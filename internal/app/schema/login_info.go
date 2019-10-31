package schema

type LoginInfo struct {
	UserId     int    `json:"user_id"`
	SourceType int    `json:"source_type"` //登录类型 1：微信小程序
	Uid        string `json:"uid"`         //1.微信小程序（openId） 2:手机号
	UidKey     string `json:"uid_key"`     //1.session_key 2.password
	Salt       string `json:"salt"`        //盐
}

type LoginInfoQueryParam struct {
	Uid    string
	UidKey string
}
