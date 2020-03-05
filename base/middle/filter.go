package middle

import (
	"github.com/astaxie/beego/logs"
	"github.com/kataras/iris"
	"his6/base/database"
	"his6/base/middle/jwt"
	"his6/base/router"
	"strings"
)

var (
	// 慢查询秒数
	//longQueryTime float64
	// url权限列表
	urlAuthList map[string]string
)

type UrlAuthOrigin struct {
	Url string
	Method string
	AuthList string
}

func init() {
	urlAuthList = getUrlAuth()

	router.App.Use(before)
	//router.App.Done(after)
}

func getUrlAuth() map[string]string {
	sql := "SELECT URL, METHOD, AUTH_LIST FROM CD_URL_AUTH"
	var lst []UrlAuthOrigin

	err := database.OraDb.Query(&lst, sql)
	if err != nil {
		return nil
	}
	uam := make(map[string]string, len(lst))
	for _, e := range lst {
		s := e.Method + e.Url
		v := e.AuthList
		uam[s] = v
	}
	return uam
}

//  调用前处理，验证JWT权限
func before(ctx iris.Context) {
	aul := urlAuthList[ctx.GetCurrentRoute().Name()]
	if hasAllAuth(aul) {
		ctx.Next()
		return
	}
	if len(aul) == 0 {
		//寻找最接近的url适配
		urlStr := ctx.GetCurrentRoute().Path()
		for idx := len(urlStr); ; {
			idx = strings.LastIndex(urlStr, "/")
			if idx >= 0 {
				urlStr = urlStr[0:idx]
				un := ctx.GetCurrentRoute().Method() + urlStr
				aul = urlAuthList[un+"/*"]
				if hasAllAuth(aul) {
					ctx.Next()
					return
				} else if len(aul) == 0 {
					continue	// 继续寻找
				} else {
					break	// 找到适配url,进行验证
				}
			}
		}
	}

	//  jwt认证处理
	s := ctx.GetHeader(jwt.GetName())
	token, err := jwt.CheckToken(s)
	if err != nil {
		logs.Error("JWT认证不成功: " + err.Error())
		if err.Error() == "Token is expired"{
			ctx.StatusCode(iris.StatusPreconditionFailed)
		} else {
			ctx.StatusCode(iris.StatusLengthRequired)
		}
		return
	}
	ip := jwt.GetClientIp(ctx.Request())

	iip := token.GetIp()
	if iip != ip {
		logs.Error("IP:" + ip + " 与Token IP:" + iip + "不一致。")
		ctx.StatusCode(iris.StatusExpectationFailed)
		return
	}

	authStr := " " + token.GetAuth() + " "
	autStrs := strings.Split(aul, " ")
	for _, str := range autStrs {
		if strings.Index(authStr, " " + str + " ") >= 0 {
			//  记录JWT token, 用于返回时回写
			ctx.Values().Set("jwtToken", token)
			logs.Debug(ctx.GetCurrentRoute().Path() + token.ToString())

			//  回写JWT token
			jw := ctx.Values().Get("jwtToken").(jwt.Info)
			tokenNew, _ := jw.GenToken()
			ctx.Header(jwt.GetName(), tokenNew)

			ctx.Next()
			return
		}
	}
	// 找不到可用权限
	logs.Error("没有url[" + ctx.GetCurrentRoute().Path() + "]的权限！")
	ctx.StatusCode(iris.StatusUnauthorized)
}

func hasAllAuth(aul string) bool {
	return strings.Index(aul, "**") >= 0
}

//  调用后处理，刷新JWT
//func after(ctx iris.Context) {
//	//  回写JWT token
//	jw := ctx.Values().Get("jwtToken").(jwt.Info)
//	token, _ := jw.GenToken()
//	ctx.Header(jwt.GetName(), token)
//}
