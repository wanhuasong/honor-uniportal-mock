package controllers

import (
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
)

const (
	iamAccessTokenKey = "honor_iam_access_token"
	iamStatusOK       = http.StatusCreated

	IamTokenCodeCreated = "201"
	IamTokenReqType     = "JWT-Token"
	IamTokenReqMethod   = "CREATE"
)

type IamTokenResponse struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type IamTokenRequestData struct {
	Type       string                    `json:"type"`
	Attributes IamTokenRequestAttributes `json:"attributes"`
}

type IamTokenRequestAttributes struct {
	Method     string `json:"method"`
	Account    string `json:"account"`
	Secret     string `json:"secret"`
	Project    string `json:"project"`
	Enterprise string `json:"enterprise"`
}

type IamTokenRequest struct {
	Data IamTokenRequestData `json:"data"`
}

func GenerateIamToken(ctx iris.Context) {
	var req IamTokenRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}
	log.Printf("IAM token req: %+v", req)

	resp := IamTokenResponse{
		Code:        IamTokenCodeCreated,
		Message:     "OK",
		AccessToken: "mockaccesstoken",
		ExpiresIn:   86400,
	}
	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(&resp)
}
