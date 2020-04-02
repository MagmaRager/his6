package handler

import (
	"github.com/kataras/iris"
	"his6/base/middle/jwt"
	"his6/base/router"
	"his6/server/common/model"
	"his6/server/common/service"
	"strconv"
)

//var (
//	login service.loginService
//)

//LoginController struct
type LoginController struct {
}

func init() {
	router.RegisterMvc("/sys", new(LoginController))
}

//func init() {
//	router.RegisterGetHandler("/sys/date", sysdateHandle)                      //SYS-001
//	router.RegisterPostHandler("/sys/login", loginHandle)                      //SYS-002
//	router.RegisterGetHandler("/sys/getbranch", getBranchHandle)               //SYS-003
//	router.RegisterGetHandler("/sys/getsystem", getSystemEmpHandle)            //SYS-004
//	router.RegisterGetHandler("/sys/getempid", getEmpidHandle)                 //SYS-005
//	router.RegisterGetHandler("/sys/getsystemdefault", getSystemDefaultHandle) //SYS-006
//
//}

/// SYS-001
/// /sys/date
func (lgn *LoginController) GetDate(ctx iris.Context) {
	now, err := service.NewLogin(ctx).Sysdate()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.Text(now.String())
}

/// SYS-002
/// /sys/login
func (lgn *LoginController) PostLogin(ctx iris.Context) {
	var ety = model.LoginInput{}
	_ = ctx.ReadForm(&ety)

	status, _ := service.NewLogin(ctx).Login(ety.EmpCode, ety.Password)
	s := strconv.Itoa(status)

	// 新的JWT token
	empId, _ := service.NewLogin(ctx).GetEmpid(ety.EmpCode)
	actions, _ := service.NewLogin(ctx).GetAction(empId)
	var autStr string
	for _, action := range actions {
		autStr += action.Code
		autStr += " "
	}
	autStr = autStr[0:len(autStr) - 1]

	jwtInfo := jwt.CreateToken(ety.BranchCode, strconv.Itoa(empId), ety.Ip, "PSW LOGIN", autStr)
	tokenNew, _ := jwtInfo.GenToken()
	ctx.Header(jwt.GetName(), tokenNew)

	ctx.Text(s)
}

/// SYS-003
/// /sys/branch
func (lgn *LoginController) GetBranch(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")

	branch, err := service.NewLogin(ctx).GetBranch(branchId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(branch)
}

/// SYS-004
/// /sys/system
func (lgn *LoginController) GetSystem(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	systems, err := service.NewLogin(ctx).GetSystemEmp(empId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(systems)
}

/// SYS-005
/// /sys/empid
func (lgn *LoginController) GetEmpid(ctx iris.Context) {
	empCode := ctx.URLParam("empCode")

	empId, err := service.NewLogin(ctx).GetEmpid(empCode)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	s := strconv.Itoa(empId)
	ctx.Text(s)
}

/// SYS-006
/// /sys/systemdefault
func (lgn *LoginController) GetSystemdefault(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	parmVal, err := service.NewLogin(ctx).GetParamEmp(empId, "DEF_SYSTEM")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	} else if parmVal.Valid {
		ctx.Text(parmVal.String)
	} else {
		ctx.Text("")
	}
}
