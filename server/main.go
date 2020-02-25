package main

import (
	//"his6/base/eureka"
	"his6/base/router"

	//_ "his6/base/middle" //PS: 仅写在middle后面的模块受middle影响
	_ "his6/base/consul"
	_ "his6/server/common/handler"
	_ "his6/server/metrics"
	//_ "his6/server/demo"
	//_ "his6/server/demo1"
)

func main() {

	//region eureka register

	// client := eureka.NewClient(&eureka.Config{
	// 	DefaultZone:           "http://10.240.35.14:8761/eureka/",
	// 	App:                   "GoHis",
	// 	Port:                  8079,
	// 	RenewalIntervalInSecs: 10,
	// 	DurationInSecs:        30,
	// 	Metadata: map[string]interface{}{
	// 		"VERSION":              "0.1.0",
	// 		"NODE_GROUP_ID":        0,
	// 		"PRODUCT_CODE":         "DEFAULT",
	// 		"PRODUCT_VERSION_CODE": "DEFAULT",
	// 		"PRODUCT_ENV_CODE":     "DEFAULT",
	// 		"SERVICE_VERSION_CODE": "DEFAULT",
	// 	},
	// })
	// // start client, register、heartbeat、refresh
	// client.Start()

	// // http server
	// http.HandleFunc("/v1/services", func(writer http.ResponseWriter, request *http.Request) {
	// 	// full applications from eureka server
	// 	apps := client.Applications

	// 	b, _ := json.Marshal(apps)
	// 	_, _ = writer.Write(b)
	// })
	//endregion

	router.Run()
}
