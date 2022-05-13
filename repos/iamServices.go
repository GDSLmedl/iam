package repos

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type IamService interface {
	InitClient(c *gin.Context) IamService
	GetUsers(c *gin.Context) []User
}

type Iam struct {
	Provider string     `json:"provider"`
	IamSvc   IamService `json:"iamService"`
}

var lock = &sync.Mutex{}
var singleIam *Iam

func GetIam(ctx *gin.Context) *Iam {
	provider := os.Getenv("PROVIDER")
	if singleIam == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleIam == nil {
			if provider == "keycloak" {
				singleIam = &Iam{
					Provider: provider,
					IamSvc:   KeycloakIamService{}.InitClient(ctx),
				}
			}
		}
	}
	return singleIam
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
