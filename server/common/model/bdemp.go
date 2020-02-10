package model

import (
	"his6/base/database"
)

//BdEmp 模块
type BdEmp struct {
	Id          int
	Ceid        int
	Code        string
	Name        string
	Inputcode1  string
	Inputcode2  string
	KindCode    string
	DeptId      int
	BizDeptId   int
	GroupId     database.NullableInt64
	TitlesId    database.NullableInt64
	IsAdmin     int
	IsTemp      int
	TakeEmpid   database.NullableInt64
	State       int
	ModifyEmpid database.NullableInt64
	ModifyTime  string
}
