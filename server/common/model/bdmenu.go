package model

import "his6/base/database"

//BdMenu
type BdMenu struct {
	SystemId   int
	Code       string
	Title      string
	ObjectCode database.NullableString
	Parameter  database.NullableString
	WinState   int
	Ico        database.NullableString
	Prompt     database.NullableString
}
