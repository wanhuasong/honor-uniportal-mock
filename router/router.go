package router

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/wanhuasong/uniportal/asset"
	. "github.com/wanhuasong/uniportal/controllers"
)

func Run() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Binary(asset.Asset, asset.AssetNames)
	app.RegisterView(tmpl)

	app.Get("/uniportal", func(ctx iris.Context) {
		redirect := ctx.URLParamEscape("redirect")
		log.Printf("redirect: %s", redirect)
		ctx.ViewData("redirect", redirect)
		ctx.View("uniportal-login-form.html")
	})

	api := app.Party("/api")
	api.Post("/uniportal/login", UniportalLogin)
	api.Post("/iam/token/create", GenerateIamToken)
	api.Post("/idaas/token/auth", IdaasAuth)
	api.Post("/mdm/user/query", GetUserDetail)

	app.Listen(":14501")
}
