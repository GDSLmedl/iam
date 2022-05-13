package endpoints

import (
	"iam/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Call Service GetUsers
	services.GetUsers(c)
}
