package schema

// LoginParam 登录参数
type LoginParam struct {
	Code string `json:"code"`
}

// LoginTokenInfo
type LoginTokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpireAt    int64  `json:"expire_at"`
}

type LoginRes struct {
	LoginTokenInfo *LoginTokenInfo `json:"login_token_info"`
	IsRegister     bool            `json:"is_register"`
}
