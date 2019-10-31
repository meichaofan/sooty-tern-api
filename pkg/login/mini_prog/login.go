package mini_prog

import (
	"github.com/medivhzhan/weapp"
	"sooty-tern/internal/app/config"
)

func getAuthCode() (string, string) {
	cfg := config.GetGlobalConfig()
	appId := cfg.Wechat.AppId
	appSecret := cfg.Wechat.AppSecret
	return appId, appSecret
}

/**
通过code，获取openId , sessionKey
*/
func Code2Session(code string) (*weapp.LoginResponse, error) {
	appId, appSecret := getAuthCode()
	res, err := weapp.Login(appId, appSecret, code)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
