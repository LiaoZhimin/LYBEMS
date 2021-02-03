package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// EquitmentDAO 设备信息DAO
type EquitmentDAO struct {
}

// GetList 获取用户信息
func (ud *EquitmentDAO) GetList() *[]models.Equitment {
	var equits []models.Equitment
	var myx services.MyXorm
	eng := myx.InitMySQLEgn()
	eng.Desc("No").Limit(100, 0).Find(&equits)
	return &equits
}
