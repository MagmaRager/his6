package model

import (
	"his6/base/database"
)

//CdObject 模块
type CdObject struct {
	Code              string
	Name              string
	Object            string
	ModuleCode        string
	UsedFlag          int
	IsFunction        int
	HasFunctionPoint  int
	Describe          database.NullableString
	FunctionPointList []CdObjectFp
}
