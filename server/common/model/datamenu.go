package model

import (
	"his6/base/database"
)

//DataMenu
type DataMenu struct {
	Code       string
	Title      string
	ModuleName database.NullableString
	ObjectName database.NullableString
	Parameter  database.NullableString
	WinState   int
	Ico        database.NullableString
	Prompt     database.NullableString
}
