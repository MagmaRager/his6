package model

import (
	"his6/base/database"
)

//CdModule 模块
type CdModule struct {
	Code        string
	Name        string
	FileName    string
	Version     string
	UsedFlag    int
	Describe    database.NullableString
	FileTime    database.NullableString
	UpdateEmpid database.NullableInt64
	//UpdateTime  time.Time
}
