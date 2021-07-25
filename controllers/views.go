package controllers

import (
	"log"

	"github.com/kataras/iris/v12"
)

func ViewUniportalLogin(ctx iris.Context) {
	uid := ctx.GetCookie(AuthCookieUid)
	if uid != "" {
		ctx.Redirect("/home")
		return
	}

	redirect := ctx.URLParamEscape("redirect")
	log.Printf("Found redirect url: %s", redirect)
	ctx.ViewData("redirect", redirect)
	ctx.View("uniportal-login-form.html")
}

func ViewUniportalHome(ctx iris.Context) {
	ctx.View("home.html")
}
