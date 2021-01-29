package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// UserDAO 用户信息DAO
type UserDAO struct {
}

// GetList 获取用户信息
func (ud *UserDAO) GetList() *[]models.User {
	var users []models.User
	var myx services.MyXorm
	eng := myx.InitMySQLEgn()
	eng.Desc("Code").Limit(100, 0).Find(&users)
	return &users
}
