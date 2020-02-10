package config

import (
	"os"
	"time"
)

//Cfg 配置文件
var Cfg *ConfigFile

func init() {
	configPath := ""

	// 需要配置环境变量HIS6_CONFIG_PATH
	path := os.Getenv("HIS6_CONFIG_PATH")
	if len(path) > 0 {
		configPath = path
	} else {
		cpath, _ := os.Getwd()
		configPath = cpath + "\\conf\\"
	}
	var err error
	Cfg, err = LoadConfigFile(configPath + "\\new_his.conf")
	if err != nil {
		panic("load config error! " + err.Error())
	}
}

//GetConfigString 获取string配置
//@param            section         配置段
//@param            key             键
func GetConfigString(section, key, def string) string {
	value, err := Cfg.GetValue(section, key)
	if err != nil || len(value) == 0 {
		return def
	}
	return value
}

//GetConfigBool 获取bool配置
//@param            section         配置段
//@param            key             键
//@param            def             默认值
func GetConfigBool(section, key string, def bool) bool {
	val, err := Cfg.Bool(section, key)
	if err != nil {
		return def
	}
	return val
}

//GetConfigInt64 获取int64配置
//@param            section         配置段
//@param            key             键
//@param            def             默认值
func GetConfigInt64(section, key string, def int64) int64 {
	val, err := Cfg.Int64(section, key)
	if err != nil {
		return def
	}
	return val
}

//GetConfigInt 获取int配置
//@param            section         配置段
//@param            key             键
//@param            def             默认值
func GetConfigInt(section, key string, def int) int {
	val, err := Cfg.Int(section, key)
	if err != nil {
		return def
	}
	return val
}

//GetConfigFloat64 获取float64配置
//@param            section         配置段
//@param            key             键
//@param            def             默认值
func GetConfigFloat64(section, key string, def float64) float64 {
	val, err := Cfg.Float64(section, key)
	if err != nil {
		return def
	}
	return val
}

//GetConfigDuration 获取duration配置
//@param            section         配置段
//@param            key             键
//@param            def             默认值
func GetConfigDuration(section, key string, def time.Duration) time.Duration {
	value, err := Cfg.GetValue(section, key)
	if err != nil || len(value) == 0 {
		return def
	}

	dv, err := time.ParseDuration(value + "s")
	if err != nil {
		return def
	}

	return dv
}
