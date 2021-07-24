package config

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg *Config

type Config struct {
	TokenAuthURL        string   `json:"honor_token_auth_url"`
	QueryUserURL        string   `json:"honor_query_user_url"`
	IamTokenURL         string   `json:"honor_iam_token_url"`
	EnableUniportal     bool     `json:"enable_honor_uniportal"`
	IdaasAccessKey      string   `json:"honor_idaas_access_key"`
	IdaasSecretKey      string   `json:"honor_idaas_secret_key"`
	UniportalLoginURL   string   `json:"honor_uniportal_login_url"`
	IamAccessKey        string   `json:"honor_iam_access_key"`
	IamSecretKey        string   `json:"honor_iam_secret_key"`
	IamProject          string   `json:"honor_iam_project"`
	IamEnterprise       string   `json:"honor_iam_enterprise"`
	InternalEmailSuffix []string `json:"honor_internal_email_suffix"`
}

func InitConfig(cfgFile string) error {
	b, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return err
	}
	Cfg = new(Config)
	return json.Unmarshal(b, &Cfg)
}
