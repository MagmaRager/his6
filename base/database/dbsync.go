package database

import (
	//"database/sql"
	"errors"
	"github.com/kataras/iris"
	"reflect"
	"strconv"
)

type DBParamValue struct {
	Value string
	Type string
}

//type CdDb struct {
//
//}

var(
	DBList []string	// 满足条件的数据库别名列表
)

/// 初始化DBList
func init() {
	sql := "SELECT ALIAS FROM CD_DB_BRANCH " +
		"WHERE IS_LOCAL = 0 AND STATE = 1"

	var c iris.Context
	err := OraDbCenter.Query(c, &DBList, sql)
	if err != nil {

	}
	if DBList == nil || len(DBList) == 0 {
		errors.New("DBList has no elements!")
	}
}

/// 根据数据库别名获取新的事物ID
func GetTransactionId(alias string) (int, error) {
	tid := -1
	sql := "SELECT MAX(TRANSACTION_ID) FROM CD_AUTO_LOG " +
		"WHERE DB_ALIAS = :1"

	var c iris.Context
	err := OraDbCenter.Find(c, &tid, sql, alias)
	if err != nil {
		return -1, err
	}
	return tid + 1, nil
}

//func GetDB(alias string) {}

/// 事务执行相应语句
func Do(tx Tx, sqlstr string, params ...interface{}) error{
	ps := make([]DBParamValue, len(params))
	i := 0
	for _, dp := range params {
		//if _, ok := dp.(*sql.RawBytes); ok {
		//	return errors.New("sql: RawBytes isn't allowed on Row.Scan")
		//}
		dbv, err := paramToValue(dp)
		if err != nil {
			return err
		}
		ps[i] = dbv
		i++
	}

	var c iris.Context
	_, err := tx.Exec(c, sqlstr, params...)
	if err != nil {
		return err
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

/// 将参数名转化为字段名
func paramToValue(arg interface{}) (DBParamValue, error) {
	t := reflect.TypeOf(arg)
	i := t.Kind()
	dbn := DBParamValue{}

	switch i {
	case 2:
		dbn.Type = "Int"
		dbn.Value = strconv.Itoa(arg.(int))
		break
	case 14:
		dbn.Type = "Float"
		dbn.Value = strconv.FormatFloat(arg.(float64), 'E', -1, 32)
		break
	case 24:
		dbn.Type = "String"
		dbn.Value = arg.(string)
		break
	default:
		dbn.Type = "Unknown"
		break
	}
	return dbn, nil
}