package database

import (
	"fmt"

	"projeto-back/database/migrations"
	"projeto-back/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES = "postgres"
)

func InicializarDatabase(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	var (
		DB  *gorm.DB
		err error
	)
	if DbDriver == POSTGRES {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("não conectou ao banco de dados %s: %v\n", DbDriver, err)
			utils.ErroLog.Printf("não conectou ao banco de dados %s: %v", DbDriver, err)
			return nil
		} else {
			fmt.Printf("conectou ao banco de dados %s\n", DbDriver)
			utils.AcaoSistema.Printf("conectou ao banco de dados %s", DbDriver)
		}
	} else {
		fmt.Println("log.banco.desconhecido")
		utils.ErroLog.Println("log.banco.desconhecido")
	}
	migrations.RunMigrations(DB)

	migrations.InicializaUsuariosBasicos(DB)
	migrations.InicializaUsuarioPerfil(DB)
	return DB
}
