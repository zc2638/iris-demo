package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/lib/database"
	"sop/model"
)

type ApsService struct{ BaseService }

// 获取所有aps
func (s *ApsService) GetAll() (all []model.Aps) {

	db := database.NewDB()
	db.Preload("ProductModel").Find(&all)
	return
}

// 根据产品型号id获取aps
func (s *ApsService) GetListByModelIds(ids interface{}) (list []model.Aps) {

	db := database.NewDB()
	db.Preload("ProductModel").Where("model_id in (?)", ids).Find(&list)
	return
}

// 根据id获取aps
func (s *ApsService) GetApsByID(id interface{}) (aps model.Aps) {

	db := database.NewDB()
	db.Where("id = ?", id).Preload("ProductModel").Preload("ApsOrder").First(&aps)
	return
}

// 根据id获取aps集合
func (s *ApsService) GetApsListByIds(ids interface{}) (list []model.Aps) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Preload("ApsOrder").Preload("ProductModel").Find(&list)
	return
}

// 根据id获取order
func (s *ApsService) GetOrderByID(id interface{}) (order model.ApsOrder) {

	db := database.NewDB()
	db.Where("id = ?", id).
		Preload("SopProcess").
		Preload("Aps").
		Preload("ApsOrderQuality").
		First(&order)
	return
}

// 根据station和uid获取aps集合
func (s *ApsService) GetOrdersByUidAndStation(uid, station interface{}) (list []model.ApsOrder) {

	db := database.NewDB()
	db.Where("uid = ?", uid).
		Where("station = ?", station).
		//Where("status", model.APS_STATUS_START).
		Preload("SopProcess").
		Preload("ApsOrderQuality").
		Preload("Aps").
		Limit(20).
		Find(&list)
	return
}

// 更新aps和order
func (s *ApsService) UpdateApsAndOrder(apsList []model.Aps, orders []model.ApsOrder) error {

	db := database.NewDB()
	tx := db.Begin()

	for _, aps := range apsList {
		if tx.Save(&aps).RowsAffected != 1 {
			tx.Rollback()
			return errors.New("aps更新失败")
		}
	}

	for _, order := range orders {
		if tx.Save(&order).RowsAffected != 1 {
			tx.Rollback()
			return errors.New("aps_order更新失败")
		}
	}

	tx.Commit()
	return nil
}