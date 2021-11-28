package main

import (
	"go.uber.org/zap"
	_ "mark-v1/cmd/docs"
	"mark-v1/common/db"
	"mark-v1/common/logger"
	"mark-v1/routers"
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

	r := routers.SetRouters()

	err := logger.Init()
	if err != nil {
		zap.L().Error("日志初始化失败", zap.Error(err))
	}

	r.Run(":8080")

}
