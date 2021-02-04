package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// UserDAO 用户信息DAO
type UserDAO struct {
}

// GetList 获取用户信息
func (ud *UserDAO) GetList() (*[]models.User, error) {
	var users []models.User
	eng := services.InitMySQLEgn()
	err := eng.Desc("Code").Limit(100, 0).Find(&users)
	return &users, err
}
