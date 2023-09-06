package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/config"
	"github.com/yeom-c/golang-gin-api/database"
	"github.com/yeom-c/golang-gin-api/route"
)

func main() {
	r := gin.New()
	r.ContextWithFallback = true

	defer func() {
		gormInstance, _ := database.DB.Gorm.DB()
		gormInstance.Close()
	}()

	// route 등록.
	route.Init(r)

	r.Run(fmt.Sprintf(":%s", config.EnvVar.ServerPort))
}
