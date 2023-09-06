package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeom-c/golang-gin-api/service"
)

type AccountController struct {
	service *service.AccountService
}

func NewAccountController(service *service.AccountService) *AccountController {
	return &AccountController{service}
}

func (c *AccountController) GetAccountById(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, _ := strconv.Atoi(sid)

	acc, err := c.service.GetAccountById(id)
	if err != nil {
		if err.Error() != "record not found" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	accDynamo, err := c.service.GetDynamoAccountById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": acc, "dynamoAccount": accDynamo})
}
