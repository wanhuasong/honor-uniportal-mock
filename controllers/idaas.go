package controllers

import "github.com/kataras/iris/v12"

const UserScopeInternetUser = "INTERNET_USER"

type IdaasAuthRequest struct {
	URL       string `json:"url"`
	UserScope string `json:"userScope"`
	Token     struct {
		AuthMethod     string `json:"authmethod"`
		HwssoUniportal string `json:"hwsso_uniportal"`
		Hwssotinter    string `json:"hwssotinter"`
		Hwssotinter3   string `json:"hwssotinter3"`
		LogFlag        string `json:"logFlag"`
		Sid            string `json:"sid"`
		Sip            string `json:"sip"`
		Uid            string `json:"uid"`
	} `json:"token"`
}

type IdaasAuthResponse struct {
	AllProperties struct {
		EmployeeNumber string `json:"employeenumber"`
		Mail           string `json:"mail"`
	} `json:"allProperties"`
}

func IdaasAuth(ctx iris.Context) {}
