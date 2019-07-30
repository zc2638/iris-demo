package controller

import (
	"encoding/json"
	"fmt"
	"sop/model"
	"sop/service"
	"strings"
	"time"
)

type SopController struct{ Base }

// 获取sop列表(不分页)
func (c *SopController) GetList() {

	list := new(service.SopService).GetAll()

	var data = make([]map[string]interface{}, 0)
	for _, v := range list {

		models := make([]string, 0)
		if v.SopModel != nil && len(v.SopModel) > 0 {
			for _, m := range v.SopModel {
				models = append(models, m.Name)
			}
		}

		data = append(data, map[string]interface{}{
			"id": v.ID,
			// 产品名称
			"productName": v.Product.Name,
			// 产品类型
			"productType": strings.Join(models, " "),
			// sop名称
			"sopName": v.Title,
			// sop类型 表示是存为模版还是存为普通sop
			"sopType": v.IsTemplate,
			// 驳回理由
			"tip": v.Comment,
			// 版本
			"version": v.Version,
			// 状态 0是已提交 1是已通过 2是模版通过 3是驳回
			"sopStatus": v.Status,
		})
	}
	c.Succ("", data)
}

// 获取所有sop模板
func (c *SopController) GetTemplate() {

	craftID := c.Ctx.URLParam("craft_id")

	all := new(service.SopService).GetTemplateAllByCraftID(craftID)

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {

		table := make([]map[string]interface{}, 0)
		if v.SopProcess != nil && len(v.SopProcess) > 0 {
			for _, process := range v.SopProcess {
				table = append(table, map[string]interface{}{
					// 工序id
					"sopStepID": process.ID,
					// 工序名称
					"sopStepName": process.Title,
					// sop工序内容
					"sop": process.Content,
					// 附件地址
					"ATT": process.Annex,
					// 是否开启防差错标识
					"py":     process.IsCheck,
					"pyInfo": process.CheckImg,
				})
			}
		}

		data = append(data, map[string]interface{}{
			// sop模板id
			"value": v.ID,
			// sop模板名称,
			"label": v.Title,
			"table": table,
		})
	}

	c.Succ("", data)
}

// 获取sop详情
func (c *SopController) GetShow() {

	id := c.Ctx.URLParam("id")

	sop := new(service.SopService).GetSopByID(id)
	models := make([]string, 0)
	modelIds := make([]uint, 0)
	if sop.SopModel != nil && len(sop.SopModel) > 0 {
		for _, m := range sop.SopModel {
			models = append(models, m.Name)
			modelIds = append(modelIds, m.ModelID)
		}
	}
	c.Succ("", map[string]interface{}{
		"id": sop.ID,
		// 产品id
		"productId": sop.ProductID,
		// 产品名称
		"productName": sop.Product.Name,
		// 产品型号id集合
		"productTypeIds": modelIds,
		// 产品类型
		"productType": strings.Join(models, " "),
		// sop名称
		"sopName": sop.Title,
		// sop类型 表示是存为模版还是存为普通sop
		"sopType": sop.IsTemplate,
		// 驳回理由
		"tip": sop.Comment,
		// 版本
		"version": sop.Version,
		// 状态 0是已提交 1是已通过 2是模版通过 3是驳回
		"sopStatus": sop.Status,
		// 工艺方案id
		"craftId": sop.CraftID,
		// 工艺方案名称
		"craftName": sop.Craft.Name,
		// 工序
		"process": sop.SopProcess,
	})
}

