package entity

import (
	"context"
	"github.com/jinzhu/gorm"
	"sooty-tern/internal/app/schema"
	"time"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, User{})
}

// SchemaUser 用户对象
type SchemaUser schema.User

// ToUser转换为用户实体
func (s SchemaUser) ToUserEntity() *User {
	entity := &User{
		RecordId: s.RecordId,
		Username: s.Username,
		Avatar:   s.Avatar,
		Info:     s.Info,
		Email:    s.Email,
		Sex:      s.Sex,
		Birthday: s.Birthday,
		City:     s.City,
		Status:   s.Status,
	}
	return entity
}

type User struct {
	Model
	RecordId string     `gorm:"column:record_id;size:16"`         // 绿洲ID
	Username string     `gorm:"column:username;size:64;not null"` // 用户名
	Avatar   string     `gorm:"column:avatar;size:128;"`          //用户头像
	Info     string     `gorm:"column:info;size:128;"`            //简介
	Email    string     `gorm:"column:email;size:32;"`            // 邮箱
	Sex      int        `gorm:"column:sex;size:4;"`               // 性别 1：男 2：女 3：未知
	Birthday *time.Time `gorm:"column:birthday;"`                 // 生日
	City     string     `gorm:"column:city;"`                     //城市
	Status   int        `gorm:"column:status;"`                   // 状态 1:启用 2:停用
	IsDel    int        `gorm:"column:is_del;size:4"`             // 是否删除 0:未删除 1:已删除
}

// TableName 表名
func (u User) TableName() string {
	return u.Model.TableName("user")
}

func (u User) String() string {
	return toString(u)
}

// 用户实体转化为schema对象
func (u User) ToSchemaUser() *schema.User {
	obj := &schema.User{
		RecordId: u.RecordId,
		Username: u.Username,
		Avatar:   u.Avatar,
		Info:     u.Info,
		Email:    u.Email,
		Sex:      u.Sex,
		Birthday: u.Birthday,
		City:     u.City,
		Status:   u.Status,
	}
	return obj
}

type Users []*User

func (u Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(u))
	for i, item := range u {
		list[i] = item.ToSchemaUser()
	}
	return list
}
