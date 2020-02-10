package cache

import (
	"encoding/json"
	"his6/base/config"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"

	"errors"
	"time"
)

var (
	//cacheType string
	cacheConn       cache.Cache
	timeoutDuration time.Duration
)

//  初始化redis数据库连接池
func init() {
	rk := config.GetConfigString("cache", "redis_key", "NewHisCache")
	rc := config.GetConfigString("cache", "redis_url", "")

	if len(rk) == 0 || len(rc) == 0 {
		memoryInit()
		return
	}
	mstr := map[string]string{}

	mstr["key"] = rk
	mstr["conn"] = rc
	mstr["dbNum"] = "0"
	//TODO more parm?

	bytes, _ := json.Marshal(mstr)

	//从redis缓存中拿数据拿数据
	cache_conn, err := cache.NewCache("redis", string(bytes))
	if err != nil {
		logs.Error("Redis is not connected, switched to memory cache.")
		memoryInit()
		return
	}

	cacheConn = cache_conn
	timeoutDuration = config.GetConfigDuration("cache", "timeout", 86400*time.Second)
}

func memoryInit() {
	mstr := map[string]string{}

	mstr["interval"] = "60"
	bytes, _ := json.Marshal(mstr)
	cacheConn, _ = cache.NewCache("memory", string(bytes))
}

//  判别key是否存在
func ExistKey(k string) (exist bool) {
	return cacheConn.IsExist(k)
}

//  删除key
func DeleteKey(k string) error {
	return cacheConn.Delete(k)
}

//  设置对象数据用Json保存value
func SetData(k string, data interface{}) error {
	value, _ := json.Marshal(data)
	err := cacheConn.Put(k, value, timeoutDuration)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetData(k string, rst interface{}) error {
	if areaData := cacheConn.Get(k); areaData != nil {
		if bts, ok := areaData.([]byte); ok {
			err := json.Unmarshal(bts, &rst)
			return err
		}
		var err = errors.New("Cache value not convertable")
		return err

	} else {
		var err = errors.New("No cache data")
		return err
	}
	return nil
}
