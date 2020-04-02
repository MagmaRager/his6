package main

import (
	"his6/base/router"

	//_ "his6/base/consul"
	_ "his6/base/message"
	_ "his6/base/middle" //PS: 仅写在middle后面的模块受middle影响
	_ "his6/server/common/handler"
	_ "his6/server/metrics"
)

func main() {
	router.Run()
}
