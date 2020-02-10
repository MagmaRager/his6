package consul

import (
	"fmt"
	"his6/base/config"
	"log"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

// func consulCheck(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "consulCheck")
// }

func init() {
	regName := config.GetConfigString("app", "name", "")
	port := config.GetConfigInt("app", "port", 8079)

	consulAddr := config.GetConfigString("consul", "url", "") //consul server
	regAddr := config.GetConfigString("consul", "regAddress", "")
	timeout := config.GetConfigString("consul", "checkTimeout", "")
	interval := config.GetConfigString("consul", "checkInterval", "")
	deregTime := config.GetConfigString("consul", "checkDereg", "")

	configcs := consulapi.DefaultConfig()
	configcs.Address = consulAddr
	client, err := consulapi.NewClient(configcs)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = regName + "_" + regAddr + ":" + strconv.Itoa(port)
	registration.Name = regName
	registration.Port = port
	registration.Tags = []string{regName}
	registration.Address = regAddr
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, port, "/check"),
		Timeout:                        timeout,
		Interval:                       interval,
		DeregisterCriticalServiceAfter: deregTime, //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	//go runCheck()

}

// func runCheck() { //反向访问check
// 	http.HandleFunc("/check", consulCheck)
// 	http.ListenAndServe(fmt.Sprintf(":%d", 8079), nil)
// }
