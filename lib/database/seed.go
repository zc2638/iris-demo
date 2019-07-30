package database

import (
	"encoding/json"
	"sop/model"
	"time"
)

func seed() {

	seedCraft()
	seedProduct()
	seedSop()
	seedApsOrder()
	seedAndon()
	seedUser()
	seedCheck()
	seedOrderQuality()
}

type checkImgData struct {
	Item string `json:"item"`
	Bg   string `json:"bg"`
}

var checkImg = checkImgData{
	"http://puh01tec3.bkt.clouddn.com/1564121305087721-m1.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564122609174237-6181564122551_.pic_hd.jpg",
}
var minefieldImg = []string{
	"http://puh01tec3.bkt.clouddn.com/1564121305087721-m1.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564121681493909-6051564121588_.pic.jpg",
}

// 工艺方案填充
func seedCraft() {

	checkImgJson, _ := json.Marshal(checkImg)
	minefieldImgJson, _ := json.Marshal(minefieldImg)

	now := time.Now()
	db := NewDB()

	var craft model.Craft
	db.First(&craft)
	if craft.ID == 0 {
		db.Exec(
			"INSERT INTO `crafts` (`name`, `status`, `created_at`, `updated_at`) VALUES (?)",
			[]interface{}{"#01鼠标外壳粗加工", 0, now, now},
		)

		var crafts []model.Craft
		db.Find(&crafts)
		if crafts != nil && len(crafts) > 0 {
			for _, c := range crafts {
				if c.Name == "#01鼠标外壳粗加工" {
					db.Exec(
						"INSERT INTO `craft_items` (`craft_id`, `name`, `check_img`, `minefield_img`, `sort`, `status`, `created_at`, `updated_at`) VALUES (?)",
						[]interface{}{c.ID, "通孔", checkImgJson, minefieldImgJson, 1, 0, now, now},
					)
				}
			}
		}
	}

}

// 产品信息填充
func seedProduct() {

	now := time.Now()
	db := NewDB()

	var product model.Product
	db.First(&product)
	if product.ID == 0 {
		db.Exec(
			"INSERT INTO `products` (`name`, `created_at`, `updated_at`) VALUES (?)",
			[]interface{}{"鼠标外壳", now, now},
		)

		var products []model.Product
		db.Find(&products)
		if products != nil && len(products) > 0 {
			for _, p := range products {
				if p.Name == "鼠标外壳" {
					db.Exec(
						"INSERT INTO `product_models` (`product_id`, `name`, `created_at`, `updated_at`) VALUES (?)",
						[]interface{}{p.ID, "#78", now, now},
					)
				}
			}
		}
	}
}

// sop数据填充
func seedSop() {

	now := time.Now()
	db := NewDB()

	var sop model.Sop
	db.First(&sop)
	if sop.ID == 0 {

		var crafts []model.Craft
		db.Find(&crafts)

		var products []model.Product
		db.Find(&products)

		if crafts != nil && products != nil && len(crafts) > 0 && len(products) > 0 {
			for _, c := range crafts {
				for _, p := range products {
					if p.Name == "鼠标外壳" {
						db.Exec(
							"INSERT INTO `sops` (`title`, `craft_id`, `product_id`, `is_template`, `version`, `status`, `created_at`, `updated_at`) VALUES (?)",
							[]interface{}{"#78鼠标外壳作业指导书", c.ID, p.ID, 1, 1, 0, now, now},
						)
					}
				}
			}
		}

		var sops []model.Sop
		db.Find(&sops)
		if sops != nil && len(sops) > 0 {

			checkImgJson, _ := json.Marshal(checkImg)
			minefieldImgJson, _ := json.Marshal(minefieldImg)

			for _, s := range sops {

				var items []model.CraftItem
				db.Find(&items)

				if items != nil && len(items) > 0 {
					for _, t := range items {
						db.Exec(
							"INSERT INTO `sop_processes` (`sop_id`, `process_id`, `title`, `check_img`, `minefield_img`, `is_check`, `sort`, `created_at`, `updated_at`) VALUES (?)",
							[]interface{}{s.ID, t.ID, "通孔", checkImgJson, minefieldImgJson, 0, t.Sort, now, now},
						)
					}
				}

				var models []model.ProductModel
				db.Where("product_id = ?", s.ProductID).Find(&models)
				if models != nil && len(models) > 0 {
					for _, m := range models {
						db.Exec(
							"INSERT INTO `sop_models` (`sop_id`, `model_id`, `name`, `created_at`, `updated_at`) VALUES (?)",
							[]interface{}{s.ID, m.ID, m.Name, now, now},
						)
					}
				}
			}
		}
	}
}

