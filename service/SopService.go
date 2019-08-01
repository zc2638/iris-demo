package service

import (
	"github.com/kataras/iris/core/errors"
	"sop/lib/database"
	"sop/model"
	"strings"
	"time"
)

type SopService struct{ BaseService }

// 获取sop列表
func (s *SopService) GetList(page, pageSize int) (list []model.Sop) {

	db := database.NewDB()
	s.Paginate(db, page, pageSize).
		Where("is_template = ?", 0).
		Preload("SopModel").
		Preload("Product").
		Preload("Craft").
		Find(&list)
	return
}

// 获取所有sop
func (s *SopService) GetAll(isAll bool) (list []model.Sop) {

	db := database.NewDB()
	if isAll == false {
		db = db.Where("is_template = ?", 0)
	}
	db.Preload("SopModel").
		Preload("SopProcess").
		Preload("Product").
		Preload("Craft").
		Find(&list)
	return
}

// 根据工艺方案id获取所有sop模板
func (s *SopService) GetTemplateAllByCraftID(id interface{}) (all []model.Sop) {

	db := database.NewDB()
	db.Where("is_template = ?", 1).Where("craft_id = ?", id).Preload("SopProcess").Find(&all)
	return
}

// 根据id获取sop信息
func (s *SopService) GetSopByID(id interface{}) (sop model.Sop) {

	db := database.NewDB()
	db.Where("id = ?", id).
		Preload("SopProcess").
		Preload("SopModel").
		Preload("Product").
		Preload("Craft").
		First(&sop)
	return
}

// 根据id获取多个sop信息
func (s *SopService) GetSopListByIds(ids interface{}) (list []model.Sop) {

	db := database.NewDB()
	db.Where("id in (?)", ids).Preload("SopProcess").Find(&list)
	return
}

// 根据productId获取sop
func (s *SopService) GetSopListByProductID(productID interface{}) (list []model.Sop) {

	db := database.NewDB()
	db.Where("product_id = ?", productID).
		Where("is_template = ?", 0).
		Preload("Product").
		Preload("SopModel").
		Preload("SopProcess").
		Find(&list)
	return
}

// 更新sop
func (s *SopService) UpdateOne(sop model.Sop, processes []model.SopProcess, models []model.SopModel) error {

	db := database.NewDB()
	tx := db.Begin()

	if tx.Save(&sop).RowsAffected != 1 {
		tx.Rollback()
		return errors.New("sop更新失败")
	}
	if processes != nil && len(processes) > 0 {
		for _, p := range processes {
			if tx.Save(&p).RowsAffected != 1 {
				tx.Rollback()
				return errors.New("sop工序更新失败")
			}
		}
	}
	if models != nil && len(models) > 0 {
		tx.Where(model.SopModel{SopID: sop.ID}).Delete(model.SopModel{})
		for _, m := range models {
			tx.Create(&m)
			if tx.NewRecord(m) == true {
				tx.Rollback()
				return errors.New("sop产品型号更新失败")
			}
		}
	}

	tx.Commit()
	return nil
}

// 批量更新sop
func (s *SopService) Updates(sopList []model.Sop) error {

	db := database.NewDB()
	tx := db.Begin()

	for _, s := range sopList {
		if tx.Save(&s).RowsAffected != 1 {
			tx.Rollback()
			return errors.New("更新失败")
		}
	}

	tx.Commit()
	return nil
}

// 创建sop
func (s *SopService) Create(sop model.Sop, process []model.SopProcess, models []model.SopModel) error {

	now := time.Now()
	db := database.NewDB()
	tx := db.Begin()

	tx.Create(&sop)

	if tx.NewRecord(sop) == true {
		tx.Rollback()
		return errors.New("sop插入失败")
	}

	processNum := len(process)
	var processNumSet []string
	for i := 0; i < processNum; i++ {
		processNumSet = append(processNumSet, "(?)")
	}

	var processExecSet = make([]interface{}, 0)
	for _, p := range process {
		processExecSet = append(processExecSet, []interface{}{sop.ID, p.ProcessID, p.Title, p.Content, p.Imgs, p.Annex, p.CheckImg, p.MinefieldImg, p.IsCheck, p.Sort, now, now})
	}

	res := tx.Exec(
		"INSERT INTO `sop_processes` (`sop_id`, `process_id`, `title`, `content`, `imgs`, `annex`, `check_img`, `minefield_img`, `is_check`, `sort`, `created_at`, `updated_at`) VALUES "+strings.Join(processNumSet, ", "),
		processExecSet...
	)
	if res.RowsAffected != int64(processNum) {
		tx.Rollback()
		return errors.New("sop_process插入失败")
	}

	modelNum := len(models)
	var modelNumSet []string
	for i := 0; i < modelNum; i++ {
		modelNumSet = append(modelNumSet, "(?)")
	}

	var modelExecSet = make([]interface{}, 0)
	for _, m := range models {
		modelExecSet = append(modelExecSet, []interface{}{sop.ID, m.ModelID, m.Name, now, now})
	}

	res = tx.Exec(
		"INSERT INTO `sop_models` (`sop_id`, `model_id`, `name`, `created_at`, `updated_at`) VALUES "+strings.Join(modelNumSet, ", "),
		modelExecSet...
	)
	if res.RowsAffected != int64(modelNum) {
		tx.Rollback()
		return errors.New("sop_model插入失败")
	}

	tx.Commit()
	return nil
}

// 根据sop id获取所有process
func (s *SopService) GetProcessBySopID(sopID interface{}) (all []model.SopProcess) {

	db := database.NewDB()
	db.Where("sop_id = ?", sopID).Find(&all)
	return
}

// 根据id获取process
func (s *SopService) GetProcessByID(id interface{}) (process model.SopProcess) {

	db := database.NewDB()
	db.Where("id = ?", id).First(&process)
	return
}

// 更新process
func (s *SopService) UpdateProcessOne(process model.SopProcess) error {

	db := database.NewDB()
	if db.Save(&process).RowsAffected != 1 {
		return errors.New("更新失败")
	}
	return nil
}