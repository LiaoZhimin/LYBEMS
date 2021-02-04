package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// ETypeDAO 设备信息DAO
type ETypeDAO struct {
}

// GetList 获取设备信息
func (ud *ETypeDAO) GetList() (*[]models.EType, error) {
	var ls []models.EType
	eng := services.InitMySQLEgn()
	err := eng.Desc("No").Limit(1000, 0).Find(&ls)
	return &ls, err
}

// GetListBy 条件查询
func (ud *ETypeDAO) GetListBy(name string) (*[]models.EType, error) {
	var ls []models.EType
	eng := services.InitMySQLEgn()
	err := eng.Desc("id").Where("name like ?+'%'", name).And("state>0").Limit(1000, 0).Find(&ls)
	return &ls, err
}

//InOrUp 更新数据，存在修改，否在新增
func (ud *ETypeDAO) InOrUp(m1 models.EType) (int64, error) {
	eng := services.InitMySQLEgn()
	b, _ := eng.Cols("id").ID(m1.ID).Exist()
	if b {
		return eng.ID(m1.ID).Update(&m1)
	} else {
		return eng.InsertOne(&m1)
	}
}
