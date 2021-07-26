package router

import (
	"github.com/kataras/iris/v12"
	"github.com/wanhuasong/uniportal/asset"
	. "github.com/wanhuasong/uniportal/controllers"
)

func Run() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Binary(asset.Asset, asset.AssetNames)
	app.RegisterView(tmpl)

	app.Get("/uniportal", ViewUniportalLogin)
	app.Get("/home", ViewUniportalHome)
	app.Get("/logout", Logout)

	api := app.Party("/api")
	api.Post("/uniportal/login", UniportalLogin)
	api.Post("/iam/token/create", GenerateIamToken)
	api.Post("/idaas/token/auth", IdaasAuth)
	api.Post("/mdm/user/query", GetUserDetail)

	app.Listen(":14501")
}
