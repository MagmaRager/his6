package dao

import (
	"errors"
	"his6/base/database"
	"his6/server/common/model"
)

//SetupDao struct
type SetupDao struct {
}

//GetPsw service
func (dao *SetupDao) GetPsw(empId int) (string, error) {
	sql := "SELECT PASSWORD FROM BD_EMP WHERE ID = :1"
	var psw = ""

	err := database.OraDb.Find(&psw, sql, empId)
	if err != nil {
		err = errors.New("未找到对应人员。" + err.Error())
		return psw, err
	}
	return psw, nil
}

//SetNewPsw service
func (dao *SetupDao) SetNewPsw(empId int, psw string) error {
	sql := "UPDATE BD_EMP SET PASSWORD = :1 WHERE ID = :2"

	tx, _ := database.OraDb.BeginTx()
	_, err := tx.Exec(sql, psw, empId)
	if err != nil {
		err = errors.New("用户不存在！" + err.Error())
		return err
	}
	tx.Commit()
	return nil
}

//GetFp service
func (dao *SetupDao) GetFp(empId int, fpCode, roles string) (model.CdObjectFp, error) {
	sql := "SELECT CODE, NAME, OBJECT_CODE, DESCRIBE FROM CD_OBJECT_FP WHERE CODE = :1 AND " +
		"(CODE IN " +
		"(SELECT FP_CODE FROM BD_OBJECT_FP_EMP WHERE EMP_ID = :2) " +
		"OR CODE IN " +
		"(SELECT FP_CODE FROM BD_OBJECT_FP_ROLE WHERE ROLE_ID IN " +
		"(SELECT ID FROM BD_ROLE WHERE ' ' || :3 || ' ' LIKE '% ' || CODE || ' %')))"
	var ofp = model.CdObjectFp{}

	err := database.OraDb.Find(&ofp, sql, fpCode, empId, roles)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return model.CdObjectFp{}, nil
		}
		err = errors.New("未能获取功能点。" + err.Error())
		return ofp, err
	}
	return ofp, nil
}

//GetModule service
func (dao *SetupDao) GetModule() ([]model.CdModule, error) {
	sql := "SELECT CODE, NAME, FILE_NAME, VERSION, USED_FLAG, DESCRIBE, FILE_TIME " +
		"FROM CD_MODULE ORDER BY CODE"
	var modules = []model.CdModule{}

	err := database.OraDb.Query(&modules, sql)
	if err != nil {
		err = errors.New("获取模块错误。" + err.Error())
		return modules, err
	}
	return modules, nil
}

//GetObject service *
func (dao *SetupDao) GetObject() ([]model.CdObject, error) {
	sql := "SELECT CODE, NAME, OBJECT, MODULE_CODE, USED_FLAG, " +
		"IS_FUNCTION, HAS_FUNCTION_POINT, DESCRIBE " +
		"FROM CD_OBJECT WHERE USED_FLAG = 1 ORDER BY MODULE_CODE, CODE"
	var objects = []model.CdObject{}

	err := database.OraDb.Query(&objects, sql)
	if err != nil {
		err = errors.New("获取对象错误。" + err.Error())
		return objects, err
	}
	return objects, nil
}

//AddModule service
func (dao *SetupDao) AddModule(m model.CdModule, tx database.Tx) error {
	sql := "INSERT INTO CD_MODULE(CODE, NAME, FILE_NAME, " +
		"VERSION, USED_FLAG, FILE_TIME, DESCRIBE, UPDATE_EMPID, UPDATE_TIME) " +
		"VALUES(:1, :2, :3, :4, 1, :5, :6, :7, SYSDATE)"
	rst, err := tx.Exec(sql, m.Code, m.Name,
		m.FileName, m.Version, m.FileTime, m.Describe, m.UpdateEmpid)
	if err != nil {
		err = errors.New("新增模块错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增模块。" + err.Error())
		return err
	}
	return nil
}

//AddObject service
func (dao *SetupDao) AddObject(o model.CdObject, tx database.Tx) error {
	sql := "INSERT INTO CD_OBJECT(CODE, NAME, OBJECT, MODULE_CODE, " +
		"USED_FLAG, IS_FUNCTION, HAS_FUNCTION_POINT, DESCRIBE)" +
		"VALUES(:1, :2, :3, :4, 1, :5, :6, :7)"

	rst, err := tx.Exec(sql, o.Code, o.Name,
		o.Object, o.ModuleCode, o.IsFunction, o.HasFunctionPoint, o.Describe)
	if err != nil {
		err = errors.New("新增对象错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增对象。" + err.Error())
		return err
	}

	return nil
}

