package models

type User struct {
	Code   string `xorm:"varchar(30) pk 'Code'"`
	Name   string `xorm:"varchar(30) notnull 'Code'"`
	PWD    string `xorm:"varchar(100) notnull 'Code'"`
	Sex    bool
	Dpt    string `xorm:"varchar(30)"`
	Office string `xorm:"varchar(30)"`
	Email  string `xorm:"varchar(30)"`
}
