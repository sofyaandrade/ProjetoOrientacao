package utils

import (
	"log"
	"os"
	"path/filepath"
)

var (
	AcaoUsuario *log.Logger
	AcaoSistema *log.Logger
	Alerta      *log.Logger
	ErroLog     *log.Logger
)

func InicializaLog() {
	LOG_FILE := "tmp/PROJETO-Back-logs.log"

	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		newpath := filepath.Join(".", "tmp")
		err := os.Mkdir(newpath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}

	flags := log.Lshortfile | log.LstdFlags

	AcaoUsuario = log.New(logFile, "AÇÃO DO USUARIO: ", flags)
	AcaoSistema = log.New(logFile, "AÇÃO DO SISTEMA: ", flags)
	ErroLog = log.New(logFile, "ERRO: ", flags)
}
