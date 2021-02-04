package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// MySQLEgn xorm的
var MySQLEgn *xorm.Engine

// InitMySQLEgn 返回xorm引擎
func InitMySQLEgn() *xorm.Engine {
	if MySQLEgn == nil {
		MySQLEgn, _ = xorm.NewEngine("mysql", "root:Lzm114321@(127.0.0.1:33062)/LYBEMS?charset=utf8")
		MySQLEgn.ShowSQL(true)
	}
	return MySQLEgn
}
