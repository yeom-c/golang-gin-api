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

	db := database.Conn
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	// route 등록.
	route.Init(r, db)

	r.Run(fmt.Sprintf(":%s", config.EnvVar.ServerPort))
}
