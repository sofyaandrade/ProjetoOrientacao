package app

import (
	"fmt"
	"os"
	"projeto-back/api/middlewares"
	"projeto-back/api/routes"
	"projeto-back/utils"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func GinConfiguracao(db *gorm.DB, enforcer *casbin.Enforcer) {
	ginMode := viper.GetString("GIN_MODE")

	if EnvValues.AppEnv == "prod" {
		if ginMode == "" {
			ginMode = gin.DebugMode
		}
		gin.SetMode(ginMode)
	}

	serverPort := fmt.Sprintf(":%s", EnvValues.ApiPort)

	gin := gin.Default()
	gin.Use(middlewares.CORSMiddleware())
	routes.RouterConfiguracao(db, gin, enforcer)

	gin.GET("/documentacao/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	gin.Run(serverPort)

}

func GinConfiguracaoFront() {
	ginMode := viper.GetString("GIN_MODE")

	if EnvValues.AppEnv == "prod" {
		if ginMode == "" {
			ginMode = gin.DebugMode
		}
		gin.SetMode(ginMode)
	}

	if _, err := os.Stat("build"); !os.IsNotExist(err) {
		frontEndPort := fmt.Sprintf(":%s", EnvValues.ApiFrontPort)
		ginFrontEnd := gin.Default()

		ginFrontEnd.Use(static.Serve("/", static.LocalFile("./build", true)))
		ginFrontEnd.Run(frontEndPort)
	} else {
		utils.ErroLog.Println("Arquivo front não encontrado")
	}
}
