package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sop/lib/face"
	"sop/lib/qiniu"
	"sop/model"
	"sop/service"
)

type FaceController struct{ Base }

// 人脸识别登陆
func (c *FaceController) PostLogin() {

	station := c.Ctx.FormValue("station") // 工位id
	file, info, err := c.Ctx.FormFile("faceImage")
	if err != nil {
		c.Err(err.Error())
		return
	}
	defer file.Close()

	result, err := face.Search(file, info)
	if err != nil {
		c.Err(err.Error())
		return
	}

	if result.ErrorMessage != "" {
		c.Err(result.ErrorMessage)
		return
	}
	if result.Results == nil || len(result.Results) == 0 {
		c.Err("未识别到人脸")
		return
	}

	var user model.User
	//if result.Results[0].Confidence < result.Thresholds.OneE5 {
	//	// 识别失败，则指定默认用户
	//	user = new(service.UserService).GetUserByID(1)
	//} else {
	//	faceToken := result.Results[0].FaceToken
	//	fmt.Println(faceToken)
	//	user = new(service.UserService).GetUserByFaceToken(faceToken)
	//}

	faceToken := result.Results[0].FaceToken
	fmt.Println(faceToken)
	user = new(service.UserService).GetUserByFaceToken(faceToken)
	if user.ID == 0 {
		c.Err("不存在的用户")
		return
	}

	var orders []model.ApsOrder
	if user.Name == "红兵" || user.Name == "祥宁" {
		ordersAll := new(service.ApsService).GetOrderAll()
		for _, v := range ordersAll {
			if v.Status == model.APS_STATUS_START {
				orders = append(orders, v)
			}
		}
	} else {
		orders = new(service.ApsService).GetOrdersByUidAndStation(user.ID, station)
	}

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

// 添加基础图片
func (c *FaceController) PostAdd() {

	uid := c.Ctx.FormValue("uid")
	file, info, err := c.Ctx.FormFile("faceImage")
	if err != nil {
		c.Err(err.Error())
		return
	}
	defer file.Close()

	userService := new(service.UserService)
	user := userService.GetUserByID(uid)
	if user.ID == 0 {
		c.Err("不存在的用户")
		return
	}

	fileBuf, err := ioutil.ReadAll(file)
	if err != nil {
		c.Err("图片信息获取异常")
		return
	}

	detectResult, err := face.DetectImage(bytes.NewReader(fileBuf), info)
	if err != nil {
		c.Err(err.Error())
		return
	}

	if detectResult.ErrorMessage != "" {
		c.Err(detectResult.ErrorMessage)
		return
	}

	if detectResult.Faces == nil || len(detectResult.Faces) == 0 {
		c.Err("未识别到人脸")
		return
	}

	imageUrl, err := qiniu.Upload(bytes.NewReader(fileBuf), info)
	if err != nil {
		fmt.Println(err)
		c.Err("图片保存失败")
		return
	}

	user.FaceImage = imageUrl
	user.FaceToken = detectResult.Faces[0].FaceToken
	fmt.Println(detectResult.Faces[0])
	fmt.Println(user.FaceToken)
	if n := userService.UpdateOne(user); n != 1 {
		c.Err("匹配结果保存失败")
		return
	}

	result, err := face.AddFace(user.FaceToken)
	if err != nil {
		c.Err(err.Error())
		return
	}
	if result.ErrorMessage != "" {
		c.Err(result.ErrorMessage)
		return
	}

	c.Succ("操作成功")
}
