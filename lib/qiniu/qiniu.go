package qiniu

import (
	"context"
	"github.com/kataras/iris/core/errors"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

const (
	ACCESS_KEY     = "HVY9Gi36B-_aqZxP-PwfAXM1Iry1d5E-StoaV9_-"
	SECRET_KEY     = "4uyHmtZVtyXS0SbceqYQDeaTXSVWbmvMwNNRvpO6"
	BUCKET         = "sop"
	DEFAULT_DOMAIN = "puh01tec3.bkt.clouddn.com"
)

// 上传base64图片
func UploadBase64(base64Str string) (string, error) {

	arr := strings.Split(base64Str, ",")
	if len(arr) != 2 {
		return "", errors.New("base64格式异常")
	}

	base64Data := arr[1]
	mime := strings.Replace(strings.Replace(arr[0], ";base64", "", -1), "data:image/", "", -1)
	key := strconv.Itoa(int(time.Now().UnixNano()/1e3)) + "." + mime

	putPolicy := storage.PutPolicy{
		Scope: BUCKET + ":" + key,
	}

	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong // 空间对应的机房
	cfg.UseHTTPS = false            // 是否使用https域名
	cfg.UseCdnDomains = false       // 上传是否使用CDN上传加速

	var ret storage.PutRet
	base64Uploader := storage.NewBase64Uploader(&cfg) // base64上传
	if err := base64Uploader.Put(context.Background(), &ret, upToken, key, []byte(base64Data), &storage.Base64PutExtra{}); err != nil {
		return "", err
	}

	imageUrl := storage.MakePublicURL(DEFAULT_DOMAIN, ret.Key)
	if !strings.Contains(imageUrl, "http://") || !strings.Contains(imageUrl, "https://") {
		imageUrl = "http://" + imageUrl
	}
	return imageUrl, nil
}

// 上传图片流
func Upload(file io.Reader, info *multipart.FileHeader) (string, error) {

	key := strconv.Itoa(int(time.Now().UnixNano()/1e3)) + "-" + info.Filename
	putPolicy := storage.PutPolicy{
		Scope: BUCKET + ":" + key,
	}

	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong // 空间对应的机房
	cfg.UseHTTPS = false            // 是否使用https域名
	cfg.UseCdnDomains = false       // 上传是否使用CDN上传加速

	var ret storage.PutRet
	FormUploader := storage.NewFormUploader(&cfg) // base64上传

	if err := FormUploader.Put(context.Background(), &ret, upToken, key, file, info.Size, &storage.PutExtra{}); err != nil {
		return "", err
	}

	imageUrl := storage.MakePublicURL(DEFAULT_DOMAIN, ret.Key)
	if !strings.Contains(imageUrl, "http://") || !strings.Contains(imageUrl, "https://") {
		imageUrl = "http://" + imageUrl
	}
	return imageUrl, nil
}