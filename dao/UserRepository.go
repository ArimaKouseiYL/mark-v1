package dao

import "mark-v1/models"

type UserRepository interface {
	GetUserByName(name string) models.Users
	GetUserById(id int64) models.Users
}
