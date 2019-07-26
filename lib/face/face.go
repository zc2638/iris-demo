package face

import (
	"encoding/json"
	"fmt"
	"github.com/zctod/go-tool/common/utils"
	"io"
	"log"
	"mime/multipart"
	"sop/lib/curl"
)

const (
	APIKEY      = "OdgqJMq685CrIyeDICSaNQBD1ia8sdPv"
	APISECRET   = "v2rDSvkqsf7MPk1tvOCxZXPbucx-nGjf"
	FACESETID   = "test03"
	FACESETNAME = "Examples"
)

func init() {

	set, err := GetSet()
	if err != nil {
		log.Fatalln(err)
	}
	if set.ErrorMessage != "" {
		log.Fatalln(set.ErrorMessage)
	}

	if len(set.Facesets) > 0 {
		for _, s := range set.Facesets {
			if s.OuterID == FACESETID && s.DisplayName == FACESETNAME {
				return
			}
		}
	}

	res, err := CreateSet()
	if err != nil {
		log.Fatalln(err)
	}
	if res.ErrorMessage != "" {
		log.Fatalln(set.ErrorMessage)
	}
	fmt.Println("创建face集合成功")
}

// 获取所有face集合
func GetSet() (result GetSetResult, err error) {

	const url = "https://api-cn.faceplusplus.com/facepp/v3/faceset/getfacesets"

	dataStruct := CreateSetData{
		ApiKey:    APIKEY,
		ApiSecret: APISECRET,
	}

	dataMap, err := utils.StrcutToMap(dataStruct)
	if err != nil {
		return
	}

	h := curl.HttpReq{
		Url:    url,
		Params: dataMap,
	}

	res, err := h.Post()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}

// 创建face集合
func CreateSet() (result CreateSetResult, err error) {

	const createSetUrl = "https://api-cn.faceplusplus.com/facepp/v3/faceset/create"

	dataStruct := CreateSetData{
		ApiKey:      APIKEY,
		ApiSecret:   APISECRET,
		DisplayName: FACESETNAME,
		OuterID:     FACESETID,
	}

	dataMap, err := utils.StrcutToMap(dataStruct)
	if err != nil {
		return
	}

	h := curl.HttpReq{
		Url:    createSetUrl,
		Params: dataMap,
	}

	res, err := h.Post()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}

// 上传图片到face++
func DetectImage(file io.Reader, info *multipart.FileHeader) (result DetectImageResult, err error) {

	const detectUrl = "https://api-cn.faceplusplus.com/facepp/v3/detect"

	h := curl.HttpReq{
		Url: detectUrl,
		FormData: curl.FormData{
			File: map[string]curl.FileInfo{
				"image_file": {info.Filename, file},
			},
			Params: map[string]string{
				"api_key":    APIKEY,
				"api_secret": APISECRET,
			},
		},
	}

	res, err := h.PostForm()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}

// 往face集合中添加face_token
func AddFace(token string) (result AddFaceResult, err error) {

	const url = "https://api-cn.faceplusplus.com/facepp/v3/faceset/addface"

	dataStruct := AddFaceData{
		ApiKey:     APIKEY,
		ApiSecret:  APISECRET,
		OuterID:    FACESETID,
		FaceTokens: token,
	}

	dataMap, err := utils.StrcutToMap(dataStruct)
	if err != nil {
		return
	}

	h := curl.HttpReq{
		Url:    url,
		Params: dataMap,
	}

	res, err := h.Post()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}

// 查询face
func Search(file multipart.File, info *multipart.FileHeader) (result SearchResult, err error) {

	const searchUrl = "https://api-cn.faceplusplus.com/facepp/v3/search"

	h := curl.HttpReq{
		Url: searchUrl,
		FormData: curl.FormData{
			File: map[string]curl.FileInfo{
				"image_file": {info.Filename, file},
			},
			Params: map[string]string{
				"api_key":    APIKEY,
				"api_secret": APISECRET,
				"outer_id": FACESETID,
			},
		},
	}

	res, err := h.PostForm()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	return
}
