package handler

import (
	"github.com/kataras/iris"
	"his6/base/router"
	"his6/server/common/model"
	"his6/server/common/service"
	"strconv"
)

var (
	login service.LoginService
)

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
	now, err := login.Sysdate()
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

	status, _ := login.Login(ety.EmpCode, ety.Password, ety.Ip)
	//if err != nil {
	//	ctx.StatusCode(iris.StatusInternalServerError)
	//	return
	//}
	s := strconv.Itoa(status)
	ctx.Text(s)
}

/// SYS-003
/// /sys/branch
func (lgn *LoginController) GetBranch(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")

	branch, err := login.GetBranch(branchId)
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

	systems, err := login.GetSystemEmp(empId)
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

	empId, err := login.GetEmpid(empCode)
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

	parmVal, err := login.GetParamEmp(empId, "DEF_SYSTEM")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	} else if parmVal.Valid {
		ctx.Text(parmVal.String)
	} else {
		ctx.Text("")
	}
}
