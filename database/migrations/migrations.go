package migrations

import (
	"projeto-back/models"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	var dbInstance *gorm.DB
	appEnv := viper.GetString("APP_ENV")

	if appEnv == "prod" {
		dbInstance = db
	} else {
		dbInstance = db.Debug()
	}

	dbInstance.AutoMigrate(&models.Usuario{})
	dbInstance.AutoMigrate(&models.Avaliacao{})
	dbInstance.AutoMigrate(&models.Cliente{})
}
