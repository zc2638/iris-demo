package controller

import (
	"github.com/kataras/iris"
	"sop/model"
	"sop/service"
)

type WebController struct{ Base }

// 工单列表
func (c *WebController) GetOrders() {

	orderAll := new(service.ApsService).GetOrderAll()
	var list []model.ApsOrder
	for _, v := range orderAll {
		if v.Status == model.APS_STATUS_START {
			list = append(list, v)
		}
	}

	var data = make([]map[string]interface{}, 0)
	for _, v := range list {
		data = append(data, map[string]interface{}{
			"id": v.ID,
			// 工单号
			"orderId": v.OrderID,
			// 工位名称
			"station": v.StationName,
			// 工序id
			"processId": v.SopProcessID,
			// 工序号
			"processSort": v.SopProcess.Sort,
			// 工序名称
			"processName": v.SopProcess.Title,
		})
	}

	c.Succ("", data)
}

// 工单详情
func (c *WebController) GetOrderShow() {

	processID, err := c.Ctx.URLParamInt("processId")
	if err != nil {
		c.Err("解析异常")
		return
	}

	process := new(service.SopService).GetProcessByID(processID)
	if process.ID == 0 {
		c.Err("不存在的工序")
		return
	}

	c.Succ("", iris.Map{
		"sopImg": process.Imgs,
	})
}