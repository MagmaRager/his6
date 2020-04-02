package service

import (
	"fmt"
	"github.com/kataras/iris"
	"his6/base/database"
	"his6/server/common/dao"
	"his6/server/common/model"
	"time"
)

//loginService Service Struct
type loginService struct {
	ctx iris.Context
}

func NewLogin(c iris.Context) *loginService {
	return &loginService{ctx: c}
}

//Sysdate service
func (ls *loginService) Sysdate() (time.Time, error) {
		fmt.Println(ls.ctx.Values().Get("SQLM"))
	return dao.NewLogin(ls.ctx).Sysdate()
}

//GetBranch service
func (ls *loginService) GetBranch(branchId int) (model.CdBranch, error) {
	return dao.NewLogin(ls.ctx).GetBranch(branchId)
}

//GetEmpid service
func (ls *loginService) GetEmpid(empCode string) (int, error) {
	return dao.NewLogin(ls.ctx).GetEmpid(empCode)
}

//GetSystemEmp service
func (ls *loginService) GetSystemEmp(empId int) ([]model.BdSystem, error) {
	return dao.NewLogin(ls.ctx).GetSystemEmp(empId)
}

//GetParam service
func (ls *loginService) GetParam(branchId int, parmName string) (database.NullableString, error) {
	return dao.NewLogin(ls.ctx).GetParam(branchId, parmName)
}

//SetParam service
func (ls *loginService) SetParam(branchId int, parmName, value string) error {
	i, err := dao.NewLogin(ls.ctx).SetParam(branchId, parmName, value)
	if err != nil {
		return err
	}
	if i == 0 {
		err = dao.NewLogin(ls.ctx).AddParam(branchId, parmName, value)
		if err != nil {
			return err
		}
	}

	return nil
}

//GetParamEmp service
func (ls *loginService) GetParamEmp(empId int, parmName string) (database.NullableString, error) {
	return dao.NewLogin(ls.ctx).GetParamEmp(empId, parmName)
}

//SetParamEmp service
func (ls *loginService) SetParamEmp(empId int, parmName, value string) error {
	i, err := dao.NewLogin(ls.ctx).SetParamEmp(empId, parmName, value)
	if err != nil {
		return err
	}
	if i == 0 {
		err = dao.NewLogin(ls.ctx).AddParamEmp(empId, parmName, value)
		if err != nil {
			return err
		}
	}

	return nil
}

//Login service
func (ls *loginService) Login(empCode, password string) (int, error) {
	return dao.NewLogin(ls.ctx).Login(empCode, password)
}

//GetEmpInfo service
func (ls *loginService) GetEmpInfo(empId int) (model.EmpInfo, error) {
	return dao.NewLogin(ls.ctx).GetEmpInfo(empId)
}

//GetRoleEmp service
func (ls *loginService) GetRoleEmp(empId int) ([]string, error) {
	return dao.NewLogin(ls.ctx).GetRoleEmp(empId)
}

//GetAction service
func (ls *loginService) GetAction(empId int) ([]model.Action, error) {
	return dao.NewLogin(ls.ctx).GetAction(empId)
}

//GetMenuEmp service
func (ls *loginService) GetMenuEmp(empId, systemID int, roles string) ([]model.DataMenu, error) {
	return dao.NewLogin(ls.ctx).GetMenuEmp(empId, systemID, roles)
}

