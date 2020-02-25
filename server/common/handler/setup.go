package handler

import (
	"encoding/json"
	"github.com/kataras/iris"
	"his6/base/router"
	"his6/server/common/model"
	"his6/server/common/service"
	"strconv"
)

var (
	setup service.SetupService
)

//SetupController struct
type SetupController struct {
}

func init() {
	router.RegisterMvc("/setup", new(SetupController))
}

//func init() {
//	router.RegisterPostHandler("/setup/newpsw", newpswHandle)	//SETUP-001
//	router.RegisterGetHandler("/setup/getfp", getFpHandle)	//SETUP-002
//	router.RegisterGetHandler("/setup/getmodule", getModuleHandle)	//SETUP-003
//	router.RegisterGetHandler("/setup/getobject", getObjectHandle)	//SETUP-004
//	router.RegisterPostHandler("/setup/addmodule", addModuleHandle)	//SETUP-005
//	router.RegisterPostHandler("/setup/addobject", addObjectHandle)	//SETUP-006
//	router.RegisterPostHandler("/setup/setmodule", setModuleHandle)	//SETUP-006
//	router.RegisterPostHandler("/setup/setobject", setObjectHandle)	//SETUP-007
//	router.RegisterGetHandler("/setup/getfplist", getFpListHandle)	//SETUP-008
//	router.RegisterGetHandler("/setup/getfprole", getFpRoleHandle)	//SETUP-009
//	router.RegisterGetHandler("/setup/getfpemp", getFpEmpHandle)	//SETUP-010
//	router.RegisterGetHandler("/setup/getrole", getRoleByBranchHandle)	//SETUP-011
//	router.RegisterGetHandler("/setup/getemp", getEmpHandle)	//SETUP-012
//	router.RegisterGetHandler("/setup/getsysrole", getSysRoleHandle)	//SETUP-013
//	router.RegisterGetHandler("/setup/getsysemp", getSysEmpHandle)	//SETUP-014
//	router.RegisterGetHandler("/setup/getmenurole", getMenuRoleHandle)	//SETUP-015
//	router.RegisterGetHandler("/setup/getmenuemp", getMenuEmpHandle)	//SETUP-016
//	router.RegisterGetHandler("/setup/getsystem", getSystemHandle)	//SETUP-017
//	router.RegisterPostHandler("/setup/addsystem", addSystemHandle)	//SETUP-018
//	router.RegisterPostHandler("/setup/setsystem", setSystemHandle)	//SETUP-019
//	router.RegisterPostHandler("/setup/deletesystem", deleteSystemHandle)	//SETUP-020
//	router.RegisterGetHandler("/setup/newsysidcenter", newSysIdCenterHandle)	//SETUP-021
//	router.RegisterGetHandler("/setup/newsysid", newSysIdHandle)	//SETUP-022
//	router.RegisterGetHandler("/setup/getmenu", getAllMenuHandle)	//SETUP-023
//	router.RegisterPostHandler("/setup/addmenu", addMenuHandle)	//SETUP-024
//	router.RegisterPostHandler("/setup/setmenu", setMenuHandle)	//SETUP-025
//	router.RegisterPostHandler("/setup/movemenu", moveMenuHandle)	//SETUP-026
//	router.RegisterPostHandler("/setup/deletemenu", deleteMenuHandle)	//SETUP-027
//	router.RegisterPostHandler("/setup/setfpright", setFpRightHandle)	//SETUP-028
//	router.RegisterGetHandler("/setup/getparams", getParamsHandle)	//SETUP-029
//	router.RegisterGetHandler("/setup/getparamsemp", getParamsEmpHandle)	//SETUP-030
//	router.RegisterPostHandler("/setup/saveparam", saveParamHandle)	//SETUP-031
//	router.RegisterPostHandler("/setup/saveparamemp", saveParamEmpHandle)	//SETUP-032
//
//}

/// SETUP-001
/// /setup/newpsw
func (stp *SetupController) PostNewpsw(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")
	opsw := ctx.URLParam("opsw")
	npsw := ctx.URLParam("npsw")

	i, err := setup.SetNewPsw(empId, opsw, npsw)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	s := strconv.Itoa(i)
	ctx.Text(s)
}

