package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/controller"
	"github.com/yeom-c/golang-gin-api/database"
	"github.com/yeom-c/golang-gin-api/repository"
	"github.com/yeom-c/golang-gin-api/service"
)

func AccountRG(rg *gin.RouterGroup) {
	accountRepo := repository.NewAccountRepository(database.DB.Gorm)
	accountDynamoRepo := repository.NewAccountDynamoRepository(database.DB.Dynamo)
	accountService := service.NewAccountService(accountRepo, accountDynamoRepo)
	accountController := controller.NewAccountController(accountService)

	accountRG := rg.Group("/account")

	accountRG.GET("/:id", accountController.GetAccountById)
}
