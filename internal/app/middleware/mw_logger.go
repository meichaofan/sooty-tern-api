package middleware

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"sooty-tern/internal/app/ginplus"
	"sooty-tern/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		method := c.Request.Method
		span := logger.StartSpan(ginplus.NewContext(c), logger.SetSpanTitle("access-log"), logger.SetSpanFuncName(JoinRouter(method, p)))
		start := time.Now()

		fields := make(map[string]interface{})
		fields["path"] = p
		fields["remote_addr"] = c.ClientIP()
		fields["method"] = method
		fields["url"] = c.Request.URL.String()
		fields["protocol"] = c.Request.Proto
		fields["header"] = c.Request.Header
		fields["user_agent"] = c.GetHeader("User-Agent")

		// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType == "application/json" {
				body, err := ioutil.ReadAll(c.Request.Body)
				c.Request.Body.Close()
				if err == nil {
					buf := bytes.NewBuffer(body)
					c.Request.Body = ioutil.NopCloser(buf)
					fields["content_length"] = c.Request.ContentLength
					fields["body"] = string(body)
				}
			}
		}
		c.Next()

		cost := time.Since(start).Nanoseconds() / 1e6
		fields["status_code"] = c.Writer.Status()
		fields["res_length"] = c.Writer.Size()
		fields["cost"] = cost
		if v, ok := c.Get(ginplus.ResBodyKey); ok {
			if b, ok := v.([]byte); ok {
				fields["response"] = string(b)
			}
		}
		fields[logger.UserIDKey] = ginplus.GetUserID(c)
		span.WithFields(fields).Infof("[request_id=%s]",
			ginplus.GetTraceID(c))
	}
}
