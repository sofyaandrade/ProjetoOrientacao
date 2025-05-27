package usecase

import (
	"log"
	"os"
	"path/filepath"
)

type SimpleLogUsecase struct {
	logger *log.Logger
}

func NewSimpleLogUsecase() (*SimpleLogUsecase, error) {
	err := os.MkdirAll("tmp", os.ModePerm)
	if err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(filepath.Join("tmp", "PROJETO-Back-logs.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return &SimpleLogUsecase{
		logger: logger,
	}, nil
}

func (l *SimpleLogUsecase) CriarLog(msg string) {
	l.logger.Println(msg)
}
