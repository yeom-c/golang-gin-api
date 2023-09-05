package main

import "github.com/yeom-c/golang-gin-api/config"

func main() {
	println(config.EnvVar.ServerPort)
}
