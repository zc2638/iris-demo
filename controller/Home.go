package controller

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"sop/lib/AICheck"
	"sop/model"
	"sop/service"
)

type HomeController struct{ Base }

func (c *HomeController) Get() interface{} {
	return iris.Map{
		"message": "Hello Sop!",
		"content": c.Ctx.Request().RequestURI,
	}
}

// 工单列表
func (c *HomeController) GetOrders() {

	uid := c.Ctx.URLParam("uid")
	station := c.Ctx.URLParamDefault("station", "1")

	user := new(service.UserService).GetUserByID(uid)
	if user.ID == 0 {
		c.Err("不存在的用户")
		return
	}

	orders := new(service.ApsService).GetOrdersByUidAndStation(user.ID, station)

	var modelIds = make([]uint, 0)
	for _, v := range orders {
		modelIds = append(modelIds, v.Aps.ModelID)
	}

	models := new(service.ProductService).GetModelByIDs(modelIds)
	var productMap = make(map[uint]model.Product)
	for _, m := range models {
		productMap[m.ID] = m.Product
	}

	var orderData = make([]map[string]interface{}, 0)
	for _, v := range orders {
		var productName string
		product, ok := productMap[v.Aps.ModelID]
		if ok {
			productName = product.Name
		}

		var qualities = make([]map[string]interface{}, 0)
		if v.ApsOrderQuality != nil && len(v.ApsOrderQuality) > 0 {
			for _, v := range v.ApsOrderQuality {
				qualities = append(qualities, map[string]interface{}{
					"id": v.ID,
					// 工件编号
					"pieceNo": v.PieceNo,
					// 质检结果
					"result": v.Result,
					// 备注
					"remark": v.Remark,
				})
			}
		}

		orderData = append(orderData, map[string]interface{}{
			"id": v.ID,
			// 工单号
			"orderId": v.OrderID,
			// 产品名称
			"productName": productName,
			// 工序名称
			"sopProcessTitle": v.SopProcess.Title,
			// 防差错图片
			"checkImg": v.SopProcess.CheckImg,
			// 加工数量
			"processQuantity": v.Total,
			// 完工数量
			"finishQuantity": v.Num,
			// 合格数量
			"qualifiedProductQuantity": v.CompleteNum,
			// 工单质检
			"QItable": qualities,
		})
	}

	c.Succ("", map[string]interface{}{
		"orders": orderData,
		"user": map[string]interface{}{
			"id": user.ID,
			// 用户名
			"name": user.Name,
			// 岗位职责
			"summary": user.Summary,
			// 操作权限
			"role": user.Role,
			// 照片
			"faceImage": user.FaceImage,
			// 工位名称
			"station": "工位1",
		},
	})
}

// 看板工单详情
func (c *HomeController) PostShow() {

	orderID := c.Ctx.PostValue("id")

	order := new(service.ApsService).GetOrderByID(orderID)
	sop := new(service.SopService).GetSopByID(order.SopProcess.SopID)
	productModel := new(service.ProductService).GetModelByID(order.Aps.ModelID)

	var processSet = make([]map[string]interface{}, 0)
	if sop.SopProcess != nil && len(sop.SopProcess) > 0 {
		for _, v := range sop.SopProcess {
			processSet = append(processSet, map[string]interface{}{
				"id": v.ID,
				// 工序名
				"title": v.Title,
			})
		}
	}

	var qualities = make([]map[string]interface{}, 0)
	if order.ApsOrderQuality != nil && len(order.ApsOrderQuality) > 0 {
		for _, v := range order.ApsOrderQuality {
			qualities = append(qualities, map[string]interface{}{
				"id": v.ID,
				// 工件编号
				"pieceNo": v.PieceNo,
				// 质检结果
				"result": v.Result,
				// 备注
				"remark": v.Remark,
			})
		}
	}

	c.Succ("", iris.Map{
		// 工单号
		"orderId": order.OrderID,
		// 文件编号 = 作业计划号
		"fileCode": order.Aps.JobPlanNumber,
		// 作业计划号
		"jobPlanNumber": order.Aps.JobPlanNumber,
		// 生效日期
		"passAt": sop.PassAt,
		// 工位
		"station": order.StationName,
		// 工序
		"sopProcessTitle": order.SopProcess.Title,
		// sop工序图片
		"sopImg": order.SopProcess.Imgs,
		// 附件信息
		"annex": order.SopProcess.Annex,
		// 防差错图片
		"checkImg": order.SopProcess.CheckImg,
		// 产品名称
		"productName": productModel.Product.Name,
		// 产线编号
		"serialNo": order.Aps.SerialNo,
		// 所有工序
		"sopProcessSet": processSet,
		// 工单质检
		"QItable": qualities,
		// 加工数量
		"processQuantity": order.Total,
		// 完工数量
		"finishQuantity": order.Num,
		// 合格数量
		"qualifiedProductQuantity": order.CompleteNum,
		// 工序id
		"sopProcessId": order.SopProcessID,
		// 版本号
		"version": sop.Version,
		// 雷区预警
		"minefieldImg": order.SopProcess.MinefieldImg,
	})
}

