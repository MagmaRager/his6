package model

import (
	"his6/base/database"
)

//BdSystem
type BdSystem struct {
	Id   int
	Code string
	Name string
	Ico  database.NullableString
}
