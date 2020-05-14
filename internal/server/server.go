package server

import (
	"SerialTest/configs"
	"SerialTest/internal/serialcom"
	"github.com/kataras/iris/v12"
)

func registerHandler(ctx iris.Context) {
	type registerType struct {
		Listen string `url:"url"`
	}
	var t registerType
	err := ctx.ReadQuery(&t)
	if err != nil && !iris.IsErrPath(err) {
		ctx.JSON(iris.Map{
			"code":    500,
			"message": "参数解析错误",
		})
	} else {
		configs.ClinetListened["token"] = t.Listen
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "注册成功",
		})
	}
}

func writerHandler(ctx iris.Context) {
	type writeType struct {
		Content string `url:"content"`
	}
	var t writeType
	err := ctx.ReadQuery(&t)
	if err != nil && !iris.IsErrPath(err) {
		ctx.JSON(iris.Map{
			"code":    500,
			"message": "参数解析错误",
		})
	} else {
		_, err := serialcom.ComWrite([]byte(t.Content))
		if err == nil {
			ctx.JSON(iris.Map{
				"code":    200,
				"message": "指令写入成功",
			})
		} else {
			ctx.JSON(iris.Map{
				"code":    500,
				"message": "指令写入失败",
			})
		}
	}
}

func logHandler(ctx iris.Context) {
	ctx.SendFile(configs.LogPathConfig, configs.LogPathConfig)
}

func RuntimeServer() {
	app := iris.New()
	app.Get("/serialtest/v1/register", registerHandler)
	app.Get("/serialtest/v1/writer", writerHandler)
	app.Get("/serialtest/v1/log", logHandler)
	app.Listen(configs.ServerConfig.ServerAddr)
}
