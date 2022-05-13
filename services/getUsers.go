package services

import (
	"iam/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	r := repos.GetIam(c)
	users := r.IamSvc.GetUsers(c)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
