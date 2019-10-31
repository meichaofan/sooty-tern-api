package schema

// 异常返回定义
type HTTPErrRes struct {
	Error     *ErrorItem `json:"error"`
	RequestId string     `json:"request_id"`
}

type ErrorItem struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// 正常返回定义
type HTTPSucRes struct {
	Result    interface{} `json:"result"`
	RequestId string      `json:"request_id"`
}

type HTTPDetail struct {
	Data interface{} `json:"data"`
}

type HTTPList struct {
	Data       interface{} `json:"data"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	Total      int         `json:"total"`
}

//分页查询信息
type PageParam struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

//分页结果信息
type PageInfo struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
}
