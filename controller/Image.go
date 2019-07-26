package controller

import (
	"encoding/json"
	"sop/lib/qiniu"
	"strconv"
	"sync"
)

type ImageController struct{ Base }

// 上传图片(base64)
func (c *ImageController) PostUploads() {

	imageBase64SetStr := c.Ctx.PostValue("imageBase64Set")
	if imageBase64SetStr == "" {
		c.Err("请选择上传的图片")
		return
	}

	var imageBase64Set []string
	if err := json.Unmarshal([]byte(imageBase64SetStr), &imageBase64Set); err != nil {
		c.Err("解析失败")
		return
	}

	var wg sync.WaitGroup
	var imageUrlSet = make([]string, len(imageBase64Set))
	for k, v := range imageBase64Set {
		wg.Add(1)
		go func(imageBase64 string, key int) {
			imageUrl, err := qiniu.UploadBase64(imageBase64)
			if err == nil {
				imageUrlSet[key] = imageUrl
			}
			wg.Done()
		}(v, k)
	}
	wg.Wait()

	for k, v := range imageUrlSet {
		if v == "" {
			c.Err("图片" + strconv.Itoa(k + 1) + "上传失败")
			return
		}
	}
	c.Succ("上传成功", imageUrlSet)
}

// 上传图片(数据流)
func (c *ImageController) PostUploadStream() {

	// 限制 10MB
	c.Ctx.SetMaxRequestBodySize(10 << 20)
	file, info, err := c.Ctx.FormFile("image")
	if err != nil {
		c.Err(err.Error())
		return
	}

	imageUrl, err := qiniu.Upload(file, info)
	if err != nil {
		c.Err(err.Error())
		return
	}
	c.Succ("上传成功", imageUrl)
}