//SetModule service
func (dao *SetupDao) SetModule(m model.CdModule) error {
	sql := "UPDATE CD_MODULE SET " +
		"NAME = :1, USED_FLAG = :2, DESCRIBE = :3 WHERE CODE = :4"
	tx, _ := database.OraDb.BeginTx()
	rst, err := tx.Exec(sql, m.Name, m.UsedFlag, m.Describe, m.Code)
	if err != nil {
		err = errors.New("新增模块错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增模块。" + err.Error())
		return err
	}
	tx.Commit()
	return nil
}

//SetObject service
func (dao *SetupDao) SetObject(o model.CdObject) error {
	sql := "UPDATE CD_OBJECT SET " +
		"NAME = :1, USED_FLAG = :2, DESCRIBE = :3 WHERE CODE = :4"
	tx, _ := database.OraDb.BeginTx()
	rst, err := tx.Exec(sql, o.Name, o.UsedFlag, o.Describe, o.Code)
	if err != nil {
		err = errors.New("新增模块错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增模块。" + err.Error())
		return err
	}
	tx.Commit()
	return nil
}

//AddObjectFp service
func (dao *SetupDao) AddObjectFp(fp model.CdObjectFp, tx database.Tx) error {
	sql := "INSERT INTO CD_OBJECT_FP(CODE, NAME, OBJECT_CODE, DESCRIBE)" +
		"VALUES(:1, :2, :3, :4)"

	rst, err := tx.Exec(sql, fp.Code, fp.Name, fp.ObjectCode, fp.Describe)
	if err != nil {
		err = errors.New("新增功能点错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增功能点。" + err.Error())
		return err
	}

	return nil
}

//GetFpList service
func (dao *SetupDao) GetFpList(ocode string) ([]model.CdObjectFp, error) {
	sql := "SELECT CODE, NAME, OBJECT_CODE, DESCRIBE FROM CD_OBJECT_FP " +
		"WHERE OBJECT_CODE = :1"
	var fps = []model.CdObjectFp{}

	err := database.OraDb.Query(&fps, sql, ocode)
	if err != nil {
		err = errors.New("获取功能点列表错误。" + err.Error())
		return fps, err
	}
	return fps, nil
}

//GetFpRole service
func (dao *SetupDao) GetFpRole(fpCode string) ([]model.BdRole, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2, STATE, IS_LEAF, DESCRIBE " +
		"FROM BD_ROLE WHERE STATE = 1 AND ID IN " +
		"(SELECT ROLE_ID FROM BD_OBJECT_FP_ROLE WHERE FP_CODE = :1) " +
		"ORDER BY CODE"
	var roles = []model.BdRole{}

	err := database.OraDb.Query(&roles, sql, fpCode)
	if err != nil {
		err = errors.New("获取功能点角色错误。" + err.Error())
		return roles, err
	}
	return roles, nil
}

//GetFpEmp service
func (dao *SetupDao) GetFpEmp(fpCode string) ([]model.BdEmp, error) {
	sql := "SELECT ID, CEID, CODE, NAME, INPUTCODE1, INPUTCODE2, KIND_CODE, " +
		"DEPT_ID, BIZ_DEPT_ID, GROUP_ID, TITLES_ID, IS_ADMIN, IS_TEMP, TAKE_EMPID, STATE " +
		"FROM BD_EMP WHERE STATE = 1 AND ID IN " +
		"(SELECT EMP_ID FROM BD_OBJECT_FP_EMP WHERE FP_CODE = :1) " +
		"ORDER BY CODE"
	var emps = []model.BdEmp{}

	err := database.OraDb.Query(&emps, sql, fpCode)
	if err != nil {
		err = errors.New("获取功能点人员错误。" + err.Error())
		return emps, err
	}
	return emps, nil
}

//GetRoleByBranch service
func (dao *SetupDao) GetRoleByBranch(branchId int) ([]model.BdRole, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2, STATE, IS_LEAF, DESCRIBE " +
		"FROM BD_ROLE WHERE STATE = 1 AND TRUNC(ID, -3) / 1000 = :1 " +
		"ORDER BY CODE"
	var roles = []model.BdRole{}

	err := database.OraDb.Query(&roles, sql, branchId)
	if err != nil {
		err = errors.New("获取机构角色失败。" + err.Error())
		return roles, err
	}
	return roles, nil
}