// 作业计划—工单 数据填充
func seedApsOrder() {

	now := time.Now()
	db := NewDB()

	var aps model.Aps
	db.First(&aps)
	if aps.ID == 0 {

		var models []model.ProductModel
		db.Find(&models)

		var crafts []model.Craft
		db.Find(&crafts)

		if models != nil && len(models) > 0 && crafts != nil && len(crafts) > 0 {
			for _, m := range models {
				for _, c := range crafts {
					db.Exec(
						"INSERT INTO `aps` (`job_plan_number`, `serial_no`, `model_id`, `craft_id`, `plan_total`, `plan_num`, `complete_num`, `status`, `created_at`, `updated_at`) VALUES (?)",
						[]interface{}{"WP201987771341", "line-m001", m.ID, c.ID, 100, 100, 0, 0, now, now},
					)
				}
			}

			var apsList []model.Aps
			db.Find(&apsList)

			if apsList != nil && len(apsList) > 0 {
				for _, a := range apsList {

					var craftItems []model.CraftItem
					db.Where("craft_id = ?", a.CraftID).Find(&craftItems)
					if craftItems != nil && len(craftItems) > 0 {
						for _, c := range craftItems {
							if a.JobPlanNumber == "WP201987771341" {
								db.Exec(
									"INSERT INTO `aps_orders` (`order_id`, `aps_id`, `uid`, `process_id`, `station`, `station_name`, `total`, `num`, `complete_num`, `status`, `created_at`, `updated_at`) VALUES (?)",
									[]interface{}{"WO1987701242", a.ID, 1, c.ID, 1, "通孔工位", 100, 100, 0, 0, now, now},
								)
							}
						}
					}
				}
			}
		}
	}
}

// andon数据填充
func seedAndon() {

	now := time.Now()
	db := NewDB()

	var andon model.Andon
	db.First(&andon)
	if andon.ID == 0 {
		db.Exec(
			"INSERT INTO `andons` (`type`, `content`, `station`, `code`, `information`, `trigger_time`, `created_at`, `updated_at`) VALUES (?), (?)",
			[]interface{}{"警报", "测试andon内容1", "1", "001", "andon测试业务信息", now, now, now},
			[]interface{}{"提示", "测试andon内容2", "1", "002", "andon测试业务信息2", now, now, now},
		)
	}
}

// 用户信息数据填充
func seedUser() {

	now := time.Now()
	db := NewDB()

	var user model.User
	db.First(&user)
	if user.ID == 0 {
		db.Exec(
			"INSERT INTO `users` (`uid`, `name`, `gender`, `job_number`, `summary`, `role`, `department`, `face_token`, `face_image`, `status`, `created_at`, `updated_at`) VALUES (?), (?), (?)",
			[]interface{}{"1", "葛旭", "1", "sh003", "激光操作工", "工位工人", "实操部", "7c9a6ef3a7fdbd2269b08e8e0944dc92", "http://puh01tec3.bkt.clouddn.com/1563864233447917-%E5%9B%BE%E7%89%87%201.png", 0, now, now},
			[]interface{}{"2", "祥宁", "1", "sh004", "车间主任", "系统管理员", "管理部", "ea37e8fe7b6fd79c49ce690ac55803e9", "http://puh01tec3.bkt.clouddn.com/1563864293042131-%E5%9B%BE%E7%89%87%202.png", 0, now, now},
			[]interface{}{"3", "红兵", "1", "sh005", "厂长", "系统管理员", "管理部", "3c5a5e36036913ccd02e8475e99b5811", "http://puh01tec3.bkt.clouddn.com/1563864472441280-%E5%9B%BE%E7%89%87%203.png", 0, now, now},
		)
	}
}

// 颜色识别数据填充
func seedCheck() {

	db := NewDB()

	var check model.Check
	db.First(&check)
	if check.ID == 0 {
		db.Exec(
			"INSERT INTO `checks` (`url`, `colors`, `size`) VALUES (?)",
			[]interface{}{checkImg.Item, `{"hotpink":[[[160,90,150],[180,255,255]]],"yellow":[[[0,80,150],[50,255,255]]]}`, `{"w":600,"h":800}`},
		)
	}
}

// 工单质检数据填充
func seedOrderQuality() {

	db := NewDB()

	var quality model.ApsOrderQuality
	db.First(&quality)
	if quality.ID == 0 {

		var orders []model.ApsOrder
		db.Find(&orders)

		if orders != nil && len(orders) > 0 {
			for _, order := range orders {
				db.Exec(
					"INSERT INTO `aps_order_qualities` (`order_id`, `piece_no`, `result`, `remark`) VALUES (?), (?), (?)",
					[]interface{}{order.ID, "wu-887109902", "返工", ""},
					[]interface{}{order.ID, "wu-887109905", "报废", ""},
					[]interface{}{order.ID, "wu-887109907", "返工", ""},
				)
			}
		}
	}
}
