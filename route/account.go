package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/controller"
	"github.com/yeom-c/golang-gin-api/repository"
	"github.com/yeom-c/golang-gin-api/service"
	"gorm.io/gorm"
)

func AccountRG(rg *gin.RouterGroup, db *gorm.DB) {
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo)
	accountController := controller.NewAccountController(accountService)

	accountRG := rg.Group("/account")

	accountRG.GET("/:id", accountController.GetAccountById)
}
