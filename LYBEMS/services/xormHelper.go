package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// MySQLEgn xorm的
var MySQLEgn *xorm.Engine

// MyXorm 封装XORM操作对象，其他struct DAO用于继承
type MyXorm struct {
	Msg string `json:"msg"`
}

// InitMySQLEgn 返回xorm引擎
func (mx *MyXorm) InitMySQLEgn() *xorm.Engine {
	if MySQLEgn == nil {
		MySQLEgn, _ = xorm.NewEngine("mysql", "root:Lzm114321@(127.0.0.1:33062)/LYBEMS?charset=utf8")
		MySQLEgn.ShowSQL(true)
	}
	return MySQLEgn
}
