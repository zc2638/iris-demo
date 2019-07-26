package gueditor

import (
    "encoding/json"
)

func init()  {
    GloabConfig = &Config{}
    loadDefaultConfig()
}

var GloabConfig *Config

type Config struct {
    ImageActionName         string   `json:"imageActionName"`
    ImageFieldName          string   `json:"imageFieldName"`
    ImageMaxSize            int      `json:"imageMaxSize"`
    ImageAllowFiles         []string `json:"imageAllowFiles"`
    ImageCompressEnable     bool     `json:"imageCompressEnable"`
    ImageCompressBorder     int      `json:"imageCompressBorder"`
    ImageInsertAlign        string   `json:"imageInsertAlign"`
    ImageURLPrefix          string   `json:"imageUrlPrefix"`
    ImagePathFormat         string   `json:"imagePathFormat"`
    ScrawlActionName        string   `json:"scrawlActionName"`
    ScrawlFieldName         string   `json:"scrawlFieldName"`
    ScrawlPathFormat        string   `json:"scrawlPathFormat"`
    ScrawlMaxSize           int      `json:"scrawlMaxSize"`
    ScrawlURLPrefix         string   `json:"scrawlUrlPrefix"`
    ScrawlInsertAlign       string   `json:"scrawlInsertAlign"`
    SnapscreenActionName    string   `json:"snapscreenActionName"`
    SnapscreenPathFormat    string   `json:"snapscreenPathFormat"`
    SnapscreenURLPrefix     string   `json:"snapscreenUrlPrefix"`
    SnapscreenInsertAlign   string   `json:"snapscreenInsertAlign"`
    CatcherLocalDomain      []string `json:"catcherLocalDomain"`
    CatcherActionName       string   `json:"catcherActionName"`
    CatcherFieldName        string   `json:"catcherFieldName"`
    CatcherPathFormat       string   `json:"catcherPathFormat"`
    CatcherURLPrefix        string   `json:"catcherUrlPrefix"`
    CatcherMaxSize          int      `json:"catcherMaxSize"`
    CatcherAllowFiles       []string `json:"catcherAllowFiles"`
    VideoActionName         string   `json:"videoActionName"`
    VideoFieldName          string   `json:"videoFieldName"`
    VideoPathFormat         string   `json:"videoPathFormat"`
    VideoURLPrefix          string   `json:"videoUrlPrefix"`
    VideoMaxSize            int      `json:"videoMaxSize"`
    VideoAllowFiles         []string `json:"videoAllowFiles"`
    FileActionName          string   `json:"fileActionName"`
    FileFieldName           string   `json:"fileFieldName"`
    FilePathFormat          string   `json:"filePathFormat"`
    FileURLPrefix           string   `json:"fileUrlPrefix"`
    FileMaxSize             int      `json:"fileMaxSize"`
    FileAllowFiles          []string `json:"fileAllowFiles"`
    ImageManagerActionName  string   `json:"imageManagerActionName"`
    ImageManagerListPath    string   `json:"imageManagerListPath"`
    ImageManagerListSize    int      `json:"imageManagerListSize"`
    ImageManagerURLPrefix   string   `json:"imageManagerUrlPrefix"`
    ImageManagerInsertAlign string   `json:"imageManagerInsertAlign"`
    ImageManagerAllowFiles  []string `json:"imageManagerAllowFiles"`
    FileManagerActionName   string   `json:"fileManagerActionName"`
    FileManagerListPath     string   `json:"fileManagerListPath"`
    FileManagerURLPrefix    string   `json:"fileManagerUrlPrefix"`
    FileManagerListSize     int      `json:"fileManagerListSize"`
    FileManagerAllowFiles   []string `json:"fileManagerAllowFiles"`
}

/**
加载默认配置
 */
func loadDefaultConfig() (err error) {
    //filePath, err := getDefaultConfigFile()
    //if err != nil {
    //    return
    //}
    //cnfJson, err := ioutil.ReadFile(filePath)
    //if err != nil {
    //    return
    //}

    cnfJson := []byte(configJson)
    err = json.Unmarshal(cnfJson, GloabConfig)
    return
}

/**
加载配置
 */
func loadConfigFromFile(filePath string) (err error) {

    //exists, err := pathExists(filePath)
    //if !exists {
    //    err = errors.New(fmt.Sprintf("config file not exists:%s", filePath))
    //    return
    //}
    //cnfJson, err := ioutil.ReadFile(filePath)
    //if err != nil {
    //    return
    //}
    cnfJson := []byte(configJson)
    err = json.Unmarshal(cnfJson, GloabConfig)
    return
}

const configJson = `{
  "imageActionName": "uploadimage",
  "imageFieldName": "upfile",
  "imageMaxSize": 2048000,
  "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
  "imageCompressEnable": true,
  "imageCompressBorder": 1600,
  "imageInsertAlign": "none",
  "imageUrlPrefix": "",
  "imagePathFormat": "/ueditor/golang/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}",
  "scrawlActionName": "uploadscrawl",
  "scrawlFieldName": "upfile",
  "scrawlPathFormat": "/ueditor/golang/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}",
  "scrawlMaxSize": 2048000,
  "scrawlUrlPrefix": "",
  "scrawlInsertAlign": "none",
  "snapscreenActionName": "uploadimage",
  "snapscreenPathFormat": "/ueditor/golang/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}",
  "snapscreenUrlPrefix": "",
  "snapscreenInsertAlign": "none",
  "catcherLocalDomain": ["127.0.0.1", "localhost", "img.baidu.com"],
  "catcherActionName": "catchimage",
  "catcherFieldName": "source",
  "catcherPathFormat": "/ueditor/golang/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}",
  "catcherUrlPrefix": "",
  "catcherMaxSize": 2048000,
  "catcherAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
  "videoActionName": "uploadvideo",
  "videoFieldName": "upfile",
  "videoPathFormat": "/ueditor/golang/upload/video/{yyyy}{mm}{dd}/{time}{rand:6}",
  "videoUrlPrefix": "",
  "videoMaxSize": 102400000,
  "videoAllowFiles": [
    ".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
    ".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid"],
  "fileActionName": "uploadfile",
  "fileFieldName": "upfile",
  "filePathFormat": "/ueditor/golang/upload/file/{yyyy}{mm}{dd}/{time}{rand:6}",
  "fileUrlPrefix": "",
  "fileMaxSize": 51200000,
  "fileAllowFiles": [
    ".png", ".jpg", ".jpeg", ".gif", ".bmp",
    ".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
    ".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid",
    ".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso",
    ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".xml"
  ],
  "imageManagerActionName": "listimage",
  "imageManagerListPath": "/ueditor/golang/upload/image/",
  "imageManagerListSize": 20,
  "imageManagerUrlPrefix": "",
  "imageManagerInsertAlign": "none",
  "imageManagerAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
  "fileManagerActionName": "listfile",
  "fileManagerListPath": "/ueditor/golang/upload/file/",
  "fileManagerUrlPrefix": "",
  "fileManagerListSize": 20,
  "fileManagerAllowFiles": [
    ".png", ".jpg", ".jpeg", ".gif", ".bmp",
    ".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
    ".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid",
    ".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso",
    ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".xml"
  ]
}`