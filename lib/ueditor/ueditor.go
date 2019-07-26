package ueditor

import (
	"github.com/dazhenghu/util/fileutil"
	"net/http"
	"path/filepath"
	"sop/lib/gueditor"
)

var uedService *gueditor.Service

func init() {
	rootPath, _    := fileutil.GetCurrentDirectory()
	//configFilePath := filepath.Join(rootPath, "ueditor-config.json") // 设置自定义配置文件路径

	rootPath      = filepath.Join(rootPath, "../") // 设置项目根目录
	uedService, _ = gueditor.NewService(nil, nil, rootPath, "")
}

// 获取配置
func GetConfig() *gueditor.Config {
	return uedService.Config()
}

// 上传图片
func UploadImage(r *http.Request) (*gueditor.ResFileInfoWithState, error) {
	return uedService.Uploadimage(r)
}