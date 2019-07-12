package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main(){


	app := iris.New()
	// 错误恢复正常中间件
	app.Use(recover.New())
	// 日志中间件
	app.Use(logger.New())
	// 设置日志级别
	app.Logger().SetLevel("info")

	// 错误处理
	app.OnAnyErrorCode(func(ctx iris.Context){
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Text("not found the router")
	})

	//设置静态文件
	app.StaticWeb("/static","./static")

	// 注册视图
	tmpl := iris.HTML("./templates", ".html")
	tmpl.Delims("{%","%}")
	tmpl.Layout("layout.html")
	app.RegisterView(tmpl)

	// 路由
	app.Get("/{name:string regexp(^[a-z]+\\.html)}",IndexHandler)

	// 启动
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

func IndexHandler(ctx iris.Context){
	ctx.ViewData("Title","IRIS")
	name := ctx.Params().Get("name")
	fmt.Println("---> ",name)
	ctx.ViewData("Tab",name)
	ctx.View("home.html")
}

func GyHandler(ctx iris.Context){
	ctx.ViewData("Tab","gy.html")
	ctx.View("gy.html")
}
