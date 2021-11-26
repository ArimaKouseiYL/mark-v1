package controller

import (
	"mark-v1/models"
	"mark-v1/models/Vo"
	"mark-v1/service"
)

func Registry(user Vo.UserRegistryVo) {

}

func GetUserByName(username string) models.Users {
	userService := new(service.UserService)
	return userService.GetUserByName(username)
}
