package model

import "his6/base/database"

//BdParamEmp
type BdParamEmp struct {
	EmpId    string
	Name     string
	Value    string
	NameChn  database.NullableString
	Describe database.NullableString
}
