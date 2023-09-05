package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/config"
)

func main() {
	r := gin.New()
	r.ContextWithFallback = true

	r.Run(fmt.Sprintf(":%s", config.EnvVar.ServerPort))
}
