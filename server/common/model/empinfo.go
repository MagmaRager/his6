package model

import (
	"his6/base/database"
)

//EmpInfo
type EmpInfo struct {
	Id          int
	Ceid        int
	Code        string
	Name        string
	DeptId      int
	DeptName    database.NullableString
	BizDeptId   int
	BizDeptName database.NullableString
	GroupId     database.NullableInt64
	GroupName   database.NullableString
	TitlesId    database.NullableInt64
	TitlesName  database.NullableString
	TakeEmpid   database.NullableInt64
}
