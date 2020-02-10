package model

import "time"

//SlowRequest 慢查询信息
type SlowRequest struct {
	Reqtime  time.Time `json:"reqtime"`
	URL      string    `json:"url"`
	Duration float64   `json:"duration"`
}
