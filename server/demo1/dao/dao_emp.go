package dao

import (
	"errors"
	"his6/base/crypto"
	"his6/base/database"
	"his6/server/model"
)

//EmpDao struct
type EmpDao struct {
}

//Login service
func (dao *EmpDao) Login(empCode string, passwd string) (model.BdictEmp, error) {

	pd := crypto.MD5(passwd)
	sql := "SELECT ID, CODE, NAME FROM BDICT_EMP WHERE CODE = :1 AND PASSWORD = :2 AND STATE = 1"
	var emp = model.BdictEmp{}

	err := database.OraDb.Find(&emp, sql, empCode, pd)
	if err != nil {
		err = errors.New("无此代码的在用工号或口令错误。" + err.Error())
		return emp, err
	}
	//  判断数据

	return emp, nil
}
