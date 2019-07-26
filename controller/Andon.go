package controller

import "sop/service"

type AndonController struct{ Base }

func (c *AndonController) GetList() {

	all := new(service.AndonService).GetAll()

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {
		data = append(data, map[string]interface{}{
			"triggerTime": v.TriggerTime, // 触发时间
			"station": v.StationName, // 工位名称
			"content": v.Content, // 内容
			"type": v.Type, // 类型
			"code": v.Code, // 业务号
			"info": v.Information, // 业务信息
		})
	}
	c.Succ("", data)
}