package controller

import (
	"sop/model"
	"sop/service"
	"strconv"
	"strings"
)

type ApsController struct{ Base }

func (c *ApsController) GetList() {

	productID := c.Ctx.URLParam("productId")
	modelID := c.Ctx.URLParam("modelId")

	var modelIds = make([]uint, 0)
	if modelID == "" {
		models := new(service.ProductService).GetModelAllByProductID(productID)
		for _, v := range models {
			modelIds = append(modelIds, v.ID)
		}
	} else {
		mid, err := strconv.Atoi(modelID)
		if err != nil {
			c.Err("产品型号解析失败")
			return
		}
		modelIds = append(modelIds, uint(mid))
	}

	var all []model.Aps
	if productID == "" {
		all = new(service.ApsService).GetAll()
	} else {
		all = new(service.ApsService).GetListByModelIds(modelIds)

	}

	var productIds = make([]uint, 0)
	for _, v := range all {
		var sts = false
		for _, pid := range productIds {
			if v.ProductModel.ProductID == pid {
				sts = true
				break
			}
		}
		if sts == false {
			productIds = append(productIds, v.ProductModel.ProductID)
		}
	}
	products := new(service.ProductService).GetProductListByIds(productIds)

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {
		if v.Status == model.APS_STATUS_START {
			continue
		}
		if v.SopID != 0 {
			continue
		}

		var product model.Product
		for _, p := range products {
			if p.ID == v.ProductModel.ProductID {
				product = p
			}
		}

		data = append(data, map[string]interface{}{
			"id": v.ID,
			// 作业计划号
			"jobPlanNumber": v.JobPlanNumber,
			// 产品名称
			"productName": product.Name,
			// 产品型号
			"productModelName": v.ProductModel.Name,
			// 产线编号
			"serialNo": v.SerialNo,
		})
	}

	c.Succ("", data)
}

// 手动匹配 —— sop列表
func (c *ApsController) GetSopList() {

	// todo type 和 model匹配
	apsID, err := c.Ctx.URLParamInt("apsId")
	if err != nil {
		c.Err("作业计划解析失败")
		return
	}

	aps := new(service.ApsService).GetApsByID(apsID)
	if aps.ID == 0 {
		c.Err("不存在的作业计划")
		return
	}

	sopList := new(service.SopService).GetSopListByProductID(aps.ProductModel.ProductID)

	var data = make([]map[string]interface{}, 0)
	for _, sop := range sopList {

		models := make([]string, 0)
		modelIds := make([]uint, 0)
		if sop.SopModel != nil && len(sop.SopModel) > 0 {
			for _, m := range sop.SopModel {
				models = append(models, m.Name)
				modelIds = append(modelIds, m.ModelID)
			}
		}

		data = append(data, map[string]interface{}{
			"id": sop.ID,
			// sop名称
			"title": sop.Title,
			// 版本号
			"version": sop.Version,
			// 产品id
			"productId": sop.ProductID,
			// 产品名称
			"productName": sop.Product.Name,
			// 产品型号id集合
			"productTypeIds": modelIds,
			// 产品类型
			"productType": strings.Join(models, " "),
			// 工艺方案id
			"craftId": sop.CraftID,
			// 工艺方案名称
			"craftName": sop.Craft.Name,
			// 工序
			"process": sop.SopProcess,
		})
	}

	c.Succ("", data)
}

// sop手动下发
func (c *ApsController) PostSopIssued() {

	apsId := c.Ctx.PostValue("apsId")
	sopId := c.Ctx.PostValue("sopId")

	apsService := new(service.ApsService)
	aps := apsService.GetApsByID(apsId)
	if aps.ID == 0 {
		c.Err("不存在的作业计划")
		return
	}

	sop := new(service.SopService).GetSopByID(sopId)
	if sop.ID == 0 {
		c.Err("不存在的sop")
		return
	}

	apsCraft := new(service.CraftService).GetCraftByID(aps.CraftID)

	orders := aps.ApsOrder
	if len(sop.SopProcess) != len(orders) {
		c.Err("工单数量与sop工序不匹配")
		return
	}

	var orderData = make([]model.ApsOrder, 0)
	for _, order := range orders {

		var item model.CraftItem
		for _, c := range apsCraft.CraftItem {
			if c.ID == order.ProcessID {
				item = c
				break
			}
		}
		if item.ID == 0 {
			c.Err("工单异常")
			return
		}

		var processId uint
		for _, p := range sop.SopProcess {
			if item.ID == p.ProcessID {
				processId = p.ID
				break
			}
		}

		if processId == 0 {
			for _, p := range sop.SopProcess {
				if item.Sort == p.Sort {
					processId = p.ID
					break
				}
			}
			if processId == 0 {
				c.Err("工序步骤异常")
				return
			}
		}

		order.SopProcessID = processId
		order.Status = model.APS_STATUS_START
		orderData = append(orderData, order)
	}

	aps.SopID = sop.ID
	aps.Status = model.APS_STATUS_START

	if err := apsService.UpdateApsAndOrder([]model.Aps{aps}, orderData); err != nil {
		c.Err("操作失败")
		return
	}
	c.Succ("操作成功")
}
