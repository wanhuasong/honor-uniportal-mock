package controllers

import (
	"log"

	"github.com/kataras/iris/v12"
)

func ViewUniportalLogin(ctx iris.Context) {
	redirect := ctx.URLParam("redirect")
	log.Printf("Found redirect url: %s", redirect)

	uid := ctx.GetCookie(AuthCookieUid)
	if uid != "" {
		if redirect != "" {
			ctx.ViewData("redirect", redirect)
			ctx.View("login-response.html")
			return
		}
		ctx.Redirect("/home")
		return
	}

	ctx.ViewData("redirect", redirect)
	ctx.View("uniportal-login-form.html")
}

func ViewUniportalHome(ctx iris.Context) {
	uid := ctx.GetCookie(AuthCookieUid)
	if uid == "" {
		ctx.Redirect("/uniportal")
		return
	}
	ctx.View("index.html")
}
