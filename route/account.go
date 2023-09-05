package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/controller"
)

func AccountRG(rg *gin.RouterGroup) {
	accountController := controller.NewAccountController()
	accountRG := rg.Group("/account")

	accountRG.GET("", accountController.Get)
}
