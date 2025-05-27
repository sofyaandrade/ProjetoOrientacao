package app

import (
	"fmt"
	"strings"

	"projeto-back/utils"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ApiPort                string `mapstructure:"API_PORT"`
	ApiFrontPort           string `mapstructure:"API_FRONT_PORT"`
	DbHost                 string `mapstructure:"DB_HOST"`
	DbDriver               string `mapstructure:"DB_DRIVER"`
	DbUser                 string `mapstructure:"DB_USER"`
	DbPassword             string `mapstructure:"DB_PASSWORD"`
	DbName                 string `mapstructure:"DB_NAME"`
	DbPort                 string `mapstructure:"DB_PORT"`
	ApiPassword            string `mapstructure:"API_PASSWORD"`
	TokenSecretExpiryHour  string `mapstructure:"TOKEN_SECRET_EXPIRY_HOUR"`
	TokenRefreshExpiryHour string `mapstructure:"TOKEN_REFRESH_EXPIRY_HOUR"`
	TokenSecret            string `mapstructure:"TOKEN_SECRET"`
	TokenRefresh           string `mapstructure:"TOKEN_REFRESH"`
	TokenEmissor           string `mapstructure:"TOKEN_EMISSOR"`
}

const (
	SQLITE          = "sqlite"
	POSTGRES        = "postgres"
	STANDALONE      = "STANDALONE"
	PROD            = "PROD"
	DEV             = "DEV"
	PROJETOBACK_ENV = "projeto-back.env"
)

func CarregarValoresEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile(PROJETOBACK_ENV)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Não foi possivel localizar o arquivo %s: %s", PROJETOBACK_ENV, err)
		utils.ErroLog.Printf("Não foi possivel localizar o arquivo %s: %s", PROJETOBACK_ENV, err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		utils.ErroLog.Println("Não foi possível carregar o ambiente:", err)
	}

	if strings.Contains(strings.ToUpper(env.AppEnv), DEV) {
		fmt.Println("ambiente rodando em modo desenvolvimento")
	} else if strings.Contains(strings.ToUpper(env.AppEnv), PROD) {
		fmt.Println("ambiente rodando em modo produção")
	} else if strings.Contains(strings.ToUpper(env.AppEnv), STANDALONE) {
		env.DbDriver = SQLITE
		fmt.Println("ambiente rodando em modo standalone")
	} else {
		env.DbDriver = POSTGRES
		fmt.Println("ambiente rodando em modo produção")
	}
	return &env, err
}
