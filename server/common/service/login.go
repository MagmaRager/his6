package service

import (
	"his6/base/database"
	"his6/server/common/dao"
	"his6/server/common/model"
	"time"
)

//LoginService Service Struct
type LoginService struct {
	sysDao dao.SysDao
}

//Sysdate service
func (ls *LoginService) Sysdate() (time.Time, error) {
	return ls.sysDao.Sysdate()
}

//GetBranch service
func (ls *LoginService) GetBranch(branchId int) (model.CdBranch, error) {
	return ls.sysDao.GetBranch(branchId)
}

//GetEmpid service
func (ls *LoginService) GetEmpid(empCode string) (int, error) {
	return ls.sysDao.GetEmpid(empCode)
}

//GetSystemEmp service
func (ls *LoginService) GetSystemEmp(empId int) ([]model.BdSystem, error) {
	return ls.sysDao.GetSystemEmp(empId)
}

//GetParam service
func (ls *LoginService) GetParam(branchId int, parmName string) (database.NullableString, error) {
	return ls.sysDao.GetParam(branchId, parmName)
}

//SetParam service
func (ls *LoginService) SetParam(branchId int, parmName, value string) error {
	i, err := ls.sysDao.SetParam(branchId, parmName, value)
	if err != nil {
		return err
	}
	if i == 0 {
		err = ls.sysDao.AddParam(branchId, parmName, value)
		if err != nil {
			return err
		}
	}

	return nil
}

//GetParamEmp service
func (ls *LoginService) GetParamEmp(empId int, parmName string) (database.NullableString, error) {
	return ls.sysDao.GetParamEmp(empId, parmName)
}

//SetParamEmp service
func (ls *LoginService) SetParamEmp(empId int, parmName, value string) error {
	i, err := ls.sysDao.SetParamEmp(empId, parmName, value)
	if err != nil {
		return err
	}
	if i == 0 {
		err = ls.sysDao.AddParamEmp(empId, parmName, value)
		if err != nil {
			return err
		}
	}

	return nil
}

//Login service
func (ls *LoginService) Login(empCode, password string) (int, error) {
	return ls.sysDao.Login(empCode, password)
}

//GetEmpInfo service
func (ls *LoginService) GetEmpInfo(empId int) (model.EmpInfo, error) {
	return ls.sysDao.GetEmpInfo(empId)
}

//GetRoleEmp service
func (ls *LoginService) GetRoleEmp(empId int) ([]string, error) {
	return ls.sysDao.GetRoleEmp(empId)
}

//GetAction service
func (ls *LoginService) GetAction(empId int) ([]model.Action, error) {
	return ls.sysDao.GetAction(empId)
}

//GetMenuEmp service
func (ls *LoginService) GetMenuEmp(empId, systemID int, roles string) ([]model.DataMenu, error) {
	return ls.sysDao.GetMenuEmp(empId, systemID, roles)
}

