package api

import (
	"github.com/gin-gonic/gin"
	"sooty-tern/internal/app/ginplus"
	"sooty-tern/internal/app/schema"
	"sooty-tern/internal/app/service"
	"sooty-tern/pkg/util"
)

// User
type User struct {
	UserService service.IUser
}

func NewUser(userService service.IUser) *User {
	return &User{
		UserService: userService,
	}
}

// @Router /api/v1/users [get]
func (u *User) Query(c *gin.Context) {
	var params schema.UserQueryParam
	params.LikeUserName = c.Query("username")
	params.LikeRecordId = c.Query("recordId")
	if v := util.S(c.Query("status")).DefaultInt(0); v > 0 {
		params.Status = v
	}
	result, err := u.UserService.Query(ginplus.NewContext(c), params, schema.UserQueryOptions{
		PageParam: ginplus.GetPageParam(c),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResList(c, result.Data, result.PageResult)
}

// @Router /api/v1/users/{id} [get]
func (u *User) Get(c *gin.Context) {
	item, err := u.UserService.Get(ginplus.NewContext(c), c.Param("record_id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResDetail(c, item)
}

// @Router /api/v1/users [post]
func (u *User) Create(c *gin.Context) {
	var item schema.User
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	newItem, err := u.UserService.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResDetail(c, newItem)
}

// @Router /api/v1/users/{id}/enable [patch]
func (u *User) Enable(c *gin.Context) {
	err := u.UserService.UpdateStatus(ginplus.NewContext(c), c.Param("record_id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// @Router /api/v1/users/{id}/disable [patch]
func (u *User) Disable(c *gin.Context) {
	err := u.UserService.UpdateStatus(ginplus.NewContext(c), c.Param("record_id"), 2)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
