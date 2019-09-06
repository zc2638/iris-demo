package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/data"
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

// 获取aps当前数量
func (s *ApsService) Count() (count int) {

	db := database.NewDB()
	db.Model(&model.Aps{}).Count(&count)
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

// 获取所有order
func (s *ApsService) GetOrderAll() (all []model.ApsOrder) {

	db := database.NewDB()
	db.Preload("SopProcess").
		Preload("ApsOrderQuality").
		Preload("Aps").
		Order("id DESC").
		Find(&all)
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

// 根据orderID获取order
func (s *ApsService) GetOrderByOrderID(orderID interface{}) (order model.ApsOrder) {

	db := database.NewDB()
	db.Where("order_id = ?", orderID).First(&order)
	return
}

// 根据station和uid获取aps集合
func (s *ApsService) GetOrdersByUidAndStation(uid, station interface{}) (list []model.ApsOrder) {

	db := database.NewDB()
	db.Where("uid = ?", uid).
		Where("station = ?", station).
		Where("status", model.APS_STATUS_START).
		Preload("SopProcess").
		Preload("ApsOrderQuality").
		Preload("Aps").
		Order("id DESC").
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

// 重置aps和order状态
func (s *ApsService) UpdateApsAndOrderToDefaultStatus() {

	db := database.NewDB()

	db.Model(&model.Aps{}).Updates(map[string]interface{}{
		"status": model.APS_STATUS_DEFAULT,
		"sop_id": 0,
	})
	db.Model(&model.ApsOrder{}).Updates(map[string]interface{}{
		"status":         model.APS_STATUS_DEFAULT,
		"sop_process_id": 0,
	})
}

// 批量创建
func (s *ApsService) Insert(apsList []data.PartyAps) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, a := range apsList {
		if a.Order == nil || len(a.Order) == 0 {
			continue
		}

		productService := new(ProductService)
		product := productService.GetProductByName(a.Product)
		if product.ID == 0 {
			continue
		}
		productModel := productService.GetModelByName(product.ID, a.Model)
		if productModel.ID == 0 {
			continue
		}

		craft := new(CraftService).GetCraftByName(a.CraftName)
		if craft.ID == 0 {
			continue
		}

		aps := model.Aps{
			JobPlanNumber: a.JobPlanNumber,
			SerialNo:      a.SerialNo,
			ModelID:       productModel.ID,
			CraftID:       craft.ID,
			PlanTotal:     a.PlanTotal,
			PlanNum:       a.PlanNum,
			CompleteNum:   a.CompleteNum,
		}
		tx.Create(&aps)
		if tx.NewRecord(aps) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}

		for _, o := range a.Order {
			var processID uint
			for _, m := range craft.CraftItem {
				if o.CraftItemName == m.Name {
					processID = m.ID
				}
			}
			if processID == 0 {
				tx.Rollback()
				return errors.New("创建失败")
			}

			user := new(UserService).GetUserByUid(o.Uid)
			if user.ID == 0 {
				tx.Rollback()
				return errors.New("创建失败")
			}

			order := model.ApsOrder{
				OrderID:     o.OrderID,
				ApsID:       aps.ID,
				Uid:         user.ID,
				ProcessID:   processID,
				Station:     o.Station,
				StationName: o.StationName,
				Total:       o.Total,
				Num:         o.Num,
				CompleteNum: o.CompleteNum,
				StartAt:     o.StartAt,
				EndAt:       o.EndAt,
			}
			tx.Create(&order)
			if tx.NewRecord(order) == true {
				tx.Rollback()
				return errors.New("创建失败")
			}
		}
	}
	tx.Commit()
	return nil
}

// 批量创建质检
func (s *ApsService) InsertQuality(qualities []data.PartyQuality) error {

	db := database.NewDB()
	tx := db.Begin()
	for _, q := range qualities {
		order := s.GetOrderByOrderID(q.OrderID)
		if order.ID == 0 {
			continue
		}
		quality := model.ApsOrderQuality{
			OrderID: order.ID,
			PieceNo: q.PieceNo,
			Result:  q.Result,
			Remark:  q.Remark,
		}
		tx.Create(&quality)
		if tx.NewRecord(quality) == true {
			tx.Rollback()
			return errors.New("创建失败")
		}
	}
	tx.Commit()
	return nil
}
