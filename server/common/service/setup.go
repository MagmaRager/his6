package service

import (
	"errors"
	"his6/base/database"
	"his6/server/common/dao"
	"his6/server/common/model"
)

//SetupService Service Struct
type SetupService struct {
	setupDao dao.SetupDao
}

//SetNewPsw service
func (ss *SetupService) SetNewPsw(empId int, opsw, npsw string) (int, error) {
	psw, err := ss.setupDao.GetPsw(empId)
	if err != nil {
		return -1, err
	}
	if psw != opsw {
		return -2, errors.New("原密码验证失败！")
	}
	err = ss.setupDao.SetNewPsw(empId, npsw)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

//GetFp service
func (ss *SetupService) GetFp(empId int, fpCode, roles string) (model.CdObjectFp, error) {
	return ss.setupDao.GetFp(empId, fpCode, roles)
}

//GetModule service
func (ss *SetupService) GetModule() ([]model.CdModule, error) {
	return ss.setupDao.GetModule()
}

//GetObject service
func (ss *SetupService) GetObject() ([]model.CdObject, error) {
	return ss.setupDao.GetObject()
}

//AddModule service
func (ss *SetupService) AddModule(module model.CdModule, objects []model.CdObject) error {
	tx, _ := database.OraDb.BeginTx()
	err := ss.setupDao.AddModule(module, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(objects); i++ {
		o := objects[i]
		err := ss.setupDao.AddObject(o, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		fps := o.FunctionPointList
		j := 0
		for ; j < len(fps); j++ {
			fp := fps[j]
			err := ss.setupDao.AddObjectFp(fp, *tx)
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
func (ss *SetupService) AddObject(objects []model.CdObject) error {
	tx, _ := database.OraDb.BeginTx()
	i := 0
	for ; i < len(objects); i++ {
		o := objects[i]
		err := ss.setupDao.AddObject(o, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		fps := o.FunctionPointList
		j := 0
		for ; j < len(fps); j++ {
			fp := fps[j]
			err := ss.setupDao.AddObjectFp(fp, *tx)
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
func (ss *SetupService) SetModule(module model.CdModule) error {
	return ss.setupDao.SetModule(module)
}

//SetObject service
func (ss *SetupService) SetObject(object model.CdObject) error {
	return ss.setupDao.SetObject(object)
}

//GetFpList service
func (ss *SetupService) GetFpList(objectCode string) ([]model.CdObjectFp, error) {
	return ss.setupDao.GetFpList(objectCode)
}

//GetFpRole service
func (ss *SetupService) GetFpRole(fpCode string) ([]model.BdRole, error) {
	return ss.setupDao.GetFpRole(fpCode)
}

//GetFpEmp service
func (ss *SetupService) GetFpEmp(fpCode string) ([]model.BdEmp, error) {
	return ss.setupDao.GetFpEmp(fpCode)
}

//GetRole service
func (ss *SetupService) GetRoleByBranch(branchId int) ([]model.BdRole, error) {
	return ss.setupDao.GetRoleByBranch(branchId)
}

//GetEmp service
func (ss *SetupService) GetEmp(branchId int) ([]model.DataEmpDir, error) {
	return ss.setupDao.GetEmp(branchId)
}

//GetSysRoleHandle service
func (ss *SetupService) GetSysRoleHandle(systemId int) ([]model.BdRole, error) {
	return ss.setupDao.GetSysRoleHandle(systemId)
}

//GetSysEmpHandle service
func (ss *SetupService) GetSysEmpHandle(systemId int) ([]model.DataEmpDir, error) {
	return ss.setupDao.GetSysEmpHandle(systemId)
}

//GetMenuRoleHandle service
func (ss *SetupService) GetMenuRoleHandle(menuCode string) ([]model.BdRole, error) {
	return ss.setupDao.GetMenuRoleHandle(menuCode)
}

//GetMenuEmpHandle service
func (ss *SetupService) GetMenuEmpHandle(menuCode string) ([]model.DataEmpDir, error) {
	return ss.setupDao.GetMenuEmpHandle(menuCode)
}

//GetSystem service
func (ss *SetupService) GetSystem() ([]model.BdSystem, error) {
	return ss.setupDao.GetSystem()
}

//AddSystem service
func (ss *SetupService) AddSystem(system model.SystemWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	systeminfo := system.SystemInfo
	roles := system.RoleList
	emps := system.EmpList
	err := ss.setupDao.AddSystem(systeminfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := systeminfo.Id
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = ss.setupDao.AddSystemRole(systemId, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = ss.setupDao.AddSystemEmp(systemId, e.Id, *tx)
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
func (ss *SetupService) SetSystem(system model.SystemWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	systeminfo := system.SystemInfo
	roles := system.RoleList
	emps := system.EmpList
	err := ss.setupDao.SetSystem(systeminfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := systeminfo.Id
	err = ss.setupDao.DeleteSystemRole(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteSystemEmp(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = ss.setupDao.AddSystemRole(systemId, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = ss.setupDao.AddSystemEmp(systemId, e.Id, *tx)
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
func (ss *SetupService) DeleteSystem(systemId int) error {
	tx, _ := database.OraDb.BeginTx()

	err := ss.setupDao.DeleteSystem(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteSystemRole(systemId, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteSystemEmp(systemId, *tx)
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
func (ss *SetupService) NewSysIdCenter() (int, error) {
	return ss.setupDao.NewSysIdCenter()
}

//NewSysId service
func (ss *SetupService) NewSysId(branchId int) (int, error) {
	return ss.setupDao.NewSysId(branchId)
}

//GetAllMenu service
func (ss *SetupService) GetAllMenu() ([]model.BdMenu, error) {
	return ss.setupDao.GetAllMenu()
}

//AddMenu service
func (ss *SetupService) AddMenu(menu model.MenuWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	menuinfo := menu.MenuInfo
	roles := menu.RoleList
	emps := menu.EmpList
	err := ss.setupDao.AddMenu(menuinfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := menuinfo.SystemId
	menuCode := menuinfo.Code
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = ss.setupDao.AddMenuRole(systemId, r.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = ss.setupDao.AddMenuEmp(systemId, e.Id, menuCode, *tx)
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
func (ss *SetupService) SetMenu(menu model.MenuWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	menuinfo := menu.MenuInfo
	roles := menu.RoleList
	emps := menu.EmpList
	err := ss.setupDao.SetMenu(menuinfo, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	systemId := menuinfo.SystemId
	menuCode := menuinfo.Code
	err = ss.setupDao.DeleteMenuRole(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteMenuEmp(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = ss.setupDao.AddMenuRole(systemId, r.Id, menuCode, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = ss.setupDao.AddMenuEmp(systemId, e.Id, menuCode, *tx)
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
func (ss *SetupService) MoveMenu(os, ns int, om, nm string) error {
	tx, _ := database.OraDb.BeginTx()
	err := ss.setupDao.MoveMenu(os, ns, om, nm, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.MoveMenuRole(os, ns, om, nm, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.MoveMenuEmp(os, ns, om, nm, *tx)
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
func (ss *SetupService) DeleteMenu(systemId int, menuCode string) error {
	tx, _ := database.OraDb.BeginTx()

	err := ss.setupDao.DeleteMenu(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteMenuRole(systemId, menuCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteMenuEmp(systemId, menuCode, *tx)
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
func (ss *SetupService) SetFpRight(fp model.FpWithRight) error {
	tx, _ := database.OraDb.BeginTx()
	fpCode := fp.FpCode
	roles := fp.RoleList
	emps := fp.EmpList
	err := ss.setupDao.DeleteFpRole(fpCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ss.setupDao.DeleteFpEmp(fpCode, *tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 0
	for ; i < len(roles); i++ {
		r := roles[i]
		err = ss.setupDao.AddFpRole(fpCode, r.Id, *tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i = 0; i < len(emps); i++ {
		e := emps[i]
		err = ss.setupDao.AddFpEmp(fpCode, e.Id, *tx)
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
func (ss *SetupService) GetParams(branchId int) ([]model.BdParam, error) {
	return ss.setupDao.GetParams(branchId)
}

//SetParam service
func (ss *SetupService) SetParam(param model.BdParam) error {
	return ss.setupDao.SetParam(param)
}

//GetParamsEmp service
func (ss *SetupService) GetParamsEmp(empId int) ([]model.BdParamEmp, error) {
	return ss.setupDao.GetParamsEmp(empId)
}

//SetParamEmp service
func (ss *SetupService) SetParamEmp(param model.BdParamEmp) error {
	return ss.setupDao.SetParamEmp(param)
}

