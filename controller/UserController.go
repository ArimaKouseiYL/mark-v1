package controller

import (
	"github.com/gin-gonic/gin"
	"mark-v1/models"
	"mark-v1/models/Vo"
	"mark-v1/service"
	"strconv"
)

func Registry(user Vo.UserRegistryVo) {

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