//GetEmp service
func (dao *SetupDao) GetEmp(branchId int) ([]model.DataEmpDir, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2, DEPT_ID, " +
		"(SELECT NAME FROM BD_DEPT WHERE ID = BD_EMP.DEPT_ID) AS DEPT_NAME " +
		"FROM BD_EMP WHERE STATE = 1 AND TRUNC(ID, -3) / 1000 = :1 AND STATE = 1 " +
		"ORDER BY CODE"
	var emps = []model.DataEmpDir{}

	err := database.OraDb.Query(&emps, sql, branchId)
	if err != nil {
		err = errors.New("获取机构人员失败。" + err.Error())
		return emps, err
	}
	return emps, nil
}

//GetSysRoleHandle service
func (dao *SetupDao) GetSysRoleHandle(systemId int) ([]model.BdRole, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2 " +
		"FROM BD_ROLE WHERE ID IN " +
		"(SELECT ROLE_ID FROM BD_SYSTEM_ROLE WHERE SYSTEM_ID = :1) " +
		"ORDER BY CODE"
	var roles = []model.BdRole{}

	err := database.OraDb.Query(&roles, sql, systemId)
	if err != nil {
		err = errors.New("获取机构角色失败。" + err.Error())
		return roles, err
	}
	return roles, nil
}

//GetSysEmpHandle service
func (dao *SetupDao) GetSysEmpHandle(systemId int) ([]model.DataEmpDir, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2, DEPT_ID, " +
		"(SELECT NAME FROM BD_DEPT WHERE ID = BD_EMP.DEPT_ID) AS DEPT_NAME " +
		"FROM BD_EMP WHERE ID IN " +
		"(SELECT EMP_ID FROM BD_SYSTEM_EMP WHERE SYSTEM_ID = :1) " +
		"ORDER BY CODE"
	var emps = []model.DataEmpDir{}

	err := database.OraDb.Query(&emps, sql, systemId)
	if err != nil {
		err = errors.New("获取机构人员失败。" + err.Error())
		return emps, err
	}
	return emps, nil
}

//GetMenuRoleHandle service
func (dao *SetupDao) GetMenuRoleHandle(menuCode string) ([]model.BdRole, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2 " +
		"FROM BD_ROLE WHERE ID IN " +
		"(SELECT ROLE_ID FROM BD_MENU_ROLE WHERE MENU_CODE = :1) " +
		"ORDER BY CODE"
	var roles = []model.BdRole{}

	err := database.OraDb.Query(&roles, sql, menuCode)
	if err != nil {
		err = errors.New("获取机构角色失败。" + err.Error())
		return roles, err
	}
	return roles, nil
}

//GetMenuEmpHandle service
func (dao *SetupDao) GetMenuEmpHandle(menuCode string) ([]model.DataEmpDir, error) {
	sql := "SELECT ID, CODE, NAME, INPUTCODE1, INPUTCODE2, DEPT_ID, " +
		"(SELECT NAME FROM BD_DEPT WHERE ID = BD_EMP.DEPT_ID) AS DEPT_NAME " +
		"FROM BD_EMP WHERE ID IN " +
		"(SELECT EMP_ID FROM BD_MENU_EMP WHERE MENU_CODE = :1) " +
		"ORDER BY CODE"
	var emps = []model.DataEmpDir{}

	err := database.OraDb.Query(&emps, sql, menuCode)
	if err != nil {
		err = errors.New("获取机构人员失败。" + err.Error())
		return emps, err
	}
	return emps, nil
}

//GetSystem service
func (dao *SetupDao) GetSystem() ([]model.BdSystem, error) {
	sql := "SELECT ID, CODE, NAME, ICO FROM BD_SYSTEM ORDER BY ID"
	var systems = []model.BdSystem{}

	err := database.OraDb.Query(&systems, sql)
	if err != nil {
		err = errors.New("获取子系统错误。" + err.Error())
		return systems, err
	}
	return systems, nil
}

//AddSystem service
func (dao *SetupDao) AddSystem(s model.BdSystem, tx database.Tx) error {
	sql := "INSERT INTO BD_SYSTEM(ID, CODE, NAME, ICO) " +
		"VALUES(:1, :2, :3, :4)"

	rst, err := tx.Exec(sql, s.Id, s.Code, s.Name, s.Ico)
	if err != nil {
		err = errors.New("新增系统错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增系统。" + err.Error())
		return err
	}

	return nil
}

