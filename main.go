package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"sop/controller"
	_ "sop/lib/database"
	_ "sop/lib/face"
	"sop/lib/ueditor"
	"sop/middleware"
)

func main() {

	app := iris.Default()

	app.Use(middleware.Cors())

	app.Get("/", func(ctx iris.Context) {
		_, _ = ctx.Text("Hello World!")
	})

	// UEditor配置
	app.Any("/staticUE", UEditor)

	mvc.New(app.Party("/home")).Handle(new(controller.HomeController))
	mvc.New(app.Party("/face")).Handle(new(controller.FaceController))

	api := app.Party("/api")
	{
		mvc.New(api.Party("/admin")).Handle(new(controller.AdminController))
		mvc.New(api.Party("/user")).Handle(new(controller.UserController))
		mvc.New(api.Party("/sop")).Handle(new(controller.SopController))
		mvc.New(api.Party("/product")).Handle(new(controller.ProductController))
		mvc.New(api.Party("/craft")).Handle(new(controller.CraftController))
		mvc.New(api.Party("/image")).Handle(new(controller.ImageController))
		mvc.New(api.Party("/andon")).Handle(new(controller.AndonController))
		mvc.New(api.Party("/aps")).Handle(new(controller.ApsController))
	}

	if err := app.Run(iris.Addr(":8080")); err != nil {
		log.Fatal(err)
	}
}

func UEditor(ctx iris.Context) {

	action := ctx.URLParam("action")

	var res interface{}
	switch action {
	case "config":
		// config接口
		res = ueditor.GetConfig()
	case "uploadimage":
		// 上传图片
		res, _ = ueditor.UploadImage(ctx.Request())
	//case "uploadscrawl":
	//	// 上传涂鸦
	//	ued.uploadScrawl(context)
	//case "uploadvideo":
	//	// 上传视频
	//	ued.uploadVideo(context)
	//case "uploadfile":
	//	// 上传附件
	//	ued.uploadfile(context)
	//case "listfile":
	//	// 查询上传的文件列表
	//	ued.listFile(context)
	//case "listimage":
	//	// 查询上传的图片列表
	//	ued.listImage(context)
	}

	_, _ = ctx.JSON(res)
}