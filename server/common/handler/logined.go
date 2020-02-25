package handler

import (
	"github.com/kataras/iris"
	"his6/base/router"
	"his6/server/common/service"
)

var (
	logined service.LoginService
)

//LoginedController struct
type LoginedController struct {
}

func init() {
	router.RegisterMvc("/sys", new(LoginedController))
}

//func init() {
//	router.RegisterGetHandler("/sys/getparam", getParamHandle)        //SYS-007
//	router.RegisterPostHandler("/sys/setparam", setParamHandle)       //SYS-008
//	router.RegisterGetHandler("/sys/getparamemp", getParamEmpHandle)  //SYS-009
//	router.RegisterPostHandler("/sys/setparamemp", setParamEmpHandle) //SYS-010
//	router.RegisterGetHandler("/sys/getempinfo", getEmpInfoHandle)    //SYS-011
//	router.RegisterGetHandler("/sys/getrole", getRoleByEmpHandle)     //SYS-012
//	router.RegisterGetHandler("/sys/getaction", getActionHandle)      //SYS-013
//	router.RegisterGetHandler("/sys/getmenu", getMenuHandle)          //SYS-014
//}

/// SYS-007
/// /sys/parameter
func (lgn *LoginedController) GetParameter(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")
	parmName := ctx.URLParam("parmName")

	parmVal, err := logined.GetParam(branchId, parmName)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	} else if parmVal.Valid {
		ctx.Text(parmVal.String)
	} else {
		ctx.Text("")
	}
	ctx.Next()
}

/// SYS-008
/// /sys/parameter
func (lgn *LoginedController) PostParameter(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")
	parmName := ctx.URLParam("parmName")
	value := ctx.URLParam("value")

	err := logined.SetParam(branchId, parmName, value)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Next()
}

/// SYS-009
/// /sys/parameteremp
func (lgn *LoginedController) GetParameteremp(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")
	parmName := ctx.URLParam("parmName")

	parmVal, err := logined.GetParamEmp(empId, parmName)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	} else if parmVal.Valid {
		ctx.Text(parmVal.String)
	} else {
		ctx.Text("")
	}
	ctx.Next()
}

/// SYS-010
/// /sys/parameteremp
func (lgn *LoginedController) PostParameteremp(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")
	parmName := ctx.URLParam("parmName")
	value := ctx.URLParam("value")

	err := logined.SetParamEmp(empId, parmName, value)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Next()
}

/// SYS-011
/// /sys/empinfo
func (lgn *LoginedController) GetEmpinfo(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	empInfo, err := logined.GetEmpInfo(empId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(empInfo)
}

/// SYS-012
/// /sys/role
func (lgn *LoginedController) GetRole(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	roles, err := logined.GetRoleEmp(empId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(roles)
}

/// SYS-013
/// /sys/action
func (lgn *LoginedController) GetAction(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	actions, err := logined.GetAction(empId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(actions)
}

/// SYS-014
/// /sys/menu
func (lgn *LoginedController) GetMenu(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")
	systemId, _ := ctx.URLParamInt("systemId")
	roles := ctx.URLParam("roles")

	menus, err := logined.GetMenuEmp(empId, systemId, roles)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(menus)
}
