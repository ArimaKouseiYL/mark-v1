package routers

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"mark-v1/common/auth"
	"mark-v1/controller"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	//url := gs.URL("http://localhost:8080/swagger/doc.json")

	// 执行swagger ,  swag init --parseDependency --parseInternal
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/users/registry", controller.Registry)
	v1.POST("/users/login", controller.Login)

	v1.Use(auth.JWTAuthMiddleware())
	{
		v1.GET("/users/:username", controller.GetUserByNameHandle)
		v1.GET("/users/userId/:userId", controller.GetUserByIdHandle)
	}

	return r
}
