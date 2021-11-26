package Vo

type UserRegistryVo struct {
	Id         int64  `json:"userId"`
	UserName   string `json:"userName" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"password" validate:"required, eqfield=Password"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Age        int    `json:"age" validate:"max=120,min=0"`
}
