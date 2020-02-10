package demo

import (
	"his6/base/message"
	"his6/base/router"

	"github.com/kataras/iris"
)

func init() {
	router.RegisterGetHandler("/json", jsonHandler)
	//routers.RegisterMvc("/mvc", new(MyController))
}

//  返回json数据demo
func jsonHandler(ctx iris.Context) {
	p := ctx.URLParam("id")
	data := make(map[string]int)
	data[p] = 1
	//发动消息给自动服务端
	message.Send("slowreq", "json here")
	message.Send("slowsql", "json there")

	ctx.JSON(data)
	ctx.Next()
}