//SetSystem service
func (dao *SetupDao) SetSystem(s model.BdSystem, tx database.Tx) error {
	sql := "UPDATE BD_SYSTEM SET CODE = :1, NAME = :2, ICO = :3 WHERE ID = :4"

	rst, err := tx.Exec(sql, s.Code, s.Name, s.Ico, s.Id)
	if err != nil {
		err = errors.New("设置系统错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功设置系统。" + err.Error())
		return err
	}

	return nil
}

//DeleteSystem service
func (dao *SetupDao) DeleteSystem(sid int, tx database.Tx) error {
	sql := "DELETE FROM BD_SYSTEM WHERE ID = :1"

	rst, err := tx.Exec(sql, sid)
	if err != nil {
		err = errors.New("删除系统角色错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功删除系统角色。" + err.Error())
		return err
	}

	return nil
}

//AddSystemRole service
func (dao *SetupDao) AddSystemRole(sid, rid int, tx database.Tx) error {
	sql := "INSERT INTO BD_SYSTEM_ROLE(SYSTEM_ID, ROLE_ID) " +
		"VALUES(:1, :2)"

	rst, err := tx.Exec(sql, sid, rid)
	if err != nil {
		err = errors.New("新增系统角色错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增系统角色。" + err.Error())
		return err
	}

	return nil
}

//AddSystemEmp service
func (dao *SetupDao) AddSystemEmp(sid, eid int, tx database.Tx) error {
	sql := "INSERT INTO BD_SYSTEM_EMP(SYSTEM_ID, EMP_ID) " +
		"VALUES(:1, :2)"

	rst, err := tx.Exec(sql, sid, eid)
	if err != nil {
		err = errors.New("新增系统人员错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增系统人员。" + err.Error())
		return err
	}

	return nil
}

//DeleteSystemRole service
func (dao *SetupDao) DeleteSystemRole(sid int, tx database.Tx) error {
	sql := "DELETE FROM BD_SYSTEM_ROLE WHERE SYSTEM_ID = :1"

	_, err := tx.Exec(sql, sid)
	if err != nil {
		err = errors.New("删除系统角色错误。" + err.Error())
		return err
	}

	return nil
}

//DeleteSystemEmp service
func (dao *SetupDao) DeleteSystemEmp(sid int, tx database.Tx) error {
	sql := "DELETE FROM BD_SYSTEM_EMP WHERE SYSTEM_ID = :1"

	_, err := tx.Exec(sql, sid)
	if err != nil {
		err = errors.New("删除系统人员错误。" + err.Error())
		return err
	}

	return nil
}

//NewSysIdCenter service
func (dao *SetupDao) NewSysIdCenter() (int, error) {
	sql := "SELECT MAX(ID) FROM BD_SYSTEM WHERE ID < 10000"
	maxidf:= 0.0
	err := database.OraDb.Find(&maxidf, sql)
	maxid := int(maxidf)
	if err != nil {
		err = errors.New("获取子系统错误。" + err.Error())
		return -1, err
	}
	if maxid == 9999 {
		err = errors.New("子系统序列号溢出。" + err.Error())
		return -2, err
	}
	return maxid + 1, nil
}

//NewSysId service
func (dao *SetupDao) NewSysId(branchId int) (int, error) {
	sql := "SELECT MAX(ID) FROM BD_SYSTEM WHERE TRUNC(ID, -4) / 10000 = :1"
	maxidf := 0.0
	err := database.OraDb.Find(&maxidf, sql, branchId)
	maxid := int(maxidf)
	if err != nil {
		err = errors.New("获取子系统错误。" + err.Error())
		return -1, err
	}
	if maxid - (branchId * 10000) == 9999 {
		err = errors.New("子系统序列号溢出。" + err.Error())
		return -2, err
	}
	return maxid + 1, nil
}

//GetAllMenu service
func (dao *SetupDao) GetAllMenu() ([]model.BdMenu, error) {
	sql := "SELECT SYSTEM_ID, CODE, TITLE, OBJECT_CODE, PARAMETER, WIN_STATE, ICO, PROMPT " +
		"FROM BD_MENU ORDER BY SYSTEM_ID, CODE"
	var menus = []model.BdMenu{}

	err := database.OraDb.Query(&menus, sql)
	if err != nil {
		err = errors.New("获取子系统错误。" + err.Error())
		return menus, err
	}
	return menus, nil
}

