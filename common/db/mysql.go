package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"mark-v1/models"
)

var Db *sql.DB

func GetDBConfig() *models.DbConfig {

	viper.SetConfigFile("./config/application.yaml")

	error := viper.ReadInConfig()
	if error != nil {
		log.Fatalln("读取配置文件失败")
	}

	database := viper.GetString("db.database")
	fmt.Println(database)

	db := new(models.DbConfig)

	viper.Sub("db").Unmarshal(db)

	return db
}

func MysqlInit() {

	dbConfig := GetDBConfig()

	var error error
	Db, error = sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@tcp("+dbConfig.Url+":"+dbConfig.Port+")/"+dbConfig.Database)

	if error != nil {
		log.Fatalln("数据库连接失败")
	}

	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

}
