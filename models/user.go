package models

type Users struct {
	Id       int64  `db:"id"`        // 用户id
	UserName string `db:"user_name"` // 用户名称
	Password string `db:"password"`  // 用户密码
	Address  string `db:"address"`   // 地址
	Phone    string `db:"phone"`     // 手机号码
	Age      int    `db:"age"`       // 年龄
}