// 修改防差错图片
func (c *HomeController) PostCheckSave() {

	processId := c.Ctx.PostValue("processId")
	checkImg := c.Ctx.PostValue("checkImg")

	sopService := new(service.SopService)
	process := sopService.GetProcessByID(processId)
	if process.ID == 0 {
		c.Err("不存在的工序")
		return
	}

	var checkImgData struct {
		Item string `json:"item"`
		Bg   string `json:"bg"`
	}

	if err := json.Unmarshal([]byte(process.CheckImg), &checkImgData); err != nil {
		c.Err("解析失败")
		return
	}
	checkImgData.Item = checkImg

	b, err := json.Marshal(checkImgData)
	if err != nil {
		c.Err("编码失败")
		return
	}
	process.CheckImg = string(b)

	if err := sopService.UpdateProcessOne(process); err != nil {
		c.Err(err.Error())
		return
	}
	c.Succ("操作成功")
}

// AI纠错接口
func (c *HomeController) PostCheckImage() {

	processID := c.Ctx.PostValue("processId")
	file, info, err := c.Ctx.FormFile("image")
	if err != nil {
		c.Err(err.Error())
		return
	}

	process := new(service.SopService).GetProcessByID(processID)
	if process.ID == 0 {
		c.Err("不存在的工序")
		return
	}

	var checkImg struct {
		Item string `json:"item"`
		Bg   string `json:"bg"`
	}
	if err := json.Unmarshal([]byte(process.CheckImg), &checkImg); err != nil {
		c.Err("防差错图片解析失败")
		return
	}

	checkService := new(service.CheckService)
	check := checkService.GetCheckByUrl(checkImg.Item)
	if check.ID == 0 {
		c.Err("防差错图片基本信息获取失败")
		return
	}

	var colors map[string]interface{}
	if err := json.Unmarshal([]byte(check.Colors), &colors); err != nil {
		c.Err("防差错图片基本信息解析失败")
		return
	}

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		c.Err("图片信息异常")
		return
	}

	if _, ok := colors["scale"]; !ok {
		infoRes, err := AICheck.CheckInfo(bytes.NewReader(fileByte), info, map[string]string{
			"colors": check.Colors,
		})
		if err != nil {
			c.Err(err.Error())
			return
		}

		size, err := json.Marshal(infoRes.Info.Size)
		if err != nil {
			c.Err("尺寸解析失败")
			return
		}
		check.Size = string(size)

		var colorData = make(map[string]struct {
			Ranges interface{} `json:"ranges"`
			Rect   interface{} `json:"rect"`
			Scale  interface{} `json:"scale"`
		})
		position := infoRes.Info.Position
		for k, v := range position {
			if cv, ok := colors[k]; ok {

				colorData[k] = struct {
					Ranges interface{} `json:"ranges"`
					Rect   interface{} `json:"rect"`
					Scale  interface{} `json:"scale"`
				}{cv, v, map[string]interface{}{
					"x": 0.1,
					"y": 0.1,
				}}
			}
		}

		colorByte, err := json.Marshal(colorData)
		if err != nil {
			c.Err("图像组合失败")
			return
		}

		check.Colors = string(colorByte)
		if err := checkService.UpdateOne(check); err != nil {
			c.Err(err.Error())
			return
		}
	}

	res, err := AICheck.CheckPos(bytes.NewReader(fileByte), info, map[string]string{
		"colors": check.Colors,
		"size":   check.Size,
	})
	if err != nil {
		c.Err(err.Error())
		return
	}

	//infoArr := strings.Split(res.Img, ",")
	//mime := strings.Replace(strings.Replace(infoArr[0], ";base64", "", -1), "data:image/", "", -1)
	//
	//f, err := os.OpenFile("test."+mime, os.O_WRONLY|os.O_CREATE, 0600)
	//if err != nil {
	//	c.Err(err.Error())
	//	return
	//}
	//defer f.Close()

	//b, err := base64.StdEncoding.DecodeString(infoArr[1])
	//if err != nil {
	//	c.Err(err.Error())
	//	return
	//}
	//
	//if _, err := f.Write(b); err != nil {
	//	c.Err(err.Error())
	//	return
	//}

	c.Succ("ok", res)
}
