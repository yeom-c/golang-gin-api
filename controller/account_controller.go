package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/service"
)

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{
		accountService: service.NewAccountService(),
	}
}

func (c *AccountController) Get(ctx *gin.Context) {
	account, err := c.accountService.Get()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": account,
	})
}
