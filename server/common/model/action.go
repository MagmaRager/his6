package model

import (
	"his6/base/database"
)

//Action
type Action struct {
	Code          string
	GrantEmpid    database.NullableInt64
	EffectiveTime database.NullableString
	ExpiryTime    database.NullableString
}
