package controller

import (
	"sop/service"
)

type CraftController struct{ Base }

// 获取所有工艺方案
func (c *CraftController) GetAll() {

	all := new(service.CraftService).GetAll()

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {

		table := make([]map[string]interface{}, 0)
		if v.CraftItem != nil && len(v.CraftItem) > 0 {
			for _, item := range v.CraftItem {
				table = append(table, map[string]interface{}{
					// 工艺工序id
					"sopStepID": item.ID,
					// 工艺工序名称
					"sopStepName": item.Name,
					// 状态
					"sopStatus": item.Status,
					// 防差错图片
					"pyInfo":    item.CheckImg,
				})
			}
		}

		data = append(data, map[string]interface{}{
			// 工艺方案id
			"value": v.ID,
			// 工艺方案名称
			"label": v.Name,
			"table": table,
		})
	}
	c.Succ("", data)
}
