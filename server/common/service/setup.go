package service

import (
	"errors"
	"github.com/kataras/iris"
	"his6/base/database"
	"his6/server/common/dao"
	"his6/server/common/model"
)

//setupService Service Struct
type setupService struct {
	ctx iris.Context
	//setupDao dao.setupDao
}

func NewSetup(c iris.Context) *setupService {
	return &setupService{ctx: c}
}

//SetNewPsw service
func (ss *setupService) SetNewPsw(empId int, opsw, npsw string) (int, error) {
	psw, err := dao.NewSetup(ss.ctx).GetPsw(empId)
	if err != nil {
		return -1, err
	}
	if psw != opsw {
		return -2, errors.New("原密码验证失败！")
	}
	err = dao.NewSetup(ss.ctx).SetNewPsw(empId, npsw)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

//GetFp service
func (ss *setupService) GetFp(empId int, fpCode, roles string) (model.CdObjectFp, error) {
	return dao.NewSetup(ss.ctx).GetFp(empId, fpCode, roles)
}

//GetModule service
func (ss *setupService) GetModule() ([]model.CdModule, error) {
	return dao.NewSetup(ss.ctx).GetModule()
}

//GetObject service
func (ss *setupService) GetObject() ([]model.CdObject, error) {
	return dao.NewSetup(ss.ctx).GetObject()
}

//AddModule service
func (ss *setupService) AddModule(module model.CdModule, objects []model.CdObject) error {
	tx, _ := database.OraDb.BeginTx()
	err := dao.NewSetup(ss.ctx).AddModule(module, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(objects); i++ {
		o := objects[i]
		err := dao.NewSetup(ss.ctx).AddObject(o, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		fps := o.FunctionPointList
		j := 0
		for ; j < len(fps); j++ {
			fp := fps[j]
			err := dao.NewSetup(ss.ctx).AddObjectFp(fp, *tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("AddModule提交失败！" + err.Error())
		return err
	}
	return nil
}

//AddObject service
func (ss *setupService) AddObject(objects []model.CdObject) error {
	tx, _ := database.OraDb.BeginTx()
	i := 0
	for ; i < len(objects); i++ {
		o := objects[i]
		err := dao.NewSetup(ss.ctx).AddObject(o, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		fps := o.FunctionPointList
		j := 0
		for ; j < len(fps); j++ {
			fp := fps[j]
			err := dao.NewSetup(ss.ctx).AddObjectFp(fp, *tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

//SetModule service
func (ss *setupService) SetModule(module model.CdModule) error {
	return dao.NewSetup(ss.ctx).SetModule(module)
}

//SetObject service
func (ss *setupService) SetObject(object model.CdObject) error {
	return dao.NewSetup(ss.ctx).SetObject(object)
}

//GetFpList service
func (ss *setupService) GetFpList(objectCode string) ([]model.CdObjectFp, error) {
	return dao.NewSetup(ss.ctx).GetFpList(objectCode)
}

//GetFpRole service
func (ss *setupService) GetFpRole(fpCode string) ([]model.BdRole, error) {
	return dao.NewSetup(ss.ctx).GetFpRole(fpCode)
}

//GetFpEmp service
func (ss *setupService) GetFpEmp(fpCode string) ([]model.BdEmp, error) {
	return dao.NewSetup(ss.ctx).GetFpEmp(fpCode)
}

//GetRole service
func (ss *setupService) GetRoleByBranch(branchId int) ([]model.BdRole, error) {
	return dao.NewSetup(ss.ctx).GetRoleByBranch(branchId)
}

//GetEmp service
func (ss *setupService) GetEmp(branchId int) ([]model.DataEmpDir, error) {
	return dao.NewSetup(ss.ctx).GetEmp(branchId)
}

//GetSysRoleHandle service
func (ss *setupService) GetSysRoleHandle(systemId int) ([]model.BdRole, error) {
	return dao.NewSetup(ss.ctx).GetSysRoleHandle(systemId)
}

//GetSysEmpHandle service
func (ss *setupService) GetSysEmpHandle(systemId int) ([]model.DataEmpDir, error) {
	return dao.NewSetup(ss.ctx).GetSysEmpHandle(systemId)
}

//GetMenuRoleHandle service
func (ss *setupService) GetMenuRoleHandle(menuCode string) ([]model.BdRole, error) {
	return dao.NewSetup(ss.ctx).GetMenuRoleHandle(menuCode)
}

//GetMenuEmpHandle service
func (ss *setupService) GetMenuEmpHandle(menuCode string) ([]model.DataEmpDir, error) {
	return dao.NewSetup(ss.ctx).GetMenuEmpHandle(menuCode)
}

//GetSystem service
func (ss *setupService) GetSystem() ([]model.BdSystem, error) {
	return dao.NewSetup(ss.ctx).GetSystem()
}

//AddSystem service
func (ss *setupService) AddSystem(system model.SystemWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	systeminfo := system.SystemInfo
	roles := system.RoleList
	emps := system.EmpList
	err := dao.NewSetup(ss.ctx).AddSystem(systeminfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := systeminfo.Id
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = dao.NewSetup(ss.ctx).AddSystemRole(systemId, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = dao.NewSetup(ss.ctx).AddSystemEmp(systemId, e.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("AddSystem提交失败！" + err.Error())
		return err
	}
	return nil
}

//SetSystem service
func (ss *setupService) SetSystem(system model.SystemWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	systeminfo := system.SystemInfo
	roles := system.RoleList
	emps := system.EmpList
	err := dao.NewSetup(ss.ctx).SetSystem(systeminfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := systeminfo.Id
	err = dao.NewSetup(ss.ctx).DeleteSystemRole(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteSystemEmp(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = dao.NewSetup(ss.ctx).AddSystemRole(systemId, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = dao.NewSetup(ss.ctx).AddSystemEmp(systemId, e.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("SetSystem提交失败！" + err.Error())
		return err
	}
	return nil
}

//DeleteSystem service
func (ss *setupService) DeleteSystem(systemId int) error {
	tx, _ := database.OraDb.BeginTx()

	err := dao.NewSetup(ss.ctx).DeleteSystem(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteSystemRole(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteSystemEmp(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("DeleteSystem提交失败！" + err.Error())
		return err
	}
	return nil
}

//NewSysIdCenter service
func (ss *setupService) NewSysIdCenter() (int, error) {
	return dao.NewSetup(ss.ctx).NewSysIdCenter()
}

//NewSysId service
func (ss *setupService) NewSysId(branchId int) (int, error) {
	return dao.NewSetup(ss.ctx).NewSysId(branchId)
}

//GetAllMenu service
func (ss *setupService) GetAllMenu() ([]model.BdMenu, error) {
	return dao.NewSetup(ss.ctx).GetAllMenu()
}

//AddMenu service
func (ss *setupService) AddMenu(menu model.MenuWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	menuinfo := menu.MenuInfo
	roles := menu.RoleList
	emps := menu.EmpList
	err := dao.NewSetup(ss.ctx).AddMenu(menuinfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := menuinfo.SystemId
	menuCode := menuinfo.Code
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = dao.NewSetup(ss.ctx).AddMenuRole(systemId, r.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = dao.NewSetup(ss.ctx).AddMenuEmp(systemId, e.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("AddMenu提交失败！" + err.Error())
		return err
	}
	return nil
}

//SetMenu service
func (ss *setupService) SetMenu(menu model.MenuWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	menuinfo := menu.MenuInfo
	roles := menu.RoleList
	emps := menu.EmpList
	err := dao.NewSetup(ss.ctx).SetMenu(menuinfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := menuinfo.SystemId
	menuCode := menuinfo.Code
	err = dao.NewSetup(ss.ctx).DeleteMenuRole(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteMenuEmp(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = dao.NewSetup(ss.ctx).AddMenuRole(systemId, r.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = dao.NewSetup(ss.ctx).AddMenuEmp(systemId, e.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("SetMenu提交失败！" + err.Error())
		return err
	}
	return nil
}

//MoveMenu service
func (ss *setupService) MoveMenu(os, ns int, om, nm string) error {
	tx, _ := database.OraDb.BeginTx()
	err := dao.NewSetup(ss.ctx).MoveMenu(os, ns, om, nm, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).MoveMenuRole(os, ns, om, nm, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).MoveMenuEmp(os, ns, om, nm, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("MoveMenu提交失败！" + err.Error())
		return err
	}
	return nil
}

//DeleteMenu service
func (ss *setupService) DeleteMenu(systemId int, menuCode string) error {
	tx, _ := database.OraDb.BeginTx()

	err := dao.NewSetup(ss.ctx).DeleteMenu(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteMenuRole(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteMenuEmp(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("DeleteMenu提交失败！" + err.Error())
		return err
	}
	return nil
}

//SetFpRight service
func (ss *setupService) SetFpRight(fp model.FpWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	fpCode := fp.FpCode
	roles := fp.RoleList
	emps := fp.EmpList
	err := dao.NewSetup(ss.ctx).DeleteFpRole(fpCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = dao.NewSetup(ss.ctx).DeleteFpEmp(fpCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = dao.NewSetup(ss.ctx).AddFpRole(fpCode, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = dao.NewSetup(ss.ctx).AddFpEmp(fpCode, e.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("SetFpRight提交失败！" + err.Error())
		return err
	}
	return nil
}

//GetParams service
func (ss *setupService) GetParams(branchId int) ([]model.BdParam, error) {
	return dao.NewSetup(ss.ctx).GetParams(branchId)
}

//SetParam service
func (ss *setupService) SetParam(param model.BdParam) error {
	return dao.NewSetup(ss.ctx).SetParam(param)
}

//GetParamsEmp service
func (ss *setupService) GetParamsEmp(empId int) ([]model.BdParamEmp, error) {
	return dao.NewSetup(ss.ctx).GetParamsEmp(empId)
}

//SetParamEmp service
func (ss *setupService) SetParamEmp(param model.BdParamEmp) error {
	return dao.NewSetup(ss.ctx).SetParamEmp(param)
}

