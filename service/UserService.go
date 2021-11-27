package service

import (
	"mark-v1/common/db"
	"mark-v1/models"
)

type UserService struct {
}

func (UserService UserService) Registry(user *models.Users) error {
	sqlStr := "insert into users(id,user_name,password,age,address,phone) values (?,?,?,?,?,?)"
	_, err := db.Db.Exec(sqlStr, user.Id, user.UserName, user.Password, user.Age, user.Address, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (UserService UserService) GetUserByName(username string) models.Users {

	var user models.Users

	sqlStr := "select user_name,password,age,id,address,phone from users where user_name=?"

	error := db.Db.QueryRow(sqlStr, username).Scan(&user.UserName, &user.Password, &user.Age, &user.Id, &user.Address, &user.Phone)

	if error != nil {

	}

	return user
}

func (UserService UserService) GetUserById(id int) models.Users {

	var user models.Users

	sqlStr := "select user_name,password,age,id,address,phone from users where id=?"

	error := db.Db.QueryRow(sqlStr, id).Scan(&user.UserName, &user.Password, &user.Age, &user.Id, &user.Address, &user.Phone)

	if error != nil {

	}

	return user
}
