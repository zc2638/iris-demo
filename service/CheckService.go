package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/lib/database"
	"sop/model"
)

type CheckService struct{ BaseService }

// 根据url获取单条数据
func (s *CheckService) GetCheckByUrl(url interface{}) (check model.Check) {

	db := database.NewDB()
	db.Where("url = ?", url).First(&check)
	return
}

// 更新
func (s *CheckService) UpdateOne(check model.Check) error {

	db := database.NewDB()
	if db.Save(&check).RowsAffected != 1 {
		return errors.New("操作失败")
	}
	return nil
}