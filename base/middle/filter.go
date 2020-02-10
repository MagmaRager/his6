package middle

import (
	"github.com/astaxie/beego/logs"
	"github.com/kataras/iris"

	"his6/base/middle/jwt"
	"his6/base/router"
)

var (
	//  慢查询秒数
	longQueryTime float64
)

func init() {
	router.App.Use(before)
	router.App.Done(after)
}

//  调用前处理，验证JWT权限
func before(ctx iris.Context) {
	//  jwt认证处理
	s := ctx.GetHeader(jwt.GetName())
	token, err := jwt.CheckToken(s)
	if err != nil {
		logs.Error("JWT认证不成功。" + s)
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ip := jwt.GetClientIp(ctx.Request())

	iip := token.GetIp()
	if iip != ip {
		logs.Error("IP:" + ip + " 与Token IP:" + iip + "不一致。")
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	//  记录JWT token, 用于返回时回写
	ctx.Values().Set("jwtToken", token)

	logs.Info(token)
	ctx.Next() // 执行下一个处理器。
}

//  调用后处理，刷新JWT
func after(ctx iris.Context) {
	//  回写JWT token
	jw := ctx.Values().Get("jwtToken").(jwt.Info)
	token, _ := jw.GenToken()
	ctx.Header(jwt.GetName(), token)
}
