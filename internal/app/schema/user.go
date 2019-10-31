package schema

import "time"

// Users 用户对象列表
type Users []*User

// User 用户对象
type User struct {
	RecordId string    `json:"record_id"` //绿洲号 uniq
	Username string    `json:"username" `
	Avatar   string    `json:"avatar"`
	Info     string    `json:"info"`
	Sex      int       `json:"sex"`
	Email    string    `json:"email" `
	Birthday *time.Time `json:"birthday,omitempty"`
	City     string    `json:"city"`
	Status   int       `json:"status" `
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	UserName     string // 用户名
	LikeUserName string // 用户名(模糊查询)
	RecordId     string // 绿洲号
	LikeRecordId string //绿洲号(模糊查询)
	Status       int
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	PageParam *PageParam // 分页参数
}

// UserQueryResult 查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PageInfo
}