/// SETUP-002
/// /setup/funcpoint
func (stp *SetupController) GetFuncpoint(ctx iris.Context) {
	fpCode := ctx.URLParam("fpCode")
	empId, _ := ctx.URLParamInt("empId")
	roles := ctx.URLParam("roles")

	fp, err := setup.GetFp(empId, fpCode, roles)
	if err != nil {
		// 业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(fp)
}

/// SETUP-003
/// /setup/module
func (stp *SetupController) GetModule(ctx iris.Context) {
	modules, err := setup.GetModule()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(modules)
}

/// SETUP-004
/// /setup/object
func (stp *SetupController) GetObject(ctx iris.Context) {
	objects, err := setup.GetObject()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(objects)
}

/// SETUP-005
/// /setup/moduleadd
func (stp *SetupController) PostModuleadd(ctx iris.Context) {
	var ety = model.ModuleObjectInfo{}
	_ = ctx.ReadForm(&ety)

	var module = model.CdModule{}
	json.Unmarshal([]byte(ety.ModuleJson), &module)

	var objects = []model.CdObject{}
	json.Unmarshal([]byte(ety.ObjectJson), &objects)

	err := setup.AddModule(module, objects)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Next()
		return
	}

	ctx.JSON(1)
}

/// SETUP-006
/// /setup/objectadd
func (stp *SetupController) PostObjectadd(ctx iris.Context) {
	var ety = []model.CdObject{}
	_ = ctx.ReadJSON(&ety)

	err := setup.AddObject(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Next()
		return
	}

	ctx.JSON(1)
}

