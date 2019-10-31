package internal

import (
	"context"
	"fmt"
	"github.com/medivhzhan/weapp"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/model"
	"sooty-tern/internal/app/schema"
	"sooty-tern/pkg/auth"
)

type LoginInfo struct {
	auth           auth.Auth
	LoginInfoModel model.ILoginInfoModel
}

func NewLoginInfo(loginInfoModel model.ILoginInfoModel, auth auth.Auth) *LoginInfo {
	return &LoginInfo{
		auth:           auth,
		LoginInfoModel: loginInfoModel,
	}
}

// auth
func (l *LoginInfo) Login(ctx context.Context, code string) (*schema.LoginRes, error) {
	//1.code -> openId 、sessionKey
/*	result, err := mini_prog.Code2Session(code)
	if err != nil {
		return nil, errors.WithStack(err)
	}*/
	result := &weapp.LoginResponse{
		OpenID:     "abcdefff",
		SessionKey: "wqeswcds12123",
	}
	openId := result.OpenID
	sessionKey := result.SessionKey
	token, err := l.GenerateToken(openId, sessionKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	//2. openId 去查数据库
	params := schema.LoginInfoQueryParam{
		Uid: openId,
	}
	exist, err := l.LoginInfoModel.Get(ctx, params)
	var isRegister bool = false
	if exist != nil {
		isRegister = true
	}
	return &schema.LoginRes{
		LoginTokenInfo: token,
		IsRegister:     isRegister,
	}, nil
}

// GenerateToken
func (l *LoginInfo) GenerateToken(openId, sessionKey string) (*schema.LoginTokenInfo, error) {
	data := fmt.Sprintf("%s:%s", openId, sessionKey)
	tokenInfo, err := l.auth.GenerateToken(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	token := &schema.LoginTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpireAt:    tokenInfo.GetExpiresAt(),
	}
	return token, nil
}

// DestroyToken
func (l *LoginInfo) DestroyToken(tokenString string) error {
	err := l.auth.DestroyToken(tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
