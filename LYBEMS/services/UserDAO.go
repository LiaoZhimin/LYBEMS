package main

import (
	mm "../models"
)

//UserDAO 用户DAO
type UserDAO struct {
	Msg string
	MyXorm
}

// List 获取用户列表
func (ud *UserDAO) List(mm.User) []*mm.User {

}
