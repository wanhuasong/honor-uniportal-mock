package controllers

import (
	"github.com/wanhuasong/uniportal/db"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
)

const (
	IDaasAuthHeaderTenantID = "TenantID"
	IDaasAuthHeaderJWT      = "SSO-JWT-Authorization"

	DefaultTenantID = "1"

	DefaultIDassAuthURL   = "http://example.com"
	UserScopeInternetUser = "INTERNET_USER"

	idaasStatusOK = http.StatusOK
)

type AuthToken struct {
	AuthMethod     string `json:"authmethod"`
	HwssoUniportal string `json:"hwsso_uniportal"`
	Hwssotinter    string `json:"hwssotinter"`
	Hwssotinter3   string `json:"hwssotinter3"`
	LogFlag        string `json:"logFlag"`
	Sid            string `json:"sid"`
	Sip            string `json:"sip"`
	Uid            string `json:"uid"`
}

type IDaasAuthRequest struct {
	URL       string    `json:"url"`
	UserScope string    `json:"userScope"`
	Token     AuthToken `json:"token"`
}

type IDaasAuthResponseProperties struct {
	EmployeeNumber string `json:"employeenumber"`
	Mail           string `json:"mail"`
}

type IDaasAuthResponse struct {
	AllProperties IDaasAuthResponseProperties `json:"allProperties"`
}

func IdaasAuth(ctx iris.Context) {
	var req IDaasAuthRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}
	log.Printf("IDaas req: %+v", req)

	number := "12345678"
	if db.DefaultStore != nil {
		number = db.DefaultStore.Get(StoreNumber)
	}
	resp := IDaasAuthResponse{
		AllProperties: IDaasAuthResponseProperties{
			Mail:           req.Token.Sid,
			EmployeeNumber: number,
		},
	}
	log.Printf("IDaas resp: %+v", resp)
	ctx.JSON(&resp)
}
