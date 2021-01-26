package services

import (
	"github.com/xormplus/xorm"
)

// Egn xorm的
var Egn *xorm.Engine

// MyXorm 封装XORM操作对象，其他struct DAO用于继承
type MyXorm struct {
	Msg string `json:"msg"`
}

// MyXormInit 返回xorm引擎
func (mx *MyXorm) MyXormInit() *xorm.Engine {
	if Egn == nil {
		Egn, _ = xorm.NewEngine("mysql", "root:Lzm114321@tcp(127.0.0.1:33062)/LYBEMS?charset=utf8")
	}
	return Egn
}
