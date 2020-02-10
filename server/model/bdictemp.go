package model

type BdictEmp struct {
	Id int `db:"ID"`
	Ceid int `db:"CEID"`
	Code string `db:"CODE"`
	Name string `db:"NAME"`
	DeptId int `db:"DEPT_ID"`
	DeptName string `db:"DEPT_NAME"`
	BizDeptId int `db:"BIZ_DEPT_ID"`
	BizDeptName string `db:"BIZ_DEPT_NAME"`
	GroupId int `db:"GROUP_ID"`
	GroupName string `db:"GROUP_NAME"`
	TitlesId int `db:"TITLES_ID"`
	TitlesName string `db:"TITLES_NAME"`
	TakeEmpid int `db:"TAKE_EMPID"`
}


