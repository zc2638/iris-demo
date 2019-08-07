package AICheck

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"sop/lib/curl"
)

const domain = "http://esop-imagex.shanghai.cosmoplat.com"

type InfoResult struct {
	Info struct {
		Size struct {
			W int `json:"w"`
			H int `json:"h"`
		} `json:"size"`
		Position map[string]interface{} `json:"position"`
	} `json:"info"`
	Img string `json:"img"`
}

type PosResult struct {
	Info map[string]interface{} `json:"info"`
	Img  string                 `json:"img"`
}

// 标准图像信息识别
func CheckInfo(file io.Reader, filename string, data map[string]string) (result InfoResult, err error) {

	const url = domain + "/check/info"

	h := curl.HttpReq{
		Url: url,
		FormData: curl.FormData{
			File: map[string]curl.FileInfo{
				"picture": {filename, file},
			},
			Params: data,
			//Params: map[string]string{
			//	"colors": `{"red":[[[ 0, 43, 46],[ 10, 255, 255]], [[156,  43,  46],[180, 255, 255]]]}`,
			//},
		},
	}

	res, err := h.PostForm()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}

// 图像检测
func CheckPos(file io.Reader, info *multipart.FileHeader, data map[string]string) (result PosResult, err error) {

	const url = domain + "/check/pos"

	h := curl.HttpReq{
		Url: url,
		FormData: curl.FormData{
			File: map[string]curl.FileInfo{
				"picture": {info.Filename, file},
			},
			Params: data,
			//Params: map[string]string{
			//	"colors": `{"red":{ "ranges": [[[0, 43, 46],[10, 255, 255]], [[156,  43,  46],[180, 255, 255]]],"rect": {"x": 480, "y": 365, "w": 216, "h": 233},"scale": {"x":0.1, "y":0.1}}}`,
			//},
		},
	}

	res, err := h.PostForm()
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	return
}
