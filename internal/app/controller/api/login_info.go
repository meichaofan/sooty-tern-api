package api

import (
	"github.com/gin-gonic/gin"
	"sooty-tern/internal/app/ginplus"
	"sooty-tern/internal/app/schema"
	"sooty-tern/internal/app/service"
)

type LoginInfo struct {
	LoginInfoService service.ILoginInfo
}

func NewLoginInfo(loginInfoService service.ILoginInfo) *LoginInfo {
	return &LoginInfo{
		LoginInfoService: loginInfoService,
	}
}

func (l *LoginInfo) Login(c *gin.Context) {
	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	info, err := l.LoginInfoService.Login(ginplus.NewContext(c), item.Code)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResDetail(c, info)
}