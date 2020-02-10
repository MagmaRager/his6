package router

import (
	"encoding/json"
	"his6/base/config"
	"his6/base/convert"
	"his6/base/message"
	"his6/base/middle/model"
	"strconv"
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
)

func init() {
	App = iris.New()

	p := pprof.New()
	longQueryTime = config.GetConfigDuration("logs", "long_query_time",
		time.Duration(5*time.Second)).Seconds()

	RegisterGetHandler("/check", checkHandler)

	App.Any("/debug/pprof", p)
	App.Any("/debug/pprof/{action:path}", p)
	App.Use(before)
	App.Done(after)
}

//  调用前处理，记录时间
func before(ctx iris.Context) {
	//  记录调用时间
	ctx.Values().Set("startTime", time.Now())

	ctx.Next() // 执行下一个处理器。
}

//  调用后处理，记录慢查询
func after(ctx iris.Context) {

	st := ctx.Values().Get("startTime").(time.Time)
	et := time.Now()
	bt := et.Sub(st).Seconds()
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
}

func checkHandler(ctx iris.Context) {
	ctx.Text("OK")
}

//Run Web服务启动
func Run() {
	port := config.GetConfigString("app", "httpPort", "8080")
	App.Run(iris.Addr(":" + port))
}

//RegisterGetHandler 注册Handler路由Get
func RegisterGetHandler(url string, handler context.Handler) {
	App.Get(url, handler)
}

//RegisterPostHandler 注册Handler路由Post
func RegisterPostHandler(url string, handler context.Handler) {
	App.Post(url, handler)
}

func RegisterHandler(method, url string, handler context.Handler) {
	App.Handle(method, url, handler)
}

func RegisterMvc(url string, controller interface{}) {
	mvc.Configure(App.Party(url), func(app *mvc.Application) {
		app.Handle(controller)
	})
}
