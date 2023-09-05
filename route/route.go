package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(r *gin.Engine, db *gorm.DB) {
	versionRG := r.Group("/v1")
	AccountRG(versionRG, db)
}
