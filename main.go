package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/config"
	"github.com/yeom-c/golang-gin-api/route"
)

func main() {
	r := gin.New()
	r.ContextWithFallback = true

	// route 등록.
	route.Init(r)

	r.Run(fmt.Sprintf(":%s", config.EnvVar.ServerPort))
}
