package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// EquitmentDAO 设备信息DAO
type EquitmentDAO struct {
}

// GetList 获取设备信息
func (ud *EquitmentDAO) GetList() (*[]models.Equitment, error) {
	var ls []models.Equitment
	eng := services.InitMySQLEgn()
	err := eng.Desc("No").Limit(1000, 0).Find(&ls)
	return &ls, err
}

// GetListBy 条件查询
func (ud *EquitmentDAO) GetListBy(no string, name string) (*[]models.Equitment, error) {
	var ls []models.Equitment
	eng := services.InitMySQLEgn()
	err := eng.Desc("No").Where("No like ?+'%'", no).And("name like ?+'%'", name).And("state>0").Limit(1000, 0).Find(&ls)
	return &ls, err
}

//InOrUp 更新数据，存在修改，否在新增
func (ud *EquitmentDAO) InOrUp(m1 models.Equitment) (int64, error) {
	eng := services.InitMySQLEgn()
	b, _ := eng.Cols("No").ID(m1.No).Exist()
	if b {
		return eng.ID(m1.No).Update(&m1)
	}
	return eng.InsertOne(&m1)

}
