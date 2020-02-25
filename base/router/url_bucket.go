package router

import "sync"

//UrlBucket url监视信息元素
type UrlBucket struct {
	sync.Mutex

	url string //相对url
	method string // HTTP请求类型
	handler string // 包名

	sum float64	//累计调用时长
	count [9]int // 时间范围计数（秒）： .01, .025, .05, .1, .25, .5, 1, 3, Inf

	failCount int //错误计数
}