// 创建sop
func (c *SopController) PostCreate() {

	ctx := c.Ctx
	title := ctx.PostValue("title")
	craftID := ctx.PostValue("craft_id")               // 工艺方案id
	productID := ctx.PostValue("product_id")           // 产品id
	modelIds := ctx.PostValue("model_ids")             // 产品型号id
	processContent := ctx.PostValue("process")         // 工序
	isTemplate, err := ctx.PostValueInt("is_template") // 是否模板
	if err != nil {
		c.Err("模板参数异常")
		return
	}

	if title == "" {
		c.Err("请填写标题")
		return
	}

	productService := new(service.ProductService)
	product := productService.GetProductByID(productID)
	if product.ID == 0 {
		c.Err("产品不存在")
		return
	}

	craftService := new(service.CraftService)
	craft := craftService.GetCraftByID(craftID)
	if craft.ID == 0 {
		c.Err("工艺方案不存在")
		return
	}

	var modelIdSet []uint
	if err := json.Unmarshal([]byte(modelIds), &modelIdSet); err != nil {
		c.Err("产品型号错误")
		return
	}

	modelSet := productService.GetModelByIDs(modelIdSet)
	//if modelSet == nil || len(modelSet) == 0 {
	//	c.Err("产品型号不存在")
	//	return
	//}

	var processSet []struct {
		CraftItemID uint   `json:"craft_item_id"`
		Content     string `json:"content"`
		Annex       string `json:"annex"`
		IsCheck     uint   `json:"is_check"`
	}

	if err := json.Unmarshal([]byte(processContent), &processSet); err != nil {
		fmt.Println(processContent)
		c.Err("sop工序数据解析失败")
		return
	}
	if len(processSet) == 0 {
		c.Err("请上传sop工序数据")
		return
	}

	itemIds := make([]interface{}, 0)
	for _, p := range processSet {
		itemIds = append(itemIds, p.CraftItemID)
	}

	craftSet := craftService.GetItemByIds(itemIds)
	if craftSet == nil || len(craftSet) == 0 {
		c.Err("sop工序数据不存在")
		return
	}

	var version uint = 1
	if c.Data != nil {
		if _, ok := c.Data["version"]; ok {
			version = c.Data["version"].(uint)
		}
	}

	sop := model.Sop{
		Title:      title,
		CraftID:    craft.ID,
		ProductID:  product.ID,
		IsTemplate: uint(isTemplate),
		Version:    version,
	}

	process := make([]model.SopProcess, 0)
	for _, s := range processSet {
		for _, c := range craftSet {
			if c.ID == s.CraftItemID {
				process = append(process, model.SopProcess{
					ProcessID: c.ID,
					Title:     c.Name,
					Content:   s.Content,
					Annex:     s.Annex,
					CheckImg:  c.CheckImg,
					IsCheck:   s.IsCheck,
					Sort:      c.Sort,
				})
			}
		}
	}

	models := make([]model.SopModel, 0)
	for _, m := range modelSet {
		if m.ID == 0 {
			continue
		}
		models = append(models, model.SopModel{
			ModelID: m.ID,
			Name:    m.Name,
		})
	}

	if err := new(service.SopService).Create(sop, process, models); err != nil {
		c.Err("操作失败")
		return
	}

	c.Succ("操作成功")
}

// sop修改(自动迭代版本，不在原版本上修改，该接口只用作审核通过的sop的修改)
func (c *SopController) PostUpdate() {

	id := c.Ctx.PostValue("id")
	sop := new(service.SopService).GetSopByID(id)
	if sop.ID == 0 {
		c.Err("不存在的sop")
		return
	}
	c.Data = map[string]interface{}{"version": sop.Version + 1}
	c.PostCreate()
}

