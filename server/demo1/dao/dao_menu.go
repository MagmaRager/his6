package dao

import (
	"errors"
	"his6/base/database"
	"his6/server/model"
)

//MenuDao struct
type MenuDao struct {
}

//GetMenus  获取
func (dao *MenuDao) GetMenus() ([]model.BdictMenu, error) {
	sql := "SELECT MENU_ID, OBJECT_CODE, ORDER_NO, TITLE, WIN_STATE FROM SYS_MENU_WP ORDER BY MENU_ID"
	var menus = make([]model.BdictMenu, 100)
	err := database.OraDb.Query(&menus, sql)
	return menus, err
}

//GetCacheMenus  获取cache
func (dao *MenuDao) GetCacheMenus() ([]model.BdictMenu, error) {
	sql := "SELECT MENU_ID, OBJECT_CODE, ORDER_NO, TITLE, WIN_STATE FROM SYS_MENU_WP ORDER BY MENU_ID"
	var menus = make([]model.BdictMenu, 100)
	err := database.OraDb.QueryWithCache("MenuAll", &menus, sql)
	return menus, err
}

//AddDept 新增科室
func (dao *MenuDao) AddDept(no int, name, loc string) error {
	sql := "INSERT INTO SCOTT.DEPT (DEPTNO, DNAME, LOC) VALUES (:no, :name, :loc)"
	tx, err := database.OraDb.BeginTx()
	if err != nil {
		return err
	}
	rows, err := tx.Exec(sql, no, name, loc)
	if err != nil {
		return err
	}

	if rows == 0 {
		tx.Rollback()
		return errors.New("数据加入没有成功。")
	}

	return tx.Commit()
}
