package main

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "mark-v1/cmd/docs"
	"mark-v1/common/db"
	"mark-v1/controller"
)

// @title mark-v1 api文档
// @version 1.0
// @description 一个简单的go web入门项目
// @termsOfService http://swagger.io/terms/

// @contact.name mark
// @contact.email leis17@163.com

// @host localhost:8080
// @BasePath /api/v1
func main() {

	db.MysqlInit()

	r := gin.Default()

	//url := gs.URL("http://localhost:8080/swagger/doc.json")

	// 执行swagger ,  swag init --parseDependency --parseInternal
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:username", controller.GetUserByNameHandle)
		v1.GET("/users/userId/:userId", controller.GetUserByIdHandle)
		v1.POST("/users/registry", controller.Registry)
	}

	r.Run(":8080")
}
