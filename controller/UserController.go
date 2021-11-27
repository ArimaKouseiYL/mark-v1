package controller

import (
	"github.com/gin-gonic/gin"
	"mark-v1/common/resp"
	"mark-v1/models"
	"mark-v1/models/Vo"
	"mark-v1/pkg/snowflake"
	"mark-v1/service"
	"strconv"
)

// Registry 用户注册
// @Summary 用户注册接口
// @Description 用户注册
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body Vo.UserRegistryVo true "用户Vo"
// @Success 200 {object}
// @Router /users/registry [post]
func Registry(c *gin.Context) {

	userRegistryVo := new(Vo.UserRegistryVo)
	if err := c.ShouldBind(userRegistryVo); err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeInvalidParam, err.Error())
		return
	}
	id := snowflake.GetSnowFlakeId()
	user := &models.Users{
		Id:       id,
		UserName: userRegistryVo.UserName,
		Password: userRegistryVo.Password,
		Address:  userRegistryVo.Address,
		Phone:    userRegistryVo.Phone,
		Age:      userRegistryVo.Age,
	}
	err := registry(user)
	if err != nil {
		resp.ResponseError(c, resp.CodeInvalidParam)
		return
	}
	resp.ResponseSuccess(c, nil)
}

func registry(user *models.Users) error {
	userService := new(service.UserService)
	return userService.Registry(user)
}

func getUserByName(username string) models.Users {
	userService := new(service.UserService)
	return userService.GetUserByName(username)
}

func getUserById(id int) models.Users {
	userService := new(service.UserService)
	return userService.GetUserById(id)
}

// GetUserByNameHandle 根据用户名称，获取用户详情接口
// @Summary 用户详情接口
// @Description 根据用户名称，获取用户详情接口
// @Tags 用户相关接口
// @Accept mpfd
// @Produce octet-stream
// @Param username path string true "用户姓名"
// @Success 200 {object} models.Users
// @Router /users/{username} [get]
func GetUserByNameHandle(context *gin.Context) {
	username := context.Param("username")
	name := getUserByName(username)
	context.JSON(200, name)
}

// GetUserByIdHandle 根据用户id，获取用户详情接口
// @Summary 用户详情接口
// @Description 根据用户id，获取用户详情接口
// @Tags 用户相关接口
// @Accept mpfd
// @Produce octet-stream
// @Param userId path int64 true "用户id"
// @Success 200 {object} models.Users
// @Router /users/userId/{userId} [get]
func GetUserByIdHandle(context *gin.Context) {
	idStr := context.Param("userId")
	id, _ := strconv.Atoi(idStr)
	name := getUserById(id)
	context.JSON(200, name)
}
