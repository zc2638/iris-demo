package controller

import "sop/service"

type ProductController struct{ Base }

// 获取产品列表
func (c *ProductController) GetAll() {

	all := new(service.ProductService).GetAll()

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {
		data = append(data, map[string]interface{}{
			// 产品id
			"value": v.ID,
			// 产品名称
			"label": v.Name,
		})
	}

	c.Succ("", data)
}

// 获取产品型号列表
func (c *ProductController) GetModel() {

	productID := c.Ctx.URLParam("product_id")
	all := new(service.ProductService).GetModelAllByProductID(productID)

	var data = make([]map[string]interface{}, 0)
	for _, v := range all {
		data = append(data, map[string]interface{}{
			// 产品型号id
			"value": v.ID,
			// 产品型号名称
			"label": v.Name,
		})
	}

	c.Succ("", data)
}
