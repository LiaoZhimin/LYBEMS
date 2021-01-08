package models

// Position 位置关系
type Position struct {
	ID    int64  `xorm:"id"`
	Name  string `xorm:"varchar(30) notnull 'name'"`
	State int    `xorm:"default(1) 'state'"`
}

// EType 设备类型
type EType struct {
	ID    int64  `xorm:"id"`
	Name  string `xorm:"varchar(30) notnull 'name'"`
	State int    `xorm:"default(1) 'state'"`
}

// Supplier 供应商
type Supplier struct {
	ID      int64  `xorm:"id"`
	Name    string `xorm:"varchar(30) notnull 'name'"`
	Phone   string `xorm:"varchar(30) notnull 'phone'"`
	Address string `xorm:"varchar(300) 'address'"`
	WebSite string `xorm:"varchar(300) 'website'"`
	State   int    `xorm:"default(1) 'state'"`
}

// Equitment 设备信息
type Equitment struct {
	No            int64  `xorm:"pk varchar(30) 'No'"`
	Name          string `xorm:"varchar(30) notnull 'name'"`
	User          string `xorm:"varchar(30) 'user'"`
	UseDate       string `xorm:"varchar(10) 'useDate'"`
	ELiftDayCount string `xorm:"int 'eliftDayCount'"`
	ScrapTime     string `xorm:"datetime 'scrapTime'"`
	Creator       string `xorm:"varchar(10) 'creator'"`
	CreateTime    string `xorm:"datetime 'createTime'"`
	Checker       string `xorm:"varchar(10) 'checker'"`
	CheckResult   string `xorm:"varchar(100) 'checkResult'"`
	CheckTime     string `xorm:"datetime 'checkTime'"`
	State         int    `xorm:"default(1) 'state'"`
	Supplier
	Position
	EType
}
