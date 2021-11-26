package main

import (
	"github.com/gin-gonic/gin"
	"mark-v1/common/db"
	"mark-v1/controller"
)

func main() {

	db.MysqlInit()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:username", func(context *gin.Context) {
			username := context.Param("username")
			name := controller.GetUserByName(username)
			context.JSONP(200, name)
		})
	}

	r.Run(":8080")
}
