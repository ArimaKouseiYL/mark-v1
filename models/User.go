package models

type Users struct {
	Id       int64  `db:"id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	Address  string `db:"address"`
	Phone    string `db:"phone"`
	Age      int    `db:"age"`
}