// sop更新
func (c *SopController) PostModify() {

	ctx := c.Ctx
	id := ctx.PostValue("id")
	title := ctx.PostValue("title")
	craftID := ctx.PostValue("craft_id")               // 工艺方案id
	productID := ctx.PostValue("product_id")           // 产品id
	modelIds := ctx.PostValue("model_ids")             // 产品型号id
	processContent := ctx.PostValue("process")         // 工序
	isTemplate, err := ctx.PostValueInt("is_template") // 是否模板
	if err != nil {
		c.Err("模板参数异常")
		return
	}

	if title == "" {
		c.Err("请填写标题")
		return
	}

	sop := new(service.SopService).GetSopByID(id)
	if sop.ID == 0 {
		c.Err("不存在的sop")
		return
	}

	productService := new(service.ProductService)
	product := productService.GetProductByID(productID)
	if product.ID == 0 {
		c.Err("产品不存在")
		return
	}

	craftService := new(service.CraftService)
	craft := craftService.GetCraftByID(craftID)
	if craft.ID == 0 {
		c.Err("工艺方案不存在")
		return
	}

	var modelIdSet []uint
	if err := json.Unmarshal([]byte(modelIds), &modelIdSet); err != nil {
		c.Err("产品型号错误")
		return
	}

	modelSet := productService.GetModelByIDs(modelIdSet)
	if modelSet == nil || len(modelSet) == 0 {
		c.Err("产品型号不存在")
		return
	}

	var processSet []struct {
		CraftItemID uint   `json:"craft_item_id"`
		Content     string `json:"content"`
		Annex       string `json:"annex"`
		IsCheck     uint   `json:"is_check"`
	}

	if err := json.Unmarshal([]byte(processContent), &processSet); err != nil {
		fmt.Println(processContent)
		c.Err("sop工序数据解析失败")
		return
	}
	if len(processSet) == 0 {
		c.Err("请上传sop工序数据")
		return
	}

	itemIds := make([]interface{}, 0)
	for _, p := range processSet {
		itemIds = append(itemIds, p.CraftItemID)
	}

	craftSet := craftService.GetItemByIds(itemIds)
	if craftSet == nil || len(craftSet) == 0 {
		c.Err("sop工序数据不存在")
		return
	}

	sop.Title = title
	sop.CraftID = craft.ID
	sop.ProductID = product.ID
	sop.IsTemplate = uint(isTemplate)

	process := make([]model.SopProcess, 0)
	for _, s := range processSet {
		for _, c := range craftSet {
			if c.ID == s.CraftItemID {
				process = append(process, model.SopProcess{
					ProcessID:    c.ID,
					Title:        c.Name,
					Content:      s.Content,
					Annex:        s.Annex,
					CheckImg:     c.CheckImg,
					MinefieldImg: c.MinefieldImg,
					IsCheck:      s.IsCheck,
					Sort:         c.Sort,
				})
			}
		}
	}

	models := make([]model.SopModel, 0)
	for _, m := range modelSet {
		if m.ID == 0 {
			continue
		}
		models = append(models, model.SopModel{
			SopID:   sop.ID,
			ModelID: m.ID,
			Name:    m.Name,
		})
	}

	if err := new(service.SopService).UpdateOne(sop, process, models); err != nil {
		c.Err("操作失败")
		return
	}

	c.Succ("操作成功")
}

// sop审核
func (c *SopController) PostAudit() {

	sopID := c.Ctx.PostValue("sop_id")
	reason := c.Ctx.PostValueTrim("reason")
	sopSet := c.Ctx.PostValue("sop_set")
	status, err := c.Ctx.PostValueInt("status")
	if err != nil {
		c.Err("状态错误")
		return
	}

	sopService := new(service.SopService)
	sop := sopService.GetSopByID(sopID)
	if sop.ID == 0 {
		c.Err("不存在的sop")
		return
	}

	sop.Status = uint(status)

	if status == model.SOP_REFUSE {
		if reason == "" {
			c.Err("请填写驳回理由")
			return
		}
		sop.Comment = reason
		sop.RefuseAt = time.Now()
	}

	var processes []model.SopProcess

	var set []struct {
		ProcessID uint     `json:"process_id"` // 工序id
		Images    []string `json:"images"`     // 工序图片数组
	}

	var processData = make([]model.SopProcess, 0)
	// 接收图片url
	if status == model.SOP_PASS {
		if err := json.Unmarshal([]byte(sopSet), &set); err != nil {
			c.Err("图片数据解析失败")
			return
		}
		var processIds = make([]uint, 0)
		for _, v := range set {
			if v.ProcessID != 0 {
				processIds = append(processIds, v.ProcessID)
			}
		}
		processes = sopService.GetProcessBySopID(sop.ID)

		for _, p := range processes {
			for _, s := range set {
				if s.ProcessID == p.ID {
					imgBytes, err := json.Marshal(s.Images)
					if err != nil {
						c.Err("图片格式异常")
						return
					}
					p.Imgs = string(imgBytes)
					processData = append(processData, p)
				}
			}
		}
		sop.PassAt = time.Now()
	}

	if err := sopService.UpdateOne(sop, processData, nil); err != nil {
		c.Err("操作失败")
		return
	}

	// 审核通过，推送大数据平台
	//go func(sop model.Sop) {
	//
	//	if sop.ID == 0 && sop.Status == model.SOP_PASS {
	//		return
	//	}
	//
	//	var sops = struct {
	//		ID        uint `json:"id"`
	//		CraftID   uint `json:"craft_id"`
	//		ProductID uint `json:"product_id"`
	//		Status    uint `json:"status"`
	//	}{sop.ID, sop.CraftID, sop.ProductID, sop.Status}
	//
	//	var sopModels = make([]struct {
	//		SopID   uint `json:"sop_id"`
	//		ModelID uint `json:"model_id"`
	//	}, 0)
	//	for _, m := range sop.SopModel {
	//		sopModels = append(sopModels, struct {
	//			SopID   uint `json:"sop_id"`
	//			ModelID uint `json:"model_id"`
	//		}{SopID: m.SopID, ModelID: m.ModelID})
	//	}
	//
	//	sopBytes, err := json.Marshal(sops)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	modelBytes, err := json.Marshal(sopModels)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	if err := analyze.Create(string(sopBytes), string(modelBytes)); err != nil {
	//		fmt.Println(err)
	//	}
	//}(sop)

	c.Succ("操作成功")
}

