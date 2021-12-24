package controllers

import (
	"github.com/wanhuasong/uniportal/db"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
)

const (
	mdmStatusOK = http.StatusOK

	mdmResponseStatusOK = "200"

	errorMessageTokenExpired = "com.auth0.jwt.exceptions.TokenExpiredException"

	subStatusTokenExpired = "40140117"

	personTypeWorking    = "1"   // 在职
	personTypeTerminated = "2"   // 离职
	personTypeAll        = "all" // 全部

	mdmAuthHeader = "Authorization"
)

type UserDetail struct {
	FullName string `json:"full_name"`
}

type UserDetailResponse struct {
	Result    []UserDetail `json:"result"`
	Count     int          `json:"count"`
	Status    string       `json:"status"`
	Message   string       `json:"message"`
	ActiveID  string       `json:"activeId"`
	SubStatus string       `json:"subStatus"`
}

type UserDetailRequestParams struct {
	EmployeeNumber   string `json:"employee_number"`
	SystemPersonType string `json:"system_person_type"`
}

type UserDetailRequest struct {
	PageSize      int                     `json:"pageSize"`
	PageIndex     int                     `json:"pageIndex"`
	UseDataFormat bool                    `json:"useDataFormat"`
	Params        UserDetailRequestParams `json:"params"`
}

func GetUserDetail(ctx iris.Context) {
	var req UserDetailRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}
	log.Printf("MDM req: %+v", req)

	token := ctx.GetHeader(mdmAuthHeader)
	log.Printf("MDM authorization token: %s", token)

	name := "张三"
	if db.DefaultStore != nil {
		name = db.DefaultStore.Get(StoreName)
	}
	resp := UserDetailResponse{
		Count:   1,
		Status:  mdmResponseStatusOK,
		Message: "ok",
		Result: []UserDetail{
			{
				FullName: name,
			},
		},
	}
	ctx.JSON(&resp)
}
