package errors

import (
	"net/http"
	"sooty-tern/internal/app/schema"
)

var (
	codes = make(map[error]*schema.ErrorItem)
)

// newErrorCode 设定错误码
func newErrorItem(err error, code int, message string, status string) error {
	errorItem := &schema.ErrorItem{
		Code:    code,
		Status:  status,
		Message: message,
	}
	codes[err] = errorItem
	return err
}

// FromErrorCode 获取错误码
func FromErrorCode(err error) (*schema.ErrorItem, bool) {
	v, ok := codes[err]
	return v, ok
}

// newBadRequestError 创建请求错误
func newBadRequestError(err error) {
	newErrorItem(err, 400, err.Error(), http.StatusText(400));
}

// newUnauthorizedError 创建未授权错误
func newUnauthorizedError(err error) {
	newErrorItem(err, 401, err.Error(), http.StatusText(401));
}

// newNotFoundError
func newNotFoundError(err error) {
	newErrorItem(err, 404, err.Error(), http.StatusText(404))
}

// newMethodNotAllowError
func newMethodNotAllowError(err error) {
	newErrorItem(err, 405, err.Error(), http.StatusText(404))
}

// newErrTooManyRequests
func newTooManyRequestsError(err error) {
	newErrorItem(err, 429, err.Error(), http.StatusText(429))
}

// newInternalServerError 创建服务器错误
func newInternalServerError(err error) {
	newErrorItem(err, 500, err.Error(), http.StatusText(500))
}