/// SETUP-007
/// /setup/moduleset
func (stp *SetupController) PostModuleset(ctx iris.Context) {
	var ety = model.CdModule{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetModule(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text("1")
}

/// SETUP-008
/// /setup/objectset
func (stp *SetupController) PostObjectset(ctx iris.Context) {
	var ety = model.CdObject{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetObject(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Next()
		return
	}

	ctx.Text("1")
}

/// SETUP-009
/// /setup/fplist
func (stp *SetupController) GetFplist(ctx iris.Context) {
	objectCode := ctx.URLParam("objectCode")

	fps, err := setup.GetFpList(objectCode)
		if err != nil {
		// 业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
	} else {
		ctx.JSON(fps)
	}
}

/// SETUP-010
/// /setup/fprole
func (stp *SetupController) GetFprole(ctx iris.Context) {
	fpCode := ctx.URLParam("fpCode")

	roles, err := setup.GetFpRole(fpCode)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(roles)
}

/// SETUP-011
/// /setup/fpemp
func (stp *SetupController) GetFpemp(ctx iris.Context) {
	fpCode := ctx.URLParam("fpCode")

	emps, err := setup.GetFpEmp(fpCode)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(emps)
}

/// SETUP-012
/// /setup/role
func (stp *SetupController) GetRole(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")

	roles, err := setup.GetRoleByBranch(branchId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(roles)
}

/// SETUP-013
/// /setup/emp
func (stp *SetupController) GetEmp(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")

	emps, err := setup.GetEmp(branchId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(emps)
}

/// SETUP-014
/// /setup/sysrole
func (stp *SetupController) GetSysrole(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("systemId")

	roles, err := setup.GetSysRoleHandle(branchId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(roles)
}

/// SETUP-015
/// /setup/sysemp
func (stp *SetupController) GetSysemp(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("systemId")

	emps, err := setup.GetSysEmpHandle(branchId)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(emps)
}

/// SETUP-016
/// /setup/menurole
func (stp *SetupController) GetMenurole(ctx iris.Context) {
	menuCode := ctx.URLParam("menuCode")

	roles, err := setup.GetMenuRoleHandle(menuCode)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(roles)
}

/// SETUP-017
/// /setup/menuemp
func (stp *SetupController) GetMenuemp(ctx iris.Context) {
	menuCode := ctx.URLParam("menuCode")

	emps, err := setup.GetMenuEmpHandle(menuCode)
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(emps)
}

/// SETUP-018
/// /setup/systemget
func (stp *SetupController) GetSystemget(ctx iris.Context) {
	systems, err := setup.GetSystem()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(systems)
}

/// SETUP-019
/// /setup/systemadd
func (stp *SetupController) PostSystemadd(ctx iris.Context) {
	var ety = model.SystemWithRight{}
	_ = ctx.ReadJSON(&ety)

	err := setup.AddSystem(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("1")
}

/// SETUP-020
/// /setup/systemset
func (stp *SetupController) PostSystemset(ctx iris.Context) {
	var ety = model.SystemWithRight{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetSystem(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(1)
}

/// SETUP-021
/// /setup/systemdelete
func (stp *SetupController) PostSystemdelete(ctx iris.Context) {
	//systemId := -1
	//_ = ctx.ReadForm(&systemId)
	systemId, _ := ctx.URLParamInt("systemId")

	err := setup.DeleteSystem(systemId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("1")
}

/// SETUP-022
/// /setup/newsysidcenter
func (stp *SetupController) GetNewsysidcenter(ctx iris.Context) {
	nid, err := setup.NewSysIdCenter()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	}
	s := strconv.Itoa(nid)
	ctx.Text(s)
}

/// SETUP-023
/// /setup/newsysid
func (stp *SetupController) GetNewsysid(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")
	nid, err := setup.NewSysId(branchId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
	}
	s := strconv.Itoa(nid)
	ctx.Text(s)
}

/// SETUP-024
/// /setup/menuget
func (stp *SetupController) GetMenuget(ctx iris.Context) {
	menus, err := setup.GetAllMenu()
	if err != nil {
		//  业务错误
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(menus)
}

/// SETUP-025
/// /setup/menuadd
func (stp *SetupController) PostMenuadd(ctx iris.Context) {
	var ety = model.MenuWithRight{}
	_ = ctx.ReadJSON(&ety)

	err := setup.AddMenu(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(1)
}

/// SETUP-026
/// /setup/menuset
func (stp *SetupController) PostMenuset(ctx iris.Context) {
	var ety = model.MenuWithRight{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetMenu(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON("1")
}

/// SETUP-027
/// /setup/menumove
func (stp *SetupController) PostMenumove(ctx iris.Context) {
	os, _ := ctx.URLParamInt("oldSystem")
	om := ctx.URLParam("oldMenu")
	ns, _ := ctx.URLParamInt("newSystem")
	nm := ctx.URLParam("newMenu")

	err := setup.MoveMenu(os, ns, om, nm)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("1")
}

/// SETUP-028
/// /setup/menudelete
func (stp *SetupController) PostMenudelete(ctx iris.Context) {
	systemId, _ := ctx.URLParamInt("systemId")
	menuCode := ctx.URLParam("menuCode")

	err := setup.DeleteMenu(systemId, menuCode)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("1")
}

/// SETUP-029
/// /setup/fprightset
func (stp *SetupController) PostFprightset(ctx iris.Context) {
	var ety = model.FpWithRight{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetFpRight(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON("1")
}

/// SETUP-030
/// /setup/params
func (stp *SetupController) GetParams(ctx iris.Context) {
	branchId, _ := ctx.URLParamInt("branchId")

	parms, err := setup.GetParams(branchId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(parms)
}

/// SETUP-031
/// /setup/paramset
func (stp *SetupController) PostParamset(ctx iris.Context) {
	var ety = model.BdParam{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetParam(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text("1")
}

/// SETUP-032
/// /setup/paramsemp
func (stp *SetupController) GetParamsemp(ctx iris.Context) {
	empId, _ := ctx.URLParamInt("empId")

	parms, err := setup.GetParamsEmp(empId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.JSON(parms)
}

/// SETUP-033
/// /setup/paramempset
func (stp *SetupController) PostParamempset(ctx iris.Context) {
	var ety = model.BdParamEmp{}
	_ = ctx.ReadJSON(&ety)

	err := setup.SetParamEmp(ety)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text("1")
}

