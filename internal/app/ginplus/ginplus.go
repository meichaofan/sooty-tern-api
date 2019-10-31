package ginplus

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	icontext "sooty-tern/internal/app/context"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/schema"
	"sooty-tern/pkg/logger"
	"sooty-tern/pkg/util"
)

// 定义上下文中的键
const (
	prefix     = "sooty_tern"
	UserIDKey  = prefix + "/user_id"  // UserIDKey 存储上下文中的键(用户ID)
	TraceIDKey = prefix + "/trace_id" // TraceIDKey 存储上下文中的键(跟踪ID)
	ResBodyKey = prefix + "/res_body" // ResBodyKey 存储上下文中的键(响应Body数据)
)

// NewContext 封装上线文入口
func NewContext(c *gin.Context) context.Context {
	parent := context.Background()

	if v := GetTraceID(c); v != "" {
		parent = icontext.NewTraceID(parent, v)
		parent = logger.NewTraceIDContext(parent, GetTraceID(c))
	}

	if v := GetUserID(c); v != "" {
		parent = icontext.NewUserID(parent, v)
		parent = logger.NewUserIDContext(parent, v)
	}

	return parent
}

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetPageNumber 获取分页的页索引
func GetPageNumber(c *gin.Context) int {
	defaultVal := 1
	if v := c.Query("pageNumber"); v != "" {
		if iv := util.S(v).DefaultInt(defaultVal); iv > 0 {
			return iv
		}
	}
	return defaultVal
}

// GetPageSize 获取分页的页大小(最大50)
func GetPageSize(c *gin.Context) int {
	defaultVal := 10
	if v := c.Query("pageSize"); v != "" {
		if iv := util.S(v).DefaultInt(defaultVal); iv > 0 {
			if iv > 50 {
				iv = 50
			}
			return iv
		}
	}
	return defaultVal
}

// GetPageParam 获取分页查询参数
func GetPageParam(c *gin.Context) *schema.PageParam {
	return &schema.PageParam{
		PageNumber: GetPageNumber(c),
		PageSize:   GetPageSize(c),
	}
}

// GetTraceID 获取追踪ID
func GetTraceID(c *gin.Context) string {
	return c.GetString(TraceIDKey)
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		logger.Warnf(NewContext(c), err.Error())
		return errors.ErrInvalidRequestParameter
	}
	return nil
}

// ResList	响应分页数据
func ResList(c *gin.Context, v interface{}, pi *schema.PageInfo) {
	r := &schema.HTTPSucRes{
		Result: &schema.HTTPList{
			Data:       v,
			PageSize:   pi.PageSize,
			PageNumber: pi.PageNumber,
			Total:      pi.Total,
		},
	}
	resSuccess(c, r)
}

// ResDetail 响应详情
func ResDetail(c *gin.Context, v interface{}) {
	r := &schema.HTTPSucRes{
		Result: &schema.HTTPDetail{
			Data: v,
		},
	}
	resSuccess(c, r)
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	r := &schema.HTTPSucRes{
		Result: map[string]bool{"success": true},
	}
	resSuccess(c, r)
}

// ResSuccess 响应成功
func resSuccess(c *gin.Context, v *schema.HTTPSucRes) {
	v.RequestId = GetTraceID(c)
	resJSON(c, http.StatusOK, v)
}

// ResError 响应错误
func ResError(c *gin.Context, err error) {
	errItem := &schema.ErrorItem{
		Code:    500,
		Status:  "server internal error",
		Message: "服务器发生错误",
	}

	if err, ok := errors.FromErrorCode(err); ok {
		errItem = err
	}

	if errItem.Code == 500 && err != nil {
		span := logger.StartSpan(NewContext(c))
		span = span.WithField("stack", fmt.Sprintf("%+v", err))
		span.Errorf(err.Error())
	}

	resJSON(c, errItem.Code, schema.HTTPErrRes{Error: errItem, RequestId: GetTraceID(c)})
}

// ResJSON 响应JSON数据
func resJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}