//AddMenu service
func (dao *SetupDao) AddMenu(s model.BdMenu, tx database.Tx) error {
	sql := "INSERT INTO BD_MENU" +
	"(SYSTEM_ID, CODE, TITLE, OBJECT_CODE, ICO, PROMPT, WIN_STATE, PARAMETER) " +
		" VALUES(:1, :2, :3, :4, :5, :6, :7, :8)"

	rst, err := tx.Exec(sql, s.SystemId, s.Code, s.Title, s.ObjectCode,
		s.Ico, s.Prompt, s.WinState, s.Parameter)
	if err != nil {
		err = errors.New("新增菜单错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增菜单。" + err.Error())
		return err
	}

	return nil
}

//SetMenu service
func (dao *SetupDao) SetMenu(m model.BdMenu, tx database.Tx) error {
	sql := "UPDATE BD_MENU SET TITLE = :1, OBJECT_CODE = :2, ICO = :3, PROMPT = :4 " +
		"WHERE SYSTEM_ID = :5 AND CODE = :6"

	rst, err := tx.Exec(sql, m.Title, m.ObjectCode, m.Ico, m.Prompt, m.SystemId, m.Code)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功设置菜单。" + err.Error())
		return err
	}

	return nil
}

//MoveMenu service
func (dao *SetupDao) MoveMenu(os, ns int, om, nm string, tx database.Tx) error {
	sql := "UPDATE BD_MENU SET SYSTEM_ID = :1, CODE = :2" +
		"WHERE SYSTEM_ID = :3 AND CODE = :4"

	rst, err := tx.Exec(sql, ns, nm, os, om)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功设置菜单。" + err.Error())
		return err
	}

	return nil
}

//DeleteMenu service
func (dao *SetupDao) DeleteMenu(sid int, mcode string, tx database.Tx) error {
	sql := "DELETE FROM BD_MENU WHERE SYSTEM_ID = :1 AND CODE = :2"

	rst, err := tx.Exec(sql, sid, mcode)
	if err != nil {
		err = errors.New("删除错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功删除系统角色。" + err.Error())
		return err
	}

	return nil
}

//AddMenuRole service
func (dao *SetupDao) AddMenuRole(sid, rid int, mcode string, tx database.Tx) error {
	sql := "INSERT INTO BD_MENU_ROLE(SYSTEM_ID, MENU_CODE, ROLE_ID) " +
		"VALUES(:1, :2, :3)"

	rst, err := tx.Exec(sql, sid, mcode, rid)
	if err != nil {
		err = errors.New("新增菜单角色错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增菜单角色。" + err.Error())
		return err
	}

	return nil
}

//AddMenuEmp service
func (dao *SetupDao) AddMenuEmp(sid, eid int, mcode string, tx database.Tx) error {
	sql := "INSERT INTO BD_MENU_EMP(SYSTEM_ID, MENU_CODE, EMP_ID) " +
		"VALUES(:1, :2, :3)"

	rst, err := tx.Exec(sql, sid, mcode, eid)
	if err != nil {
		err = errors.New("新增菜单人员错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增菜单人员。" + err.Error())
		return err
	}

	return nil
}

//MoveMenuRole service
func (dao *SetupDao) MoveMenuRole(os, ns int, om, nm string, tx database.Tx) error {
	sql := "UPDATE BD_MENU_ROLE SET SYSTEM_ID = :1, MENU_CODE = :2" +
		"WHERE SYSTEM_ID = :3 AND MENU_CODE = :4"

	_, err := tx.Exec(sql, ns, nm, os, om)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}

	return nil
}

//MoveMenuEmp service
func (dao *SetupDao) MoveMenuEmp(os, ns int, om, nm string, tx database.Tx) error {
	sql := "UPDATE BD_MENU_EMP SET SYSTEM_ID = :1, MENU_CODE = :2" +
		"WHERE SYSTEM_ID = :3 AND MENU_CODE = :4"

	_, err := tx.Exec(sql, ns, nm, os, om)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}

	return nil
}

//DeleteMenuRole service
func (dao *SetupDao) DeleteMenuRole(sid int, mcode string, tx database.Tx) error {
	sql := "DELETE FROM BD_MENU_ROLE WHERE SYSTEM_ID = :1 AND MENU_CODE = :2"

	_, err := tx.Exec(sql, sid, mcode)
	if err != nil {
		err = errors.New("删除系统角色错误。" + err.Error())
		return err
	}

	return nil
}

