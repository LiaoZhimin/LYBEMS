package models

// User 用户
type User struct {
	Code   string `xorm:"varchar(30) pk 'Code'"`
	Name   string `xorm:"varchar(30) notnull 'Name'"`
	PWD    string `xorm:"varchar(100) notnull 'PWD'"`
	Sex    bool   `xorm:"bit 'Sex'"`
	Dpt    string `xorm:"varchar(30) 'Dpt'"`
	Office string `xorm:"varchar(30) 'Office'"`
	Email  string `xorm:"varchar(30) 'Email'"`
}

// UserLimits 用户权限
type UserLimits struct {
	ID       int64  `xorm:"id"`
	UserCode string `xorm:"varchar(30) notnull 'UserCode'`
	Limits   string `xorm:"varchar(max) 'Limits'"`
}
