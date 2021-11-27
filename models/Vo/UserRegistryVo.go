package Vo

type UserRegistryVo struct {
	Id         int64  `json:"userId"`                                         // id
	UserName   string `json:"userName" binding:"required"`                    // 名称
	Password   string `json:"password" binding:"required"`                    // 密码
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"` // 确认密码
	Address    string `json:"address"`                                        // 地址
	Phone      string `json:"phone"`                                          // 手机号码
	Age        int    `json:"age" binding:"max=120,min=0"`                    // 年龄
}

type UserLoginVo struct {
	UserName string `json:"username" binding:"required"` // 名称
	Password string `json:"password" binding:"required"` // 密码
	Token    string `json:"token"`                       // token
}
