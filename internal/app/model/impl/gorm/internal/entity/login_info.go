package entity

import (
	"context"
	"github.com/jinzhu/gorm"
	"sooty-tern/internal/app/schema"
)

func GetLoginInfoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, LoginInfo{})
}

type SchemaLoginInfo schema.LoginInfo

type LoginInfo struct {
	Model
	UserId     int    `gorm:"column:user_id"`
	SourceType int    `gorm:"column:source_type"` //登录类型 1：微信小程序
	Uid        string `gorm:"column:uid"`         //1.微信小程序（openId） 2:手机号
	UidKey     string `gorm:"column:uid_key"`     //1.session_key 2.password
	Salt       string `gorm:"column:salt"`        //盐
}

func (l LoginInfo) ToSchemaLoginInfo() *schema.LoginInfo {
	return &schema.LoginInfo{
		UserId:     l.UserId,
		SourceType: l.SourceType,
		Uid:        l.Uid,
		UidKey:     l.UidKey,
		Salt:       l.Salt,
	}
}

func (s SchemaLoginInfo) ToLoginInfoEntity() *LoginInfo {
	return &LoginInfo{
		UserId:     s.UserId,
		SourceType: s.SourceType,
		Uid:        s.Uid,
		UidKey:     s.UidKey,
		Salt:       s.Salt,
	}
}

// TableName 表名
func (l LoginInfo) TableName() string {
	return l.Model.TableName("login_info")
}
