package database

import (
	"fmt"
	"log"

	"github.com/yeom-c/golang-gin-api/config"
	"github.com/yeom-c/golang-gin-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.EnvVar.DBUser, config.EnvVar.DBPassword, config.EnvVar.DBHost, config.EnvVar.DBPort, config.EnvVar.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// auto migration.
	db.AutoMigrate(&model.Account{})

	Conn = db
}
