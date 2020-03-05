package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"his6/base/config"
	"his6/base/convert"
	"his6/base/message"
	"his6/base/middle/model"
	sort2 "sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/kataras/iris/mvc"
)

var (
	//App : Iris服务对象
	App *iris.Application
	//longQueryTime 慢查询时间，默认5s
	longQueryTime float64

	// url监控调用桶信息，通过/metrics可以查看
	urlBuckets []*UrlBucket
	// url桶时间范围分隔
	les = []float64{.01, .025, .05, .1, .25, .5, 1, 3}
)

func init() {
	App = iris.New()

	p := pprof.New()
	longQueryTime = config.GetConfigDuration("logs",
		"long_query_time", 5 * time.Second).Seconds()

	App.Any("/debug/pprof", p)
	App.Any("/debug/pprof/{action:path}", p)
	App.Use(before)
	//App.Done(after)
}

//before 调用前处理，记录时间
func before(ctx iris.Context) {
	//  记录调用时间
	ctx.Values().Set("startTime", time.Now())
	ctx.Next() // 处理完成，执行下一个处理器。

	st := ctx.Values().Get("startTime").(time.Time)
	et := time.Now()
	bt := et.Sub(st).Seconds()

	url := ctx.GetCurrentRoute().Path()
	method := ctx.Request().Method
	status := ctx.GetStatusCode()

	// 此次调用信息入桶
	b := find(url, method)
	if b.url != "N/A" {
		if status <= 299 {
			bucket(b, bt)
		} else {
			bucketFail(b)
		}
	} else {
		err := errors.New("监视错误：url[" + url + "]未定义！")
		fmt.Println(err)
	}

	// 调用后处理，记录慢查询慢请求
	if bt >= longQueryTime {
		logs.Debug(ctx.Request().URL.String() + " 慢请求：" +
			strconv.FormatFloat(bt, 'E', -1, 64) + "秒")

		// 超过时间，记录慢请求
		var slowreq = model.SlowRequest{}
		slowreq.Reqtime = et
		slowreq.URL = ctx.Request().RequestURI
		slowreq.Duration = bt

		cvb, err := json.Marshal(slowreq)
		if err != nil {
			logs.Error("慢查询json转化失败")
		}
		cvs := convert.Byte2Str(cvb)
		message.Send("slowreq", cvs)
	}
	// bucket

}

//  调用后处理，记录慢查询
//func after(ctx iris.Context) {
//
//	st := ctx.Values().Get("startTime").(time.Time)
//	et := time.Now()
//	bt := et.Sub(st).Seconds()
//	if bt >= longQueryTime {
//		logs.Debug(ctx.Request().URL.String() + " 慢请求：" +
//			strconv.FormatFloat(bt, 'E', -1, 64) + "秒")
//
//		// 超过时间，记录慢请求
//		var slowreq = model.SlowRequest{}
//		slowreq.Reqtime = et
//		slowreq.URL = ctx.Request().RequestURI
//		slowreq.Duration = bt
//
//		cvb, err := json.Marshal(slowreq)
//		if err != nil {
//			logs.Error("慢查询json转化失败")
//		}
//		cvs := convert.Byte2Str(cvb)
//		message.Send("slowreq", cvs)
//	}
//}

//find 获取url桶
func find(url, method string) *UrlBucket {
	for _, b := range urlBuckets {
		if b.url == url && b.method == method {
			return b
		}
	}
	return &UrlBucket{url:"N/A"}
}

//urlComp url桶排序方法
func urlComp(i, j int) bool{
	if urlBuckets[i].handler == urlBuckets[j].handler {
		if urlBuckets[i].method == urlBuckets[j].method {
			return urlBuckets[i].method > urlBuckets[j].method
		} else {
			return urlBuckets[i].method > urlBuckets[j].method
		}
	} else {
		return urlBuckets[i].handler > urlBuckets[j].handler
	}
}

//bucket 调用成功入桶
func bucket(b *UrlBucket, duration float64) bool {
	b.Lock()

	b.sum += duration

	for i := 7; i >= 0; {
		b.count[i + 1] += 1
		if duration >= les[i] {
			break
		} else {
			if i == 0 {
				b.count[0] += 1
			}
			i--
		}
	}

	b.Unlock()
	return true
}

//bucketFail 调用失败入桶
func bucketFail(b *UrlBucket) bool {
	b.Lock()
	b.failCount += 1
	b.Unlock()
	return true
}


//Run Web服务启动
func Run() {
	port := config.GetConfigString("app", "httpPort", "8080")

	routes := App.APIBuilder.GetRoutes()
	total := len(routes)
	urlBuckets = make([]*UrlBucket, total)
	for i := 0; i < total; i++ {
		handlerName := routes[i].MainHandlerName
		lastDot := strings.LastIndex(handlerName, ".")
		ub := UrlBucket{url:routes[i].Path, method:routes[i].Method, sum:0,
			count:[9]int{}, failCount:0, handler:handlerName[0: lastDot]}

		urlBuckets[i] = &ub
	}
	sort2.Slice(urlBuckets, urlComp)

	App.Run(iris.Addr(":" + port))
}

//GetMetrics 输出metrics信息（/metrics）
func GetMetrics() string{
	metrics := "# HELP http_request_duration_seconds_bucket " +
		"How long it took to process the succeeded request, " +
		"partitioned by method and HTTP path and bucketed by duration. \n" +
		"# HELP http_request_fail_total_bucket " +
		"Counter of request that status code >= 300.\n"
	s := "http_request_duration_seconds_bucket"
	sf := "http_request_fail_total_bucket"
	lelist := []string{"0.01", "0.025", "0.05", "0.1", "0.25", "0.5", "1.0", "3.0", "total"}

	for _, b := range urlBuckets {
		metric := s
		metric += "{status=\"success\", method=\""
		metric += b.method
		metric += "\",path=\""
		metric += b.url
		metric += "\",handler=\""
		metric += b.handler
		metric += "\",le=\""

		for i := 0; i < 9; i++ {
			mt := metric
			mt += lelist[i]
			mt += "\"} "
			mt += strconv.Itoa(b.count[i])
			mt += "\n"
			metrics += mt
		}

		metric = sf
		metric += "{status=\"fail\", method=\""
		metric += b.method
		metric += "\",path=\""
		metric += b.url
		metric += "\",handler=\""
		metric += b.handler
		metric += "\"} "
		metric += strconv.Itoa(b.failCount)
		metric += "\n"

		metrics += metric
	}
	return metrics
}

//RegisterGetHandler 注册Handler路由Get
func RegisterGetHandler(url string, handler context.Handler) {
	App.Get(url, handler)
}

//RegisterPostHandler 注册Handler路由Post
func RegisterPostHandler(url string, handler context.Handler) {
	App.Post(url, handler)
}

//RegisterHandler 注册路由（传入方法）
func RegisterHandler(method, url string, handler context.Handler) {
	App.Handle(method, url, handler)
}

//RegisterMvc 注册Handler内部所有路由
func RegisterMvc(url string, controller interface{}) {
	mvc.Configure(App.Party(url), func(app *mvc.Application) {
		app.Handle(controller)
	})
}
