package model

import "his6/base/database"

//BdParam
type BdParam struct {
	BranchId string
	Name     string
	Value    string
	NameChn  database.NullableString
	Describe database.NullableString
}
