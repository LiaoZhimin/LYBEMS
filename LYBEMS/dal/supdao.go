package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// SUPDAO 设备信息DAO
type SUPDAO struct {
}

// GetList 获取设备信息
func (ud *SUPDAO) GetList() (*[]models.Supplier, error) {
	var ls []models.Supplier
	eng := services.InitMySQLEgn()
	err := eng.Desc("id").Limit(1000, 0).Find(&ls)
	return &ls, err
}

// GetListBy 条件查询
func (ud *SUPDAO) GetListBy(name string) (*[]models.Supplier, error) {
	var ls []models.Supplier
	eng := services.InitMySQLEgn()
	err := eng.Desc("id").Where("name like ?+'%'", name).And("state>0").Limit(1000, 0).Find(&ls)
	return &ls, err
}

//InOrUp 更新数据，存在修改，否在新增
func (ud *SUPDAO) InOrUp(m1 models.Supplier) (int64, error) {
	eng := services.InitMySQLEgn()
	b, _ := eng.Cols("id").ID(m1.ID).Exist()
	if b {
		return eng.ID(m1.ID).Update(&m1)
	} else {
		return eng.InsertOne(&m1)
	}
}
