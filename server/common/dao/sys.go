package dao

import (
	"errors"
	"his6/base/database"
	"his6/server/common/model"
	"time"
)

//SysDao struct
type SysDao struct {
}

//Sysdate service
func (dao *SysDao) Sysdate() (time.Time, error) {
	sql := "SELECT SYSDATE FROM DUAL"
	var sysdate = time.Time{}

	err := database.OraDb.Find(&sysdate, sql)
	if err != nil {
		err = errors.New("获取系统时间失败！" + err.Error())
		return sysdate, err
	}
	return sysdate, nil
}

//GetBranch service
func (dao *SysDao) GetBranch(branchId int) (model.CdBranch, error) {
	sql := "SELECT ID, CODE, NAME, SHORT_NAME FROM CD_BRANCH WHERE ID = :1 AND USED_FLAG = 1"
	var branch = model.CdBranch{}

	err := database.OraDb.Find(&branch, sql, branchId)
	if err != nil {
		err = errors.New("未找到对应机构。" + err.Error())
		return branch, err
	}
	return branch, nil
}

//GetEmpid service
func (dao *SysDao) GetEmpid(empCode string) (int, error) {
	sql := "SELECT ID FROM BD_EMP WHERE CODE = :1 AND STATE = 1"
	var id = -1

	err := database.OraDb.Find(&id, sql, empCode)
	if err != nil {
		err = errors.New("未找到对应人员。" + err.Error())
		return id, err
	}
	return id, nil
}

//GetSystemEmp service
func (dao *SysDao) GetSystemEmp(empId int) ([]model.BdSystem, error) {
	sql := "SELECT ID, CODE, NAME, ICO FROM BD_SYSTEM WHERE ID IN " +
		"(SELECT SYSTEM_ID FROM BD_SYSTEM_EMP WHERE EMP_ID = :1) " +
		"OR ID IN " +
		"(SELECT SYSTEM_ID FROM BD_SYSTEM_ROLE WHERE ROLE_ID IN " +
		"(SELECT ROLE_ID FROM BD_ROLE_EMP WHERE EMP_ID = :1)) " +
		"ORDER BY ID"
	var systems []model.BdSystem

	err := database.OraDb.Query(&systems, sql, empId)
	if err != nil {
		err = errors.New("获取子系统失败。" + err.Error())
		return systems, err
	}
	return systems, nil
}

//GetParam service
func (dao *SysDao) GetParam(branchId int, parmName string) (database.NullableString, error) {
	sql := "SELECT VALUE FROM BD_PARAMETER WHERE BRANCH_ID = :1 AND NAME = :2"
	var paramVal = database.NullableString{}

	err := database.OraDb.Find(&paramVal, sql, branchId, parmName)
	if err != nil {
		err = errors.New("无对应参数值。" + err.Error())
		return paramVal, err
	}
	return paramVal, nil
}

//SetParam service
func (dao *SysDao) SetParam(branchId int, parmName, value string) (int, error) {
	sql := "UPDATE BD_PARAMETER SET VALUE = :1 " +
		"WHERE BRANCH_ID = :2 AND NAME = :3"

	tx, _ := database.OraDb.BeginTx()
	rst, err := tx.Exec(sql, value, branchId, parmName)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		tx.Rollback()
		return -1, err
	}
	if rst != 1 {
		tx.Rollback()
		return 0, nil
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("SetParam提交失败！" + err.Error())
		return -2, err
	}
	return 1, nil
}

//AddParam service
func (dao *SysDao) AddParam(branchId int, parmName, value string) error {
	sql := "INSERT INTO BD_PARAMETER(BRANCH_ID, NAME, VALUE) " +
		"VALUES(:1, :2, :3)"

	tx, _ := database.OraDb.BeginTx()
	_, err := tx.Exec(sql, branchId, parmName, value)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("AddParam提交失败！" + err.Error())
		return err
	}
	return nil
}

//GetParamEmp service
func (dao *SysDao) GetParamEmp(empId int, parmName string) (database.NullableString, error) {
	sql := "SELECT VALUE FROM BD_PARAMETER_EMP WHERE EMP_ID = :1 AND NAME = :2"
	var paramVal = database.NullableString{}

	err := database.OraDb.Find(&paramVal, sql, empId, parmName)
	if err != nil {
		err = errors.New("无对应参数值。" + err.Error())
		return paramVal, err
	}
	return paramVal, nil
}

//SetParamEmp service
func (dao *SysDao) SetParamEmp(empId int, parmName, value string) (int, error) {
	sql := "UPDATE BD_PARAMETER_EMP SET VALUE = :1 " +
		"WHERE EMP_ID = :2 AND NAME = :5"

	tx, _ := database.OraDb.BeginTx()
	rst, err := tx.Exec(sql, value, empId, parmName)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		tx.Rollback()
		return -1, err
	}
	if rst != 1 {
		tx.Rollback()
		return 0, nil
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("SetParamEmp提交失败！" + err.Error())
		return -2, err
	}
	return 1, nil
}

//AddParamEmp service
func (dao *SysDao) AddParamEmp(empId int, parmName, value string) error {
	sql := "INSERT INTO BD_PARAMETER_EMP(EMP_ID, NAME, VALUE) " +
		"VALUES(:1, :2, :3)"

	tx, _ := database.OraDb.BeginTx()
	_, err := tx.Exec(sql, empId, parmName, value)
	if err != nil {
		err = errors.New("设置菜单错误。" + err.Error())
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.New("AddParamEmp提交失败！" + err.Error())
		return err
	}
	return nil
}

