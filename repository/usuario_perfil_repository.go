package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type usuarioPerfilRepository struct {
	database *gorm.DB
}

func NovoUsuarioPerfilRepository(db *gorm.DB) models.UsuarioPerfilRepository {
	return &usuarioPerfilRepository{
		database: db,
	}
}

func (upr *usuarioPerfilRepository) CriarUsuarioPerfil(usuarioPerfil *models.UsuarioPerfil) error {
	err := upr.database.Create(&usuarioPerfil).Error
	return err
}

func (upr *usuarioPerfilRepository) BuscarTodosUsuariosPerfis() (*[]models.UsuarioPerfil, error) {
	var usuariosPerfis *[]models.UsuarioPerfil
	err := upr.database.Find(&usuariosPerfis).Error
	return usuariosPerfis, err
}

func (upr *usuarioPerfilRepository) BuscarUsuarioPerfilPorId(id uint) (*models.UsuarioPerfil, error) {
	var usuarioPerfil *models.UsuarioPerfil
	err := upr.database.First(&usuarioPerfil, id).Error
	return usuarioPerfil, err
}

func (upr *usuarioPerfilRepository) EditarUsuarioPerfil(usuarioPerfil *models.UsuarioPerfil, id uint) error {
	err := upr.database.Model(&usuarioPerfil).Where("id = ?", id).Updates(usuarioPerfil).Error
	return err
}

func (upr *usuarioPerfilRepository) DeletarUsuarioPerfil(id uint) error {
	var usuarioPerfil *models.UsuarioPerfil
	err := upr.database.Delete(&usuarioPerfil, id).Error
	return err
}

func (upr *usuarioPerfilRepository) ExisteUsuarioPerfilPorId(id uint) error {
	var usuarioPerfil *models.UsuarioPerfil
	err := upr.database.First(&usuarioPerfil, id).Error
	return err
}
