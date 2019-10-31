package errors

import (
	"github.com/pkg/errors"
)

// 定义错误函数的别名
var (
	New       = errors.New
	Wrap      = errors.Wrap
	Wrapf     = errors.Wrapf
	WithStack = errors.WithStack
)

// 定义错误
var (
	// 公共错误
	ErrNotFound                = New("资源不存在")
	ErrMethodNotAllow          = New("方法不被允许")
	ErrBadRequest              = New("请求发生错误")
	ErrInvalidRequestParameter = New("无效的请求参数")
	ErrTooManyRequests         = New("请求过于频繁")
	ErrUnknownQuery            = New("未知的查询类型")
	ErrInvalidParent           = New("无效的父级节点")
	ErrNotAllowDeleteWithChild = New("含有子级，不能删除")
	ErrResourceExists          = New("资源已经存在")
	ErrResourceNotAllowDelete  = New("资源不允许删除")

	// 权限错误
	ErrNoPerm         = New("无访问权限")
	ErrNoResourcePerm = New("无资源的访问权限")

	// 用户错误
	ErrInvalidUserName = New("无效的用户名")
	ErrInvalidPassword = New("无效的密码")
	ErrInvalidUser     = New("无效的用户")
	ErrUserDisable     = New("用户被禁用")
	ErrUserNotEmptyPwd = New("密码不允许为空")

	// login
	ErrRegister               = New("注册失败")
	ErrLoginNotAllowModifyPwd = New("不允许修改密码")
	ErrLoginInvalidOldPwd     = New("旧密码不正确")
	ErrLoginInvalidVerifyCode = New("无效的验证码")
)

func init() {
	// 公共错误
	//400
	newBadRequestError(ErrBadRequest)
	newBadRequestError(ErrInvalidRequestParameter)
	newBadRequestError(ErrUnknownQuery)
	newBadRequestError(ErrInvalidParent)
	newBadRequestError(ErrNotAllowDeleteWithChild)
	newBadRequestError(ErrResourceExists)
	newBadRequestError(ErrResourceNotAllowDelete)

	//401
	newUnauthorizedError(ErrNoPerm)
	newUnauthorizedError(ErrNoResourcePerm)

	//404
	newNotFoundError(ErrNotFound)

	//405
	newMethodNotAllowError(ErrMethodNotAllow)

	//429
	newTooManyRequestsError(ErrTooManyRequests)

	// 业务
	newBadRequestError(ErrInvalidUserName)
	newBadRequestError(ErrInvalidPassword)
	newBadRequestError(ErrInvalidUser)
	newBadRequestError(ErrUserDisable)
	newBadRequestError(ErrUserNotEmptyPwd)
	newBadRequestError(ErrRegister)
	newBadRequestError(ErrLoginNotAllowModifyPwd)
	newBadRequestError(ErrLoginInvalidOldPwd)
	newBadRequestError(ErrLoginInvalidVerifyCode)
}
