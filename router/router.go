package router

import (
	"github.com/kataras/iris/v12"
	. "github.com/wanhuasong/uniportal/controllers"
)

func Run() {
	app := iris.New()

	app.Get("/uniportal", UniportalLoginForm)

	api := app.Party("/api")
	api.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	api.Post("/iam/auth/token", GenerateIamToken)
	api.Post("/idaas/auth", IdaasAuth)

	app.Listen(":8080")
}
