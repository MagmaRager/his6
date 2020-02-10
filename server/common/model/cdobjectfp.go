package model

import "his6/base/database"

//CdObjectFp 模块
type CdObjectFp struct {
	Code       string
	Name       string
	ObjectCode string
	Describe   database.NullableString
}
