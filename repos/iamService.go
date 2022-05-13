package repos

import (
	"os"
	"sync"

	"iam/models"

	"github.com/gin-gonic/gin"
)

var lock = &sync.Mutex{}
var singleIam *models.Iam

func GetIam(ctx *gin.Context) *models.Iam {
	provider := os.Getenv("PROVIDER")
	if singleIam == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleIam == nil {
			if provider == "keycloak" {
				singleIam = &models.Iam{
					Provider: provider,
					IamSvc:   KeycloakIamService{}.InitClient(ctx),
				}
			}
		}
	}
	return singleIam
}
