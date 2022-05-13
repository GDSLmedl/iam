package main

import (
	"errors"
	"fmt"
	"iam/endpoints"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func verifyConfig() error {
	provider := os.Getenv("PROVIDER")
	validProviders := map[string]bool{
		"keycloak": true,
		"azuread":  false,
	}
	if !validProviders[provider] {
		return errors.New(fmt.Sprintf("[ERR_CONF] PROVIDER env var is not set correctly : '%s'", provider))
	}
	authUrl := os.Getenv("AUTH_URL")
	if authUrl == "" {
		return errors.New(fmt.Sprintf("[ERR_CONF] AUTH_URL env var is not set correctly : '%s'", authUrl))
	}
	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		return errors.New(fmt.Sprintf("[ERR_CONF] CLIENT_ID env var is not set correctly : '%s'", clientId))
	}
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientId == "" {
		return errors.New(fmt.Sprintf("[ERR_CONF] CLIENT_SECRET env var is not set correctly : '%s'", clientSecret))
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("An error occured when loading env file : %s\n", err)
	}
	err = verifyConfig()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/users", endpoints.GetUsers)
	r.Run(":8080")
}
