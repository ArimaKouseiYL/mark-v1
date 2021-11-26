package dao

import "mark-v1/models"

type UserRepository interface {
	GetUserByName(name string) models.Users
}
