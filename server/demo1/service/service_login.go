package service

import (
	"his6/server/demo1/dao"
	"his6/server/model"
)

type LoginService struct{
	empDao dao.EmpDao
	menuDao dao.MenuDao
}

func (ls *LoginService) Login(empCode string, passwd string) (model.BdictEmp, error){
	return ls.empDao.Login(empCode, passwd)
}

func (ls *LoginService) QueryMenus() ([]model.BdictMenu, error){
	return ls.menuDao.GetMenus()
}

func (ls *LoginService) QueryCacheMenus() ([]model.BdictMenu, error){
	return ls.menuDao.GetCacheMenus()
}

func (ls *LoginService) AddDept(no int, name, loc string) error{
	return ls.menuDao.AddDept(no, name, loc)
}
