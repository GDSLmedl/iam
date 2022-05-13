package models

import "github.com/gin-gonic/gin"

type IamService interface {
	InitClient(c *gin.Context) IamService
	// Users functions
	GetUsers(c *gin.Context) []User
	// GetUser(c *gin.Context, userId string) User
	// GetPromoUsers(c *gin.Context, promo string) []User
	// GetGroupMembers(c *gin.Context, group string) []User
}

type Iam struct {
	Provider string     `json:"provider"`
	IamSvc   IamService `json:"iamService"`
}

type User struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Groups    []string `json:"groups,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	Promotion string   `json:"promotion,omitempty"`
}
