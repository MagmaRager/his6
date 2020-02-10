package demo1

import (
	"fmt"
	"his6/base/middle/jwt"
	"his6/base/router"
	"his6/server/demo1/service"
	"strconv"

	"github.com/kataras/iris"
)

var (
	login service.LoginService
)

func init() {
	router.RegisterGetHandler("/login", loginHandler)
	router.RegisterGetHandler("/menus", queryMenuHandle)
	router.RegisterGetHandler("/cachemenus", queryCacheMenuHandle)
	router.RegisterGetHandler("/add", addDeptHandle)
	router.RegisterPostHandler("/post", postHandler)
}

type Ety struct {
	Name string
}

//  登录demo
func loginHandler(ctx iris.Context) {
	pd := ctx.URLParam("password")
	code := ctx.URLParam("code")

	//  登录验证
	emp, err := login.Login(code, pd)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	bid := strconv.Itoa(emp.DeptId) //  临时用一下
	eid := strconv.Itoa(emp.Id)
	lwg := "A"
	ip := jwt.GetClientIp(ctx.Request())

	//  产生回写的JWT
	jt := jwt.CreateToken(bid, eid, ip, lwg)
	token, _ := jt.GenToken()
	ctx.Header(jwt.GetName(), token)

	ctx.JSON(emp)
	ctx.Next()
}

func postHandler(ctx iris.Context) {
	var ety Ety = Ety{}
	// bt, _ := ctx.GetBody()
	// str := convert.Byte2Str(bt)
	_ = ctx.ReadForm(&ety) //.FormValue("name")

	//header := ctx.GetHeader("content-type")

	mtd := ctx.Method()
	fmt.Println(ety.Name + mtd)

	ctx.Next()
}

// func mongoHandler(ctx iris.Context) bool {
// 	mongo, err := mgo.Dial("127.0.0.1") // 建立连接

// 	defer mongo.Close()

// 	if err != nil {
// 		return false
// 	}

// 	client := mongo.DB("mydb_tutorial").C("t_student") //选择数据库和集合

// 	//创建数据
// 	data := Student{
// 		Name:   "seeta",
// 		Age:    18,
// 		Sid:    "s20180907",
// 		Status: 1,
// 	}

// 	//插入数据
// 	cErr := client.Insert(&data)

// 	if cErr != nil {
// 		return false
// 	}
// 	return true
// }

//  获取所有可用菜单
func queryMenuHandle(ctx iris.Context) {

	menus, err := login.QueryMenus()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(menus)
	ctx.Next()
}

//  获取所有可用菜单有cahce
func queryCacheMenuHandle(ctx iris.Context) {

	menus, err := login.QueryCacheMenus()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(menus)
	ctx.Next()
}

//  加入部门数据
func addDeptHandle(ctx iris.Context) {
	err := login.AddDept(50, "test", "sh")
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.ResponseWriter().WriteString("加入数据成功！")
	ctx.Next()
}
