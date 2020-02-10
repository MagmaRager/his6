package database

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"his6/base/convert"
	"his6/base/crypto"
	"his6/base/message"
	"his6/base/middle/model"
	"os"
	"strconv"
	"strings"
	"time"

	"his6/base/cache"
	"his6/base/config"

	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/golog"

	//oci Oracle处理
	_ "github.com/mattn/go-oci8"
)

var (
	//OraDb ： 生产库对象
	OraDb OracleDb = NewDb("db")
	//OraDbCenter ： 中心库对象
	OraDbCenter OracleDb = NewDb("dbc")
	// //OraDbHistory ： 历史库对象对象
	// OraDbHistory OracleDb = NewDb("dbh")

	//longSQLTime ：慢SQL时间，默认为3s，若超过则发送消息
	longSQLTime float64 = config.GetConfigDuration("logs", "long_sql_time",
		time.Duration(3*time.Second)).Seconds()
)

//OracleDb ODb结构
type OracleDb struct {
	db *sqlx.DB
}

//Tx 事务结构
type Tx struct {
	sqltx *sql.Tx
}

//NewDb 创建数据库实体
func NewDb(typ string) OracleDb {
	nlsLang := config.GetConfigString(typ, "nls_lang", "") //"AMERICAN_AMERICA.AL32UTF8"
	os.Setenv("NLS_LANG", nlsLang)

	sid := config.GetConfigString(typ, "sid", "")
	user := config.GetConfigString(typ, "user", "")
	password := config.GetConfigString(typ, "password", "wdhis2018")

	key := config.GetConfigString(typ, "key", "NHIS2020")
	if len(user) == 0 || len(sid) == 0 {
		golog.Error(typ + " 数据库服务地址没有配置")
		return OracleDb{}
	}
	// crypto.EncryptDES_CBC("", key)
	pd := crypto.DecryptDES_CBC(password, key)

	url := user + "/" + pd + "@" + sid
	db, _ := sqlx.Open("oci8", url)

	maxConn, _ := strconv.Atoi(config.GetConfigString(typ, "maxOpenConns", "25"))
	minConn, _ := strconv.Atoi(config.GetConfigString(typ, "maxIdleConns", "5"))
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(minConn)
	db.MapperFunc(PascalToCol)

	return OracleDb{db}
}

// func PascalToColx(dbColumnName string) string {
// 	//fmt.Println("x1:" + time.Now().String())
// 	var cols = ""
// 	var i int
// 	for i = 0; i < len(dbColumnName); i++ {
// 		var c = dbColumnName[i]
// 		if i != 0 && c >= 'A' && c <= 'Z' {
// 			cols += "_"
// 		}
// 		cols += string(dbColumnName[i])
// 	}
// 	str := strings.ToUpper(cols)
// 	//fmt.Println("x2:" + time.Now().String())
// 	return str
// }

//PascalToCol PascalCase参数格式转化为数据库字段（下划线）
func PascalToCol(dbColumnName string) string {
	//fmt.Println("x1:" + time.Now().String())
	var cols bytes.Buffer
	cols.Grow(len(dbColumnName))
	var i int
	for i = 0; i < len(dbColumnName); i++ {
		var c = dbColumnName[i]
		if i != 0 && c >= 'A' && c <= 'Z' {
			cols.WriteByte('_')
		}
		cols.WriteByte(c)
	}
	str := strings.ToUpper(cols.String())
	//fmt.Println("x2:" + time.Now().String())
	return str
}

//BeginTx 开始事务
func (oracle OracleDb) BeginTx() (*Tx, error) {
	tx, err := oracle.db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{sqltx: tx}, nil
}

//Commit 提交
func (tx *Tx) Commit() error {
	return tx.sqltx.Commit()
}

//Rollback 回滚
func (tx *Tx) Rollback() error {
	return tx.sqltx.Rollback()
}

//Exec 执行返回影响记录数
func (tx *Tx) Exec(sql string, args ...interface{}) (int64, error) {
	st := time.Now()
	rst, err := tx.sqltx.Exec(sql, args...)

	if err != nil {
		tx.Rollback()
		return -1, err
	}
	et := time.Now()
	bt := et.Sub(st).Seconds()
	if bt >= longSQLTime {
		// 超过时间，记录慢SQL
		var slowsql = model.SlowSQL{}
		slowsql.Time = et
		slowsql.ExecuteSQL = sql
		var parms []string
		for i, arg := range args {
			str := arg.(string)
			parms[i] = str
		}
		slowsql.Params = parms
		slowsql.Duration = bt

		cvb, err := json.Marshal(slowsql)
		if err != nil {
			logs.Error("慢查询json转化失败")
		}
		cvs := convert.Byte2Str(cvb)
		message.Send("slowsql", cvs)
	}

	return rst.RowsAffected()
}



//Find 查询单行结果, 返回错误
func (oracle OracleDb) Find(rst interface{}, sql string, args ...interface{}) error {
	return sqlx.Get(oracle.db, rst, sql, args...)
}

//Query 查询多行结果，返回错误
func (oracle OracleDb) Query(rst interface{}, sql string, args ...interface{}) error {
	return sqlx.Select(oracle.db, rst, sql, args...)
}

//QueryWithCache 查询多行结果（优先从缓存获取），返回错误
func (oracle OracleDb) QueryWithCache(key string, rst interface{}, sql string, args ...interface{}) error {
	err := cache.GetData(key, rst)
	if err != nil {
		// 查询与结果遍历
		err = sqlx.Select(oracle.db, rst, sql, args...)
		if err != nil {
			return err
		}
		cache.SetData(key, rst)
	}

	return nil
}
