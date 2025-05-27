package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type logRepository struct {
	database *gorm.DB
}

func NovoLogRepository(db *gorm.DB) models.LogRepository {
	return &logRepository{
		database: db,
	}
}

func (lr *logRepository) CriarLog(log *models.Log) error {
	err := lr.database.Create(&log).Error
	return err
}

func (lr *logRepository) BuscarLogsPorTipo(logFiltro *models.LogFiltro) (*[]models.Log, error) {
	var logs *[]models.Log
	err := lr.database.Model(&models.Log{}).Where("tipo", logFiltro.Tipo).Order("created_at asc").Find(&logs).Error
	return logs, err
}

func (lr *logRepository) BuscarLogsPorPeriodo(logFiltro *models.LogFiltro) (*[]models.Log, error) {
	var logs *[]models.Log
	err := lr.database.Model(&models.Log{}).Where("created_at BETWEEN ? AND ?", logFiltro.DataInicio, logFiltro.DataFinal).Order("created_at asc").Find(&logs).Error
	return logs, err
}

func (lr *logRepository) BuscarLogsPorTipoEPeriodo(logFiltro *models.LogFiltro) (*[]models.Log, error) {
	var logs *[]models.Log
	err := lr.database.Model(&models.Log{}).Where("tipo", logFiltro.Tipo).Where("created_at BETWEEN ? AND ?", logFiltro.DataInicio, logFiltro.DataFinal).Order("created_at asc").Find(&logs).Error
	return logs, err
}
