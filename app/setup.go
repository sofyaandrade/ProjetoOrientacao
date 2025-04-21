package app

import (
	"log"
	"projeto-back/database"
	"projeto-back/utils"

	"gorm.io/gorm"
)

var (
	Database  *gorm.DB
	EnvValues *Env
	err       error
)

func Run() {
	utils.InicializaLog()

	EnvValues, err = CarregarValoresEnv()
	if err != nil {
		utils.ErroLog.Println("Erro oa carregar os valores do env", err)
		log.Fatalf("Erro oa carregar os valores do env, %v", err)
	}

	Database = database.InicializarDatabase(EnvValues.DbDriver, EnvValues.DbUser, EnvValues.DbPassword, EnvValues.DbPort, EnvValues.DbHost, EnvValues.DbName)

	enforcer := PermissoesAcessoConfiguracao(Database)

	go GinConfiguracaoFront()

	GinConfiguracao(Database, enforcer)

}
