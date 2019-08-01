package database

import (
	"encoding/json"
	"sop/model"
	"time"
)

func seed() {

	seedUser()
	seedCraft()
	seedProduct()
	seedSop()
	seedApsOrder()
	seedAndon()
	seedCheck()
	seedOrderQuality()
}

type checkImgData struct {
	Item string `json:"item"`
	Bg   string `json:"bg"`
}

var checkImg = checkImgData{
	"http://puh01tec3.bkt.clouddn.com/1564121305087721-m1.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564477521235111-%E7%BB%84%201.png",
}

var mobileCheckImg = checkImgData{
	"http://puh01tec3.bkt.clouddn.com/1564633002633601-t9.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564560192110622-%E7%BB%84%201@2x.png",
}

var minefieldImg = []string{
	"http://puh01tec3.bkt.clouddn.com/1564121305087721-m1.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564121681493909-6051564121588_.pic.jpg",
}

var mobileMinefieldImg = []string{
	"http://puh01tec3.bkt.clouddn.com/1564633002633601-t9.jpeg",
	"http://puh01tec3.bkt.clouddn.com/1564633041967638-WechatIMG279.jpeg",
}

// 用户信息数据填充
func seedUser() {

	now := time.Now()
	db := NewDB()

	var user model.User
	db.First(&user)
	if user.ID == 0 {
		db.Exec(
			"INSERT INTO `users` (`uid`, `name`, `gender`, `job_number`, `summary`, `role`, `department`, `face_token`, `face_image`, `status`, `created_at`, `updated_at`) VALUES (?), (?), (?), (?), (?)",
			[]interface{}{"1", "评审专家", "1", "sh000", "管理人员", "超级管理员", "管理部", "e10adc3949ba59abbe56e057f20f883e", "http://puh01tec3.bkt.clouddn.com/1564542514821045-User.png", 0, now, now},
			[]interface{}{"2", "葛旭", "1", "sh003", "钻孔工", "工位工人", "实操部", "7c9a6ef3a7fdbd2269b08e8e0944dc92", "http://puh01tec3.bkt.clouddn.com/1563864233447917-%E5%9B%BE%E7%89%87%201.png", 0, now, now},
			[]interface{}{"3", "祥宁", "1", "sh004", "车间主任", "系统管理员", "管理部", "ea37e8fe7b6fd79c49ce690ac55803e9", "http://puh01tec3.bkt.clouddn.com/1563864293042131-%E5%9B%BE%E7%89%87%202.png", 0, now, now},
			[]interface{}{"4", "红兵", "1", "sh005", "厂长", "系统管理员", "管理部", "3c5a5e36036913ccd02e8475e99b5811", "http://puh01tec3.bkt.clouddn.com/1563864472441280-%E5%9B%BE%E7%89%87%203.png", 0, now, now},
			[]interface{}{"5", "韩鹏", "1", "sh006", "激光工", "工位工人", "实操部", "e5458e04cf70a9ad2e920d265bd7af20", "http://puh01tec3.bkt.clouddn.com/1564556172405803-%E5%9B%BE%E7%89%87%204.png", 0, now, now},
		)
	}
}

