package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

// MyXorm 封装XORM操作对象，其他struct DAO用于继承
type MyXorm struct {
	Msg string `json:"msg"`
}

// MyXormInit 返回xorm引擎 
func MyXormInit() *xorm.Engine {
	if engine == nil {
		engine, _ = xorm.NewEngine("mysql", "root:Lzm114321@tcp(127.0.0.1:33061)/LYBEMS?charset=utf8")
	}
	return engine
}

//GetById 实例DAO函数，根据ID获取数据
func (mx *MyXorm) GetById(int64 id) rt interface{}{
	engine.Id(id).Get(&rt)
	return rt
}
