package db

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	models "mark-v1/models"
)

var Db *sql.DB

func GetDBConfig() *models.DbConfig {

	viper.SetConfigFile("./config/application.yaml")

	error := viper.ReadInConfig()
	if error != nil {
		zap.L().Error("读取配置文件失败")
	}

	database := viper.GetString("db.database")
	fmt.Println(database)

	var db = new(models.DbConfig)
	viper.Sub("db").Unmarshal(db)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		viper.Sub("db").Unmarshal(db)
		zap.L().Info("配置文件更改:")
	})

	return db
}

func MysqlInit() {

	dbConfig := GetDBConfig()

	var error error
	Db, error = sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@tcp("+dbConfig.Url+":"+dbConfig.Port+")/"+dbConfig.Database)

	if error != nil {
		zap.L().Error("数据库连接失败", zap.Error(error))
	}

	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

}