// 工艺方案填充
func seedCraft() {

	checkImgJson, _ := json.Marshal(checkImg)
	mobileCheckImgJson, _ := json.Marshal(mobileCheckImg)
	minefieldImgJson, _ := json.Marshal(minefieldImg)
	mobileMinefieldImgJson, _ := json.Marshal(mobileMinefieldImg)

	now := time.Now()
	db := NewDB()

	var craft model.Craft
	db.First(&craft)
	if craft.ID == 0 {
		db.Exec(
			"INSERT INTO `crafts` (`name`, `status`, `created_at`, `updated_at`) VALUES (?), (?), (?), (?), (?), (?)",
			[]interface{}{"#99手机外壳精加工", 0, now, now},
			[]interface{}{"#66手机外壳精加工", 0, now, now},
			[]interface{}{"#54手机外壳精加工", 0, now, now},
			[]interface{}{"#50鼠标外壳粗加工", 0, now, now},
			[]interface{}{"#90鼠标外壳粗加工", 0, now, now},
			[]interface{}{"#01鼠标外壳粗加工", 0, now, now},
		)

		var crafts []model.Craft
		db.Find(&crafts)
		if crafts != nil && len(crafts) > 0 {
			for _, c := range crafts {
				sql := "INSERT INTO `craft_items` (`craft_id`, `name`, `check_img`, `minefield_img`, `sort`, `status`, `created_at`, `updated_at`) VALUES "
				// TODO 处理防差错图片 和 雷区预警图片
				switch c.Name {
				case "#99手机外壳精加工":
					db.Exec(
						sql + "(?)",
						[]interface{}{c.ID, "镭雕", mobileCheckImgJson, mobileMinefieldImgJson, 1, 0, now, now},
					)
				case "#66手机外壳精加工":
					db.Exec(
						sql + "(?), (?)",
						[]interface{}{c.ID, "丝印", mobileCheckImgJson, mobileMinefieldImgJson, 1, 0, now, now},
						[]interface{}{c.ID, "镭雕", mobileCheckImgJson, mobileMinefieldImgJson, 2, 0, now, now},
					)
				case "#54手机外壳精加工":
					db.Exec(
						sql + "(?), (?)",
						[]interface{}{c.ID, "镭雕", mobileCheckImgJson, mobileMinefieldImgJson, 1, 0, now, now},
						[]interface{}{c.ID, "丝印", mobileCheckImgJson, mobileMinefieldImgJson, 2, 0, now, now},
					)
				case "#50鼠标外壳粗加工":
					db.Exec(
						sql + "(?), (?), (?)",
						[]interface{}{c.ID, "通圆孔", checkImgJson, minefieldImgJson, 1, 0, now, now},
						[]interface{}{c.ID, "通小孔", checkImgJson, minefieldImgJson, 2, 0, now, now},
						[]interface{}{c.ID, "表面处理", checkImgJson, minefieldImgJson, 3, 0, now, now},
					)
				case "#90鼠标外壳粗加工":
					db.Exec(
						sql + "(?), (?)",
						[]interface{}{c.ID, "通圆孔", checkImgJson, minefieldImgJson, 1, 0, now, now},
						[]interface{}{c.ID, "通小孔", checkImgJson, minefieldImgJson, 2, 0, now, now},
					)
				case "#01鼠标外壳粗加工":
					db.Exec(
						sql + "(?)",
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
			"INSERT INTO `products` (`name`, `created_at`, `updated_at`) VALUES (?), (?)",
			[]interface{}{"手机外壳", now, now},
			[]interface{}{"鼠标外壳", now, now},
		)

		var products []model.Product
		db.Find(&products)
		if products != nil && len(products) > 0 {
			for _, p := range products {
				switch p.Name {
				case "手机外壳":
					db.Exec(
						"INSERT INTO `product_models` (`product_id`, `name`, `created_at`, `updated_at`) VALUES (?), (?), (?)",
						[]interface{}{p.ID, "#11", now, now},
						[]interface{}{p.ID, "#22", now, now},
						[]interface{}{p.ID, "#33", now, now},
					)
				case "鼠标外壳":
					db.Exec(
						"INSERT INTO `product_models` (`product_id`, `name`, `created_at`, `updated_at`) VALUES (?), (?), (?)",
						[]interface{}{p.ID, "#77", now, now},
						[]interface{}{p.ID, "#78", now, now},
						[]interface{}{p.ID, "#79", now, now},
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

		var models []model.ProductModel
		db.Find(&models)

		if crafts != nil && models != nil && len(crafts) > 0 && len(models) > 0 {
			sql := "INSERT INTO `sops` (`title`, `craft_id`, `product_id`, `is_template`, `version`, `status`, `created_at`, `updated_at`) VALUES "
			for _, m := range models {
				for _, c := range crafts {
					if m.Name == "#77" {
						switch c.Name {
						case "#01鼠标外壳粗加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"鼠标外壳作业指导书模板01", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						case "#90鼠标外壳粗加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"鼠标外壳作业指导书模板02", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						case "#50鼠标外壳粗加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"鼠标外壳作业指导书模板03", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						}
					}
					if m.Name == "#22" {
						switch c.Name {
						case "#99手机外壳精加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"手机外壳作业指导书模板A", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						case "#66手机外壳精加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"手机外壳作业指导书模板B", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						case "#54手机外壳精加工":
							db.Exec(
								sql + "(?)",
								[]interface{}{"手机外壳作业指导书模板C", c.ID, m.ProductID, 1, 1, 0, now, now},
							)
						}
					}
				}
			}
		}

		var sops []model.Sop
		db.Find(&sops)
		if sops != nil && len(sops) > 0 {

			checkImgJson, _ := json.Marshal(checkImg)
			mobileCheckImgJson, _ := json.Marshal(mobileCheckImg)
			minefieldImgJson, _ := json.Marshal(minefieldImg)
			mobileMinefieldImgJson, _ := json.Marshal(mobileMinefieldImg)

			for _, s := range sops {

				var items []model.CraftItem
				db.Where("craft_id = ?", s.CraftID).Find(&items)

				if items != nil && len(items) > 0 {
					for _, t := range items {
						var content = ""
						var checkImgJsonStr = ""
						var minefieldImgJsonStr = ""
						switch s.Title {
						case "鼠标外壳作业指导书模板01":
							if t.Name == "通孔" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626480194786-1.png\" title=\"1564626480194675000326217.png\" alt=\"1.png\"/></td><td valign=\"top\" style=\"word-break: break-all;\" rowspan=\"3\" colspan=\"1\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、圆孔处理</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:#222A35;text-combine:letters\">1.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">在现有的通孔加工过程中，往往是在利用样冲冲出定位孔后，利用钻头直接钻屑，这样对于钻头的损坏较大；并且在对通孔进行扩孔、铰孔时，需要更换多种不同类型的刀具，使得加工过程比较繁琐，不利于实现快速加工。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">2.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">技术实现要素：针对现有技术中的问题，本发明所述的一种圆孔加工方法，本方法首先通过激光对工件待钻孔部位加工出一个通孔，然后利用钻头进行钻屑，可减小钻头受到的轴向力，避免钻头损坏，并且在进行通孔的扩孔时，利用多头铰刀实现快速加工，避免了频繁装夹刀具的繁琐过程，大大提高了加工的效率。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用激光对工件上需要钻孔的地方打孔，形成激光孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用钻机的钻头沿着</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的激光孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用多头铰刀对</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的通孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S4</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">完成后，将工件运送至下一步作业；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626628796256-3.png\" title=\"1564626628796205000742349.png\" alt=\"3.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"},{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626647503209-2.png\" title=\"1564626647503122000734368.png\" alt=\"2.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">二、微小孔处理</span></p><span style=\"font-size:20px\">1.</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">微小孔加工业务，最小加工微孔孔径（线宽）</span><span style=\"font-size:20px;font-family:Calibri;color:black\">1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">整体加工精度最高</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">±1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size:20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">局部特征精度＜</span><span style=\"font-size:20px;font-family:Calibri;color:black\">0.1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，加工幅面</span><span style=\"font-size:20px;font-family:Calibri;color:black\">300*300</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">mm</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">。</span> <span style=\"font-size:20px\">●</span> <p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\"><br/></span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">2.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精心选用进口高功率高稳定性光纤激光器，接近衍射极限光束质量，是微孔加工理想光源</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">3.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">进口高速精密振镜和精密远心场镜，保证打孔一致性和重复精度；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">4.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精密视觉检测和校正功能，保证系统长期精度，检测孔径，方便工艺调试；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">5.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">直线电机</span><span style=\"font-size:20px;font-family:Calibri;color:black\">XY</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">平台，长期免维护，微米级定位和重复精度，</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:black\">扩展加工幅面。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							}
							checkImgJsonStr = string(checkImgJsonStr)
							minefieldImgJsonStr = string(minefieldImgJsonStr)
						case "鼠标外壳作业指导书模板02":
							if t.Name == "通圆孔" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626875685599-1.png\" title=\"1564626875685473000729699.png\" alt=\"1.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、圆孔处理</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:#222A35;text-combine:letters\">1.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">在现有的通孔加工过程中，往往是在利用样冲冲出定位孔后，利用钻头直接钻屑，这样对于钻头的损坏较大；并且在对通孔进行扩孔、铰孔时，需要更换多种不同类型的刀具，使得加工过程比较繁琐，不利于实现快速加工。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">2.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">技术实现要素：针对现有技术中的问题，本发明所述的一种圆孔加工方法，本方法首先通过激光对工件待钻孔部位加工出一个通孔，然后利用钻头进行钻屑，可减小钻头受到的轴向力，避免钻头损坏，并且在进行通孔的扩孔时，利用多头铰刀实现快速加工，避免了频繁装夹刀具的繁琐过程，大大提高了加工的效率。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用激光对工件上需要钻孔的地方打孔，形成激光孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用钻机的钻头沿着</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的激光孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用多头铰刀对</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的通孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S4</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">完成后，将工件运送至下一步作业；</span></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626896106043-3.png\" title=\"1564626896105992000246634.png\" alt=\"3.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							} else if t.Name == "通小孔" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626920514049-2.png\" title=\"1564626920513980000582083.png\" alt=\"2.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、微小孔处理</span></p><span style=\"font-size:20px\">1.</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">微小孔加工业务，最小加工微孔孔径（线宽）</span><span style=\"font-size:20px;font-family:Calibri;color:black\">1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">整体加工精度最高</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">±1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size:20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">局部特征精度＜</span><span style=\"font-size:20px;font-family:Calibri;color:black\">0.1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，加工幅面</span><span style=\"font-size:20px;font-family:Calibri;color:black\">300*300</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">mm</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">。</span> <span style=\"font-size:20px\">●</span> <p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">2.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精心选用进口高功率高稳定性光纤激光器，接近衍射极限光束质量，是微孔加工理想光源</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">3.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">进口高速精密振镜和精密远心场镜，保证打孔一致性和重复精度；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">4.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精密视觉检测和校正功能，保证系统长期精度，检测孔径，方便工艺调试；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">5.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">直线电机</span><span style=\"font-size:20px;font-family:Calibri;color:black\">XY</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">平台，长期免维护，微米级定位和重复精度，</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:black\">扩展加工幅面。</span></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564626946196450-3.png\" title=\"1564626946196392000823320.png\" alt=\"3.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							}
							checkImgJsonStr = string(checkImgJsonStr)
							minefieldImgJsonStr = string(minefieldImgJsonStr)
						case "鼠标外壳作业指导书模板03":
							if t.Name == "通圆孔" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627014293031-1.png\" title=\"1564627014292976000040141.png\" alt=\"1.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、圆孔处理</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:#222A35;text-combine:letters\">1.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">在现有的通孔加工过程中，往往是在利用样冲冲出定位孔后，利用钻头直接钻屑，这样对于钻头的损坏较大；并且在对通孔进行扩孔、铰孔时，需要更换多种不同类型的刀具，使得加工过程比较繁琐，不利于实现快速加工。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">2.</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">技术实现要素：针对现有技术中的问题，本发明所述的一种圆孔加工方法，本方法首先通过激光对工件待钻孔部位加工出一个通孔，然后利用钻头进行钻屑，可减小钻头受到的轴向力，避免钻头损坏，并且在进行通孔的扩孔时，利用多头铰刀实现快速加工，避免了频繁装夹刀具的繁琐过程，大大提高了加工的效率。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用激光对工件上需要钻孔的地方打孔，形成激光孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用钻机的钻头沿着</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S1</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的激光孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">利用多头铰刀对</span><span style=\"font-size: 19px;font-family:Calibri;color:black\">S2</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">中的通孔进行扩孔；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:19px;font-family:Calibri;color:black\">S4</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:19px;font-family:Calibri;color:black\">S3</span><span style=\"font-size:19px;font-family:微软雅黑;color:black\">完成后，将工件运送至下一步作业；</span></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627033744450-3.png\" title=\"1564627033744402000227110.png\" alt=\"3.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							} else if t.Name == "通小孔" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627059807674-2.png\" title=\"1564627059807612000132312.png\" alt=\"2.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、微小孔处理</span></p><span style=\"font-size:20px\">1.</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">微小孔加工业务，最小加工微孔孔径（线宽）</span><span style=\"font-size:20px;font-family:Calibri;color:black\">1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">整体加工精度最高</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">±1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size:20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">局部特征精度＜</span><span style=\"font-size:20px;font-family:Calibri;color:black\">0.1</span><span style=\"font-size:20px;font-family:Calibri;color:black\">μ</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">m</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">，加工幅面</span><span style=\"font-size:20px;font-family:Calibri;color:black\">300*300</span><span style=\"font-size: 20px;font-family:Calibri;color:black\">mm</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">。</span> <span style=\"font-size:20px\">●</span> <p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">2.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精心选用进口高功率高稳定性光纤激光器，接近衍射极限光束质量，是微孔加工理想光源</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">3.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">进口高速精密振镜和精密远心场镜，保证打孔一致性和重复精度；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">4.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">精密视觉检测和校正功能，保证系统长期精度，检测孔径，方便工艺调试；</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:Calibri;color:black\">5.&nbsp;</span><span style=\"font-size: 20px;font-family:微软雅黑;color:black\">直线电机</span><span style=\"font-size:20px;font-family:Calibri;color:black\">XY</span><span style=\"font-size:20px;font-family:微软雅黑;color:black\">平台，长期免维护，微米级定位和重复精度，</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:black\">扩展加工幅面。</span></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627074986439-3.png\" title=\"1564627074986369000572194.png\" alt=\"3.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							} else if t.Name == "表面处理" {
								content = `[{"view":"<p style=\";margin-top:0;margin-bottom:0;text-align:center;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed; text-align: center;\"><span style=\"font-size:20px;font-family:微软雅黑;color:#222A35;text-combine:letters\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed; text-align: center;\"><span style=\"font-size: 20px; font-family: 微软雅黑; color: rgb(34, 42, 53);\">该工序无</span><span style=\"font-size: 20px; font-family: Calibri; color: rgb(34, 42, 53);\">sop</span><br/></p><p><br/></p>"}]`
							}
							checkImgJsonStr = string(checkImgJsonStr)
							minefieldImgJsonStr = string(minefieldImgJsonStr)
						case "手机外壳作业指导书模板A":
							if t.Name == "镭雕" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627227104547-4.png\" title=\"1564627227104486000456708.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、镭雕工艺</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:#222A35;text-combine:letters\">1</span><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">塑件清洁：塑胶部件注塑出来后表面会残留部分油污，以及空气环境中细小的灰尘，通过溶剂擦拭，静电风枪除尘将这些油污及灰尘清除，可以获得较好的良率</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">2</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在塑胶表面喷涂一层双组份热固化油漆，此层油漆与塑胶料及</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆都有很好的附着力，同时此层油漆最后通过镭雕显现给用户，可以通过增加珠光粉，色粉等达到设计的色彩要求</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">3</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆表面喷涂一层</span><span style=\"font-size: 21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆，主要是为了保证与金属镀膜层的连接，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆与金属膜层无附着力；</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆为透明清漆，烘烤一定时间后</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">光固化；固化能量一般在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">800-1200</span><span style=\"font-size:21px;font-family:Calibri;color:black\">mj</span><span style=\"font-size:21px;font-family:Calibri;color:black\">/cm2&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627246888632-5.png\" title=\"1564627246888575000923091.png\" alt=\"5.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"},{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627279451001-4.png\" title=\"1564627279450622000634679.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">4</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">真空镀膜：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆的表面进行真空镀膜，通过物理沉积的方式在产品表面形成一层纳米级厚度的不连续金属膜，赋予塑胶材料金属的外观质感；因镀膜材料都是非导电性氧化物，且膜层在微观上不连续，所以膜层不导电，对射频不会产生影响。常用的镀膜金属材质为铟。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">5</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂保护层：真空镀膜完成后在金属膜层表面喷涂一层保护层，防止后面拆装夹具及镭雕过程中产品划伤；保护层要薄，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">2-4</span><span style=\"font-size:21px;font-family:Calibri;color:black\">um</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">即可</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">6</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）镭雕图案：产品拆掉喷涂夹具后放置在设计好的精密镭雕治具上，固定紧扣。按照</span><span style=\"font-size:21px;font-family:Calibri;color:black\">ID</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">设计的图案设置镭雕程序，激光功率参数，步进速度等关键参数；自主设计的传动系统保证产品激光焦距一直在产品表面，以获得好的镭雕效果；镭雕后产品表面残留的粉尘需用专用的清洗剂擦拭除去，清洗剂一般用低分子的醇类和酯类。</span></p></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							}
							checkImgJsonStr = string(mobileCheckImgJson)
							minefieldImgJsonStr = string(mobileMinefieldImgJson)
						case "手机外壳作业指导书模板B":
							if t.Name == "丝印" {
								content = `[{"view":"<p style=\";margin-top:0;margin-bottom:0;text-align:center;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed; text-align: center;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed; text-align: center;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\">该工序无</span><span style=\"font-size:21px;font-family:Calibri;color:black\">sop</span></p><p><br/></p>"}]`
							} else if t.Name == "镭雕" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627392579554-4.png\" title=\"1564627392579442000292741.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、镭雕工艺</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:#222A35;text-combine:letters\">1</span><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">塑件清洁：塑胶部件注塑出来后表面会残留部分油污，以及空气环境中细小的灰尘，通过溶剂擦拭，静电风枪除尘将这些油污及灰尘清除，可以获得较好的良率</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">2</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在塑胶表面喷涂一层双组份热固化油漆，此层油漆与塑胶料及</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆都有很好的附着力，同时此层油漆最后通过镭雕显现给用户，可以通过增加珠光粉，色粉等达到设计的色彩要求</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">3</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆表面喷涂一层</span><span style=\"font-size: 21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆，主要是为了保证与金属镀膜层的连接，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆与金属膜层无附着力；</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆为透明清漆，烘烤一定时间后</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">光固化；固化能量一般在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">800-1200</span><span style=\"font-size:21px;font-family:Calibri;color:black\">mj</span><span style=\"font-size:21px;font-family:Calibri;color:black\">/cm2&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627413541441-5.png\" title=\"1564627413541390000806742.png\" alt=\"5.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"},{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627441331016-4.png\" title=\"1564627441330970000269729.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">4</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">真空镀膜：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆的表面进行真空镀膜，通过物理沉积的方式在产品表面形成一层纳米级厚度的不连续金属膜，赋予塑胶材料金属的外观质感；因镀膜材料都是非导电性氧化物，且膜层在微观上不连续，所以膜层不导电，对射频不会产生影响。常用的镀膜金属材质为铟。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">5</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂保护层：真空镀膜完成后在金属膜层表面喷涂一层保护层，防止后面拆装夹具及镭雕过程中产品划伤；保护层要薄，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">2-4</span><span style=\"font-size:21px;font-family:Calibri;color:black\">um</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">即可</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">6</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）镭雕图案：产品拆掉喷涂夹具后放置在设计好的精密镭雕治具上，固定紧扣。按照</span><span style=\"font-size:21px;font-family:Calibri;color:black\">ID</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">设计的图案设置镭雕程序，激光功率参数，步进速度等关键参数；自主设计的传动系统保证产品激光焦距一直在产品表面，以获得好的镭雕效果；镭雕后产品表面残留的粉尘需用专用的清洗剂擦拭除去，清洗剂一般用低分子的醇类和酯类。</span></p></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							}
							checkImgJsonStr = string(mobileCheckImgJson)
							minefieldImgJsonStr = string(mobileMinefieldImgJson)
						case "手机外壳作业指导书模板C":
							if t.Name == "镭雕" {
								content = `[{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627612741617-4.png\" title=\"1564627612741561000591906.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">操作步骤：</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">一、镭雕工艺</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:#222A35;text-combine:letters\">1</span><span style=\"font-size:21px;font-family:微软雅黑;color:#222A35;text-combine:letters\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">塑件清洁：塑胶部件注塑出来后表面会残留部分油污，以及空气环境中细小的灰尘，通过溶剂擦拭，静电风枪除尘将这些油污及灰尘清除，可以获得较好的良率</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">2</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在塑胶表面喷涂一层双组份热固化油漆，此层油漆与塑胶料及</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆都有很好的附着力，同时此层油漆最后通过镭雕显现给用户，可以通过增加珠光粉，色粉等达到设计的色彩要求</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">3</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆表面喷涂一层</span><span style=\"font-size: 21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆，主要是为了保证与金属镀膜层的连接，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">PU</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆与金属膜层无附着力；</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆为透明清漆，烘烤一定时间后</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">光固化；固化能量一般在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">800-1200</span><span style=\"font-size:21px;font-family:Calibri;color:black\">mj</span><span style=\"font-size:21px;font-family:Calibri;color:black\">/cm2&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627633082031-5.png\" title=\"1564627633081859000035664.png\" alt=\"5.png\"/></td></tr><tr></tr><tr></tr></tbody></table>"},{"view":"<table class=\"rb\" border=\"“1”\" cellpadding=\"0\" cellspacing=\"0\" data-sort=\"sortDisabled\"><tbody><tr class=\"firstRow\"><td valign=\"top\" rowspan=\"3\" colspan=\"1\"><img src=\"http://puh01tec3.bkt.clouddn.com/1564627672465791-4.png\" title=\"1564627672465723000071342.png\" alt=\"4.png\"/></td><td valign=\"top\" rowspan=\"3\" colspan=\"1\" style=\"word-break: break-all;\"><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">4</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">真空镀膜：在</span><span style=\"font-size:21px;font-family:Calibri;color:black\">UV</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">底漆的表面进行真空镀膜，通过物理沉积的方式在产品表面形成一层纳米级厚度的不连续金属膜，赋予塑胶材料金属的外观质感；因镀膜材料都是非导电性氧化物，且膜层在微观上不连续，所以膜层不导电，对射频不会产生影响。常用的镀膜金属材质为铟。</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><br/></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">5</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）喷涂保护层：真空镀膜完成后在金属膜层表面喷涂一层保护层，防止后面拆装夹具及镭雕过程中产品划伤；保护层要薄，</span><span style=\"font-size:21px;font-family:Calibri;color:black\">2-4</span><span style=\"font-size:21px;font-family:Calibri;color:black\">um</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">即可</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">&nbsp;</span></p><p style=\";margin-top:0;margin-bottom:0;text-align:left;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:Calibri;color:black\">6</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">）镭雕图案：产品拆掉喷涂夹具后放置在设计好的精密镭雕治具上，固定紧扣。按照</span><span style=\"font-size:21px;font-family:Calibri;color:black\">ID</span><span style=\"font-size:21px;font-family:微软雅黑;color:black\">设计的图案设置镭雕程序，激光功率参数，步进速度等关键参数；自主设计的传动系统保证产品激光焦距一直在产品表面，以获得好的镭雕效果；镭雕后产品表面残留的粉尘需用专用的清洗剂擦拭除去，清洗剂一般用低分子的醇类和酯类。</span></p></td></tr><tr></tr><tr></tr></tbody></table>"}]`
							} else if t.Name == "丝印" {
								content = `[{"view":"<p style=\";margin-top:0;margin-bottom:0;text-align:center;direction:ltr;unicode-bidi:embed\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size:21px;font-family:微软雅黑;color:black\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed;\"><span style=\"font-size: 21px; font-family: 微软雅黑;\"><br/></span></p><p style=\"margin-top: 0px; margin-bottom: 0px; direction: ltr; unicode-bidi: embed; text-align: center;\"><span style=\"font-size: 21px; font-family: 微软雅黑;\">该工序无</span><span style=\"font-size: 21px; font-family: Calibri;\">sop</span><br/></p><p><br/></p>"}]`
							}
							checkImgJsonStr = string(mobileCheckImgJson)
							minefieldImgJsonStr = string(mobileMinefieldImgJson)
						}
						db.Exec(
							"INSERT INTO `sop_processes` (`sop_id`, `process_id`, `title`, `content`, `check_img`, `minefield_img`, `is_check`, `sort`, `created_at`, `updated_at`) VALUES (?)",
							[]interface{}{s.ID, t.ID, t.Name, content, checkImgJsonStr, minefieldImgJsonStr, 0, t.Sort, now, now},
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

					sql := "INSERT INTO `aps` (`job_plan_number`, `serial_no`, `model_id`, `craft_id`, `plan_total`, `plan_num`, `complete_num`, `status`, `created_at`, `updated_at`) VALUES "
					switch m.Name {
					case "#78":
						if c.Name == "#01鼠标外壳粗加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201987771341", "line-m001", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					case "#77":
						if c.Name == "#90鼠标外壳粗加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201966550001", "line-m002", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					case "#79":
						if c.Name == "#50鼠标外壳粗加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201966550002", "line-m003", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					case "#11":
						if c.Name == "#99手机外壳精加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201900220011", "line-p001", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					case "#22":
						if c.Name == "#66手机外壳精加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201900220012", "line-p002", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					case "#33":
						if c.Name == "#54手机外壳精加工" {
							db.Exec(
								sql + "(?)",
								[]interface{}{"WP201900220013", "line-p003", m.ID, c.ID, 20, 7, 4, 0, now, now},
							)
						}
					}
				}
			}

			var apsList []model.Aps
			db.Find(&apsList)

			var users []model.User
			db.Find(&users)

			if apsList != nil && len(apsList) > 0 && users != nil && len(users) > 0 {
				for _, a := range apsList {
					for _, user := range users {

						sql := "INSERT INTO `aps_orders` (`order_id`, `aps_id`, `uid`, `process_id`, `station`, `station_name`, `total`, `num`, `complete_num`, `status`, `created_at`, `updated_at`) VALUES "
						var craftItems []model.CraftItem
						var craftIds = make([]uint, 0)
						db.Where("craft_id = ?", a.CraftID).Find(&craftItems)
						if craftItems != nil && len(craftItems) > 0 {
							for _, c := range craftItems {
								craftIds = append(craftIds, c.ID)
							}
						}

						if len(craftIds) > 0 {
							if user.Name == "葛旭" {
								switch a.JobPlanNumber {
								case "WP201987771341":
									db.Exec(
										sql+"(?)",
										[]interface{}{"WO1987701242", a.ID, user.ID, craftIds[0], 1, "通孔工位", 20, 7, 4, 0, now, now},
									)
								case "WP201966550001":
									db.Exec(
										sql+"(?), (?)",
										[]interface{}{"WO1987701243", a.ID, user.ID, craftIds[0], 1, "通孔工位", 20, 7, 4, 0, now, now},
										[]interface{}{"WO1987701244", a.ID, user.ID, craftIds[1], 1, "通孔工位", 20, 7, 4, 0, now, now},
									)
								case "WP201966550002":
									db.Exec(
										sql+"(?), (?), (?)",
										[]interface{}{"WO1987701245", a.ID, user.ID, craftIds[0], 1, "通孔工位", 20, 7, 4, 0, now, now},
										[]interface{}{"WO1987701246", a.ID, user.ID, craftIds[1], 1, "通孔工位", 20, 7, 4, 0, now, now},
										[]interface{}{"WO1987701247", a.ID, user.ID, craftIds[2], 2, "激光工位", 20, 7, 4, 0, now, now},
									)
								}
							}
							if user.Name == "韩鹏" {
								switch a.JobPlanNumber {
								case "WP201900220011":
									db.Exec(
										sql+"(?)",
										[]interface{}{"WO1987701248", a.ID, user.ID, craftIds[0], 2, "激光工位", 20, 7, 4, 0, now, now},
									)
								case "WP201900220012":
									db.Exec(
										sql+"(?), (?)",
										[]interface{}{"WO1987701249", a.ID, user.ID, craftIds[0], 2, "激光工位", 20, 7, 4, 0, now, now},
										[]interface{}{"WO1987701250", a.ID, user.ID, craftIds[1], 2, "激光工位", 20, 7, 4, 0, now, now},
									)
								case "WP201900220013":
									db.Exec(
										sql+"(?), (?)",
										[]interface{}{"WO1987701251", a.ID, user.ID, craftIds[0], 2, "激光工位", 20, 7, 4, 0, now, now},
										[]interface{}{"WO1987701252", a.ID, user.ID, craftIds[1], 2, "激光工位", 20, 7, 4, 0, now, now},
									)
								}
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

// 颜色识别数据填充
func seedCheck() {

	db := NewDB()

	var check model.Check
	db.First(&check)
	if check.ID == 0 {
		db.Exec(
			"INSERT INTO `checks` (`url`, `colors`, `size`) VALUES (?), (?)",
			[]interface{}{checkImg.Item, `{"hotpink":[[[160,90,150],[180,255,255]]],"yellow":[[[0,80,150],[50,255,255]]]}`, `{"w":600,"h":800}`},
			[]interface{}{mobileCheckImg.Item, `{"hotpink":[[[150,43,46],[180,255,255]]],"green":[[[35,100,101],[100,255,255]]]}`, `{"w":600,"h":800}`},
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
