package database

//SlowSQL 慢Sql语句信息
type SQLMonitor struct {
	Node       string   `json:"node"`
	Time       string   `json:"time"`
	ExecuteSQL string   `json:"executesql"`
	Parameters []string `json:"parameters"`
	Duration   float64  `json:"duration"`
}