//DeleteMenuEmp service
func (dao *SetupDao) DeleteMenuEmp(sid int, mcode string, tx database.Tx) error {
	sql := "DELETE FROM BD_MENU_EMP WHERE SYSTEM_ID = :1 AND MENU_CODE = :2"

	_, err := tx.Exec(sql, sid, mcode)
	if err != nil {
		err = errors.New("删除系统人员错误。" + err.Error())
		return err
	}

	return nil
}

//AddFpRole service
func (dao *SetupDao) AddFpRole(fpCode string, rid int, tx database.Tx) error {
	sql := "INSERT INTO BD_OBJECT_FP_ROLE(FP_CODE, ROLE_ID) VALUES(:1, :2)"

	rst, err := tx.Exec(sql, fpCode, rid)
	if err != nil {
		err = errors.New("新增菜单角色错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增菜单角色。" + err.Error())
		return err
	}

	return nil
}

//AddFpEmp service
func (dao *SetupDao) AddFpEmp(fpCode string, eid int, tx database.Tx) error {
	sql := "INSERT INTO BD_OBJECT_FP_EMP(FP_CODE, EMP_ID) VALUES(:1, :2)"

	rst, err := tx.Exec(sql, fpCode, eid)
	if err != nil {
		err = errors.New("新增菜单人员错误。" + err.Error())
		return err
	}
	if rst != 1 {
		err = errors.New("未成功新增菜单人员。" + err.Error())
		return err
	}

	return nil
}

//DeleteFpRole service
func (dao *SetupDao) DeleteFpRole(fpCode string, tx database.Tx) error {
	sql := "DELETE FROM BD_OBJECT_FP_ROLE WHERE FP_CODE = :1"

	_, err := tx.Exec(sql, fpCode)
	if err != nil {
		err = errors.New("删除系统角色错误。" + err.Error())
		return err
	}

	return nil
}

//DeleteFpEmp service
func (dao *SetupDao) DeleteFpEmp(fpCode string, tx database.Tx) error {
	sql := "DELETE FROM BD_OBJECT_FP_EMP WHERE FP_CODE = :1"

	_, err := tx.Exec(sql, fpCode)
	if err != nil {
		err = errors.New("删除系统人员错误。" + err.Error())
		return err
	}

	return nil
}

//GetParams service
func (dao *SetupDao) GetParams(branchId int) ([]model.BdParam, error) {
	sql := "SELECT BRANCH_ID, NAME, VALUE, NAME_CHN, DESCRIBE FROM BD_PARAMETER " +
		"WHERE BRANCH_ID = :1"
	var parms = []model.BdParam{}

	err := database.OraDb.Query(&parms, sql, branchId)
	if err != nil {
		err = errors.New("无对应参数值。" + err.Error())
		return nil, err
	}
	return parms, nil
}

//SetParam service
func (dao *SetupDao) SetParam(parm model.BdParam) error {
	sql := "UPDATE BD_PARAMETER SET VALUE = :1, NAME_CHN = :2, DESCRIBE = :3 " +
		"WHERE BRANCH_ID = :4 AND NAME = :5"

	tx, _ := database.OraDb.BeginTx()
	_, err := tx.Exec(sql, parm.Value, parm.NameChn, parm.Describe, parm.BranchId, parm.Name)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}

	return nil
}

//GetParamsEmp service
func (dao *SetupDao) GetParamsEmp(empId int) ([]model.BdParamEmp, error) {
	sql := "SELECT EMP_ID, NAME, VALUE, NAME_CHN, DESCRIBE FROM BD_PARAMETER_EMP " +
		"WHERE EMP_ID = :1"
	var parms = []model.BdParamEmp{}

	err := database.OraDb.Query(&parms, sql, empId)
	if err != nil {
		err = errors.New("无对应参数值。" + err.Error())
		return nil, err
	}
	return parms, nil
}

//SetParamEmp service
func (dao *SetupDao) SetParamEmp(parm model.BdParamEmp) error {
	sql := "UPDATE BD_PARAMETER_EMP SET VALUE = :1, NAME_CHN = :2, DESCRIBE = :3 " +
		"WHERE EMP_ID = :4 AND NAME = :5"

	tx, _ := database.OraDb.BeginTx()
	_, err := tx.Exec(sql, parm.Value, parm.NameChn, parm.Describe, parm.EmpId, parm.Name)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		return err
	}

	return nil
}