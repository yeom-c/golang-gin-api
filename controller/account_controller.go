package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (c *AccountController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "account",
	})
}
