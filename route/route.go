package route

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	versionRG := r.Group("/v1")
	AccountRG(versionRG)
}
