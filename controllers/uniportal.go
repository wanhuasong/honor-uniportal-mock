package controllers

import (
	"log"

	"github.com/kataras/iris/v12"
)

func UniportalLogin(ctx iris.Context) {
	ctx.SetCookieKV("authmethod", "authpwd")
	ctx.SetCookieKV("hwsso_uniportal", "lZKjdGPj4lKUaUh3a_a9i1OF_b_bOzTRhsdb9pduCH54dudRbqSDi99svSQcHtwS_bp8JX3E0eOkK9VWfVFJO3Jo4Tcm_bi2XpIjRFL_bewDxL8kgbu1w0sivkh_b8vMAUloI_bvpDxwH8_aarEU0tVZmnc7XkT84QZl3XqZjUxQmLvWEi57cqOi5kmiDKMH9k8wFXexESLf7NX5aqSMreMLt8r14lWPzZ5tXTA4Nkqxr9HN3g_a0zKp4jBFE8_a_be1Tru1Hl2js7ESLFmrtfsbb6bW9aT_aXQV6iyzb9_buFyOFNfKgnm2ZltIrU3FbxFZA3w6vJCnTggzPBeq_anYDg9_bWZ9ylISEg_c_c")
	ctx.SetCookieKV("hwssotinter", "D7-AA-90-56-3B-3B-12-87-30-7C-5D-5B-2B-A7-15-A3")
	ctx.SetCookieKV("hwssotinter3", "27527561300776")
	ctx.SetCookieKV("uid", "31303130F5553906EFFA0D3E5B2370C0A75080E2C3F7B139783FBFDEA06B2941A794BE623E1C986E31AABE17D628BB2ED7AE1755")

	email := ctx.FormValue("email")
	log.Printf("login email: %s", email)
	password := ctx.FormValue("password")
	log.Printf("login passoword: %s", password)

	redirect := ctx.FormValue("redirect")
	ctx.ViewData("redirect", redirect)
	ctx.View("login-response.html")
}
