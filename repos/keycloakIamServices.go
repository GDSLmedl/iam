package repos

import (
	"fmt"
	"os"

	"iam/models"

	keycloak "github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"
)

type KeycloakIamService struct{}

var keycloakInstance KeycloakIamService
var iamClient keycloak.GoCloak
var keycloakJWT keycloak.JWT

func (iamSvc KeycloakIamService) InitClient(ctx *gin.Context) models.IamService {
	iamClient = keycloak.NewClient(os.Getenv("AUTH_URL"))
	jwt, err := iamClient.LoginClient(ctx, os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("AUTH_REALM"))
	if err != nil {
		// Handle so that we do not panic
		panic("Could not connect to Keycloak Realm")
	}
	keycloakJWT = *jwt
	keycloakInstance = KeycloakIamService{}
	return keycloakInstance
}

func (iamSvc KeycloakIamService) GetUsers(ctx *gin.Context) []models.User {
	usersK, err := iamClient.GetUsers(ctx, keycloakJWT.AccessToken, "SIWEB", keycloak.GetUsersParams{})
	if err != nil {
		fmt.Errorf("Could not GET Users : %s", err)
		return nil
	}
	return convertUsers(usersK)
}

func convertUsers(usersK []*keycloak.User) []models.User {
	users := make([]models.User, 0, len(usersK)+1)
	for _, u := range usersK {
		users = append(users, convertUser(u))
	}
	return users
}

func getValueOrDefault[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

func convertUser(userK *keycloak.User) models.User {
	user := &models.User{
		ID:        getValueOrDefault(userK.ID, ""),
		FirstName: getValueOrDefault(userK.FirstName, ""),
		LastName:  getValueOrDefault(userK.LastName, ""),
		Email:     getValueOrDefault(userK.Email, ""),
	}
	return *user
}
