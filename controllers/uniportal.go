package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	AuthCookieAuthMethod     = "authmethod"
	AuthCookieHwssoUniportal = "hwsso_uniportal"
	AuthCookieHwssotinter    = "hwssotinter"
	AuthCookieHwssotinter3   = "hwssotinter3"
	AuthCookieLogFlag        = "logFlag"
	AuthCookieSid            = "sid"
	AuthCookieSip            = "sip"
	AuthCookieUid            = "honor_uid" // cookie uid 被 dev 环境占用

	AuthCookieValueAuthMethod     = "31303130EBEAA59AD34759AF5EF05F3802A1C5FAEE8B459A22CA0B65C7B2797877960FD16CC6"
	AuthCookieValueHwssoUniportal = "yhZIXGAAYi_ahHM1671Sx7XhNS8eiJ5kuUGADjpkMzBhl1x1ko8FUzP4DBle4n_bMD8f9qYrJ6I91iJYJ0fGok1KUnHG9qZQMwFz0lLrybZaIaES0T3nm_agQ7WFA575UonHzfnGxcUKrIENfMXX3Dt7ykSl2Sx_btWQGnFdXA7Fef8fEv6TfFsh6YGMBuw7s9arDNYdOm7DibfDe1fvnUEXLdDnhCqQ8c3aSQkd3DPuHwYHj06KDVNnkwgpDJqg9CONGyVBlLIRL_b0UU3sw_bChrgT1Qizgafw_aDiNht7ElBxBxR_bVeIHvxZe_a4_bsdyf37nnCKw9GYCAZyHRyDgZkC9_acg_c_c####kOV97xUKRgcBen718R_bPiY12z6bVkstL5wKz709qvvMW5kHr0tkeVk2YEuLvHP1O_aqBGUynW73KqdIHiBgfRrZuxOyMEup1jUltV7uboIjZS47_bA9FOvl5bvu2GMjhsRLPW_…D08eEkVUThT7L_a9eTy_ajXw6kL25GOhJGkbYxfTbpfqW4nnuCVrg68DYx8efcSO5X04MeQHIj63bLyNUK3rPWFbT_b0gL5I33RYavdkY7Llpw_c_c####LAGQBCx08jF_bzNFZlCQHWNyGxzE2gzfAYFDuifHq8LhjlFI120xPzWhn5kJ_brIGGcSh13uleEDr5w9UJAQMoU69YCCTuEADamX1qAjp5nmkk5A2KILhHa8pb3S_aLczq6PYCNN5la5iqDtYZossiOZylrFFiFMTP5XY0kVPc8HnuZbrjvSHKZZY0MtY0uQo0FixMQ2UIvkWmQQTz_bgHVx_bhZheDxqqi6523SyNQ56bZxBI4EzJSdaSX2qhd93_bIhUll9aDb05WuP_aQZpdv70sbW0w0CecyXp_bLuWovcrxVgwj8VX7SsXk0qwstROo3rxNY45TCa_bTGITEZ2orLORJtA_c_c"
	AuthCookieValueHwssotinter    = "D7-AA-90-56-3B-3B-12-87-30-7C-5D-5B-2B-A7-15-A3"
	AuthCookieValueHwssotinter3   = "27527623713077"
	AuthCookieValueLogFlag        = "in"
	AuthCookieValueSid            = "313031300CFF5C45CE8C37684F24C70A85B2658E733FE2D73F469A86B05CE96C2AFF637839F29AC8E194902AD8DF09C5AF11E44FCA28C1CE51F9B6E307FB00B86CB72BE37D2875D7"
	AuthCookieValueSip            = "313031306C27122427C274AAAAE2B73BE2E441972964B386396F5233D7424F791E641CF9C92B56C0DB55C463A82F125B01B5EF83DA6506D92EAFDCEE7833AB39274E87A9"
	AuthCookieValueUid            = "31303130F5553906EFFA0D3E5B2370C0A75080E2C3F7B139783FBFDEA06B2941A794BE623E1C986E31AABE17D628BB2ED7AE1755"

	AuthCookieDomain = ".myones.net"
)

func UniportalLogin(ctx iris.Context) {
	email := ctx.FormValue("email")
	log.Printf("login email: %s", email)
	password := ctx.FormValue("password")
	log.Printf("login passoword: %s", password)

	ctx.Header("Set-Cookie", authCookie(AuthCookieAuthMethod, AuthCookieValueAuthMethod))
	ctx.Header("Set-Cookie", authCookie(AuthCookieHwssoUniportal, AuthCookieValueHwssoUniportal))
	ctx.Header("Set-Cookie", authCookie(AuthCookieHwssotinter, AuthCookieValueHwssotinter))
	ctx.Header("Set-Cookie", authCookie(AuthCookieHwssotinter3, AuthCookieValueHwssotinter3))
	ctx.Header("Set-Cookie", authCookie(AuthCookieLogFlag, AuthCookieValueLogFlag))
	ctx.Header("Set-Cookie", authCookie(AuthCookieSid, AuthCookieValueSid))
	ctx.Header("Set-Cookie", authCookie(AuthCookieSip, AuthCookieValueSip))
	ctx.Header("Set-Cookie", authCookie(AuthCookieUid, AuthCookieValueUid))

	redirect := ctx.FormValue("redirect")
	if redirect == "" {
		redirect = "/home"
	}
	ctx.ViewData("redirect", redirect)
	ctx.View("login-response.html")
}

func Logout(ctx iris.Context) {
	removeCookieWithDomain(ctx, AuthCookieAuthMethod, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieHwssoUniportal, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieHwssotinter, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieHwssotinter3, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieLogFlag, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieSid, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieSip, AuthCookieDomain)
	removeCookieWithDomain(ctx, AuthCookieUid, AuthCookieDomain)
	ctx.Redirect("/uniportal")
}

func removeCookieWithDomain(ctx iris.Context, name, domain string) {
	cookie := &http.Cookie{
		Name:     name,
		Domain:   domain,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Duration(1) * time.Minute),
	}
	ctx.SetCookie(cookie)
}

// 解决 ctx.SetCookieKV() / ctx.SetCookie() 无法设置跨域 cookie 的问题
func authCookie(name, value string) string {
	expires := time.Now().Add(time.Duration(2) * time.Hour)
	return fmt.Sprintf("%s=%s; Path=/; Domain=%s; Expires=%s; HttpOnly",
		name, value, AuthCookieDomain, expires.UTC().Format(http.TimeFormat))
}
