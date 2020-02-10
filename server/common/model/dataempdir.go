package model

import "his6/base/database"

//DataEmpDir 模块
type DataEmpDir struct {
	Id         int
	Code       string
	Name       string
	Inputcode1 string
	Inputcode2 string
	DeptId     database.NullableInt64
	DeptName   database.NullableString
}
