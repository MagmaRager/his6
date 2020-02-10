package model

import "time"

//SlowSQL 慢Sql语句信息
type SlowSQL struct {
	Time       time.Time `json:"time"`
	ExecuteSQL string    `json:"executesql"`
	Params     []string  `json:"params"`
	Duration   float64   `json:"duration"`
}
