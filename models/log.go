package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Tipo      string
	Descricao string
}

type LogFiltro struct {
	DataInicio time.Time
	DataFinal  time.Time
	Tipo       string
}

type TipoEnum uint

const (
	ACAO_USUARIO TipoEnum = iota
	ACAO_SISTEMA
	ERRO
	ALERTA
)

type LogUsecase interface {
	CriarLog(log *Log) error
	BuscarLogs(logFiltro *LogFiltro) (*[]Log, error)
}

type LogRepository interface {
	CriarLog(log *Log) error
	BuscarLogsPorTipo(logFiltro *LogFiltro) (*[]Log, error)
	BuscarLogsPorPeriodo(logFiltro *LogFiltro) (*[]Log, error)
	BuscarLogsPorTipoEPeriodo(logFiltro *LogFiltro) (*[]Log, error)
}

func (s TipoEnum) String() string {
	return [...]string{"AÇÃO_USUARIO", "AÇÃO_SISTEMA", "ERRO", "ALERTA"}[s]
}

func (s TipoEnum) Index() int {
	return int(s)
}
