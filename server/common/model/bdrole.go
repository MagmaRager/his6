package model

import "his6/base/database"

//BdRole 模块
type BdRole struct {
	Id          int
	Code        string
	Name        string
	Inputcode1  string
	Inputcode2  string
	State       int
	IsLeaf      int
	Describe    database.NullableString
	ModifyEmpid database.NullableInt64
	ModifyTime  string
}
