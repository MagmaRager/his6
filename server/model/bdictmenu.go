package model

type BdictMenu struct {
	MenuId int `db:"MENU_ID"`
	ObjectCode string `db:"OBJECT_CODE"`
	OrderNo string `db:"ORDER_NO"`
	Title string `db:"TITLE"`
	WinState int `db:"WIN_STATE"`
}
