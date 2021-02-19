package dal

import (
	"LYBEMS/LYBEMS/models"
	"LYBEMS/LYBEMS/services"
)

// PSTDAO 设备信息DAO
type PSTDAO struct {
}

// GetList 获取设备信息
func (ud *PSTDAO) GetList() (*[]models.Position, error) {
	var ls []models.Position
	eng := services.InitMySQLEgn()
	err := eng.Desc("No").Limit(1000, 0).Find(&ls)
	return &ls, err
}

// GetListBy 条件查询
func (ud *PSTDAO) GetListBy(name string) (*[]models.Position, error) {
	var ls []models.Position
	eng := services.InitMySQLEgn()
	err := eng.Desc("id").Where("name like ?+'%'", name).And("state>0").Limit(1000, 0).Find(&ls)
	return &ls, err
}

//InOrUp 更新数据，存在修改，否在新增
func (ud *PSTDAO) InOrUp(m1 models.Position) (int64, error) {
	eng := services.InitMySQLEgn()
	b, _ := eng.Cols("id").ID(m1.ID).Exist()
	if b {
		return eng.ID(m1.ID).Update(&m1)
	}
	return eng.InsertOne(&m1)

}