//Login service
func (dao *SysDao) Login(empCode, password string) (int, error) {
	sql := "SELECT ID, PASSWORD FROM BD_EMP WHERE CODE = :1 AND STATE = 1"
	var info = model.LoginInfo{}

	err := database.OraDb.Find(&info, sql, empCode)
	if err != nil {
		err = errors.New("用户不存在！" + err.Error())
		return -1, err
	}
	if password != info.Password {
		err = errors.New("密码错误！")
		return -2, err
	}
	return 1, nil
}

//GetEmpInfo service
func (dao *SysDao) GetEmpInfo(empId int) (model.EmpInfo, error) {
	sql := "SELECT E.ID, E.CEID, E.CODE, E.NAME, E.DEPT_ID, " +
		"E.BIZ_DEPT_ID, E.GROUP_ID, E.TITLES_ID, E.TAKE_EMPID, " +
		"DECODE(E.DEPT_ID, NULL, NULL, (SELECT NAME FROM BD_DEPT WHERE ID = E.DEPT_ID)) AS DEPT_NAME, " +
		"DECODE(E.BIZ_DEPT_ID, NULL, NULL, (SELECT NAME FROM BD_DEPT WHERE ID = E.BIZ_DEPT_ID)) AS BIZ_DEPT_NAME, " +
		"DECODE(E.GROUP_ID, NULL, NULL, (SELECT NAME FROM BD_WORK_GROUP WHERE ID = E.GROUP_ID)) AS GROUP_NAME, " +
		"DECODE(E.TITLES_ID, NULL, NULL, (SELECT NAME FROM CD_TITLES_GRADE WHERE ID = E.TITLES_ID)) AS TITLES_NAME " +
		"FROM BD_EMP E WHERE ID = :1 AND STATE = 1"
	var empInfo = model.EmpInfo{}

	err := database.OraDb.Find(&empInfo, sql, empId)
	if err != nil {
		err = errors.New("获取人员信息失败。" + err.Error())
		return empInfo, err
	}
	return empInfo, nil
}

//GetRoleEmp service
func (dao *SysDao) GetRoleEmp(empId int) ([]string, error) {
	sql := "SELECT CODE FROM BD_ROLE WHERE ID IN " +
		"(SELECT ROLE_ID FROM BD_ROLE_EMP WHERE EMP_ID = :1) AND STATE = 1"
	var roles []string

	err := database.OraDb.Query(&roles, sql, empId)
	if err != nil {
		err = errors.New("获取角色列表失败。" + err.Error())
		return roles, err
	}
	return roles, nil
}

//GetAction service
func (dao *SysDao) GetAction(empId int) ([]model.Action, error) {
	sql := "SELECT ACTION_CODE AS CODE, GRANT_EMPID, " +
		"DECODE(EFFECTIVE_TIME, NULL, '', TO_CHAR(EFFECTIVE_TIME,'YYYY-MM-DD HH24:MI:SS')) AS EFFECTIVE_TIME, " +
		"DECODE(EXPIRY_TIME, NULL, '', TO_CHAR(EXPIRY_TIME,'YYYY-MM-DD HH24:MI:SS')) AS EXPIRY_TIME " +
		"FROM BD_ACTION_ROLE_EMP " +
		"WHERE EMP_ID = :1 AND (EXPIRY_TIME IS NULL OR EXPIRY_TIME > SYSDATE)"
	var actions []model.Action

	err := database.OraDb.Query(&actions, sql, empId)
	if err != nil {
		err = errors.New("获取行为角色信息失败。" + err.Error())
		return actions, err
	}
	return actions, nil
}

//GetMenuEmp service
func (dao *SysDao) GetMenuEmp(empId, systemID int, roles string) ([]model.DataMenu, error) {
	sql := "SELECT M.CODE, M.TITLE, M.PARAMETER, M.WIN_STATE, M.ICO, M.PROMPT, " +
		"(SELECT OBJECT FROM CD_OBJECT " +
		"WHERE CODE = M.OBJECT_CODE AND USED_FLAG = 1) AS OBJECT_NAME, " +
		"(SELECT CD_MODULE.FILE_NAME FROM CD_MODULE, CD_OBJECT " +
		"WHERE CD_MODULE.CODE = CD_OBJECT.MODULE_CODE AND " +
		"CD_OBJECT.CODE = M.OBJECT_CODE " +
		"AND CD_MODULE.USED_FLAG = 1) AS MODULE_NAME " +
		"FROM BD_MENU M " +
		"WHERE M.CODE IN " +
		"(SELECT MENU_CODE FROM BD_MENU_EMP " +
		"WHERE SYSTEM_ID = :2 AND EMP_ID = :1) " +
		"OR M.CODE IN " +
		"(SELECT MENU_CODE FROM BD_MENU_ROLE WHERE SYSTEM_ID = :2 AND ROLE_ID IN " +
		"(SELECT ID FROM BD_ROLE WHERE ' ' || :3 || ' ' LIKE '% ' || CODE || ' %')) " +
		"ORDER BY M.CODE"
	var menus []model.DataMenu

	err := database.OraDb.Query(&menus, sql, systemID, empId, systemID, roles)
	if err != nil {
		err = errors.New("获取菜单信息失败。" + err.Error())
		return menus, err
	}
	return menus, nil
}


