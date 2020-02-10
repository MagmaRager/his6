package model

import "his6/base/database"

//CdBranch 中心机构
type CdBranch struct {
	Id         int
	Code       string
	Name       string
	ShortName  string
	OrganCode  database.NullableString
	Inputcode1 string
	Inputcode2 string
	UsedFlag   int
	Address    database.NullableString
	ZipCode    database.NullableString
	Phone      database.NullableString
	Describe   database.NullableString
	//ModifyEmpid database.NullableInt64
	//ModifyTime time.Time
}
