package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccountRG(rg *gin.RouterGroup) {
	accountRG := rg.Group("/account")

	accountRG.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "account",
		})
	})
}