// sop匹配
func (c *SopController) PostMatch() {

	ids := c.Ctx.PostValue("ids")

	var idArr []uint
	if err := json.Unmarshal([]byte(ids), &idArr); err != nil {
		c.Err("参数格式异常")
		return
	}

	apsService := new(service.ApsService)
	sopService := new(service.SopService)

	apsList := apsService.GetApsListByIds(idArr)
	sopList := sopService.GetAll()

	var data = make([]interface{}, 0)
	for _, aps := range apsList {
		var sopCurrent model.Sop
		var sopSimilar model.Sop
		for _, sop := range sopList {
			if aps.ProductModel.ProductID == sop.ProductID {
				sopSimilar = sop
				if aps.CraftID == sop.CraftID {
					sopCurrent = sop
				}
				break
			}
		}
		if sopCurrent.ID != 0 {
			sopSimilar = sopCurrent
		}
		if sopSimilar.ID == 0 {
			continue
		}

		data = append(data, map[string]interface{}{
			"sopId": sopSimilar.ID,
			"apsId": aps.ID,
		})
	}

	c.Succ("ok", data)
}

// sop下发
func (c *SopController) PostIssued() {

	apsSet := c.Ctx.PostValue("apsSet")

	var set []struct {
		ApsID uint `json:"apsId"`
		SopID uint `json:"sopId"`
	}

	if err := json.Unmarshal([]byte(apsSet), &set); err != nil {
		c.Err("数据解析失败")
		return
	}

	sopIds := make([]uint, 0)
	apsIds := make([]uint, 0)
	for _, s := range set {
		sopIds = append(sopIds, s.SopID)
		apsIds = append(apsIds, s.ApsID)
	}

	sops := new(service.SopService).GetSopListByIds(sopIds)
	for _, s := range sops {
		if s.Status != model.SOP_PASS {
			c.Err("sop审核未通过")
			return
		}
	}

	apsService := new(service.ApsService)
	apsList := apsService.GetApsListByIds(apsIds)

	var apsData = make([]model.Aps, 0)
	var orderList []model.ApsOrder
	for _, aps := range apsList {

		orders := aps.ApsOrder
		var sopId uint
		for _, s := range set {
			if s.ApsID == aps.ID {
				sopId = s.SopID
				break
			}
		}

		var sop model.Sop
		for _, sv := range sops {
			if sv.ID == sopId {
				sop = sv
				break
			}
		}

		var orderData = make([]model.ApsOrder, 0)
		for _, order := range orders {
			var sts = false
			for _, process := range sop.SopProcess {
				if order.ProcessID == process.ProcessID {
					order.SopProcessID = process.ID
					order.Status = model.APS_STATUS_START
					sts = true
					orderData = append(orderData, order)
					break
				}
			}
			if sts == false {
				c.Err("下发异常")
				return
			}
		}

		if len(orderData) > 0 {
			orderList = append(orderList, orderData...)
		}
		aps.SopID = sop.ID
		aps.Status = model.APS_STATUS_START
		apsData = append(apsData, aps)
	}

	if err := apsService.UpdateApsAndOrder(apsData, orderList); err != nil {
		c.Err("操作失败")
		return
	}

	c.Succ("操作成功")
}
