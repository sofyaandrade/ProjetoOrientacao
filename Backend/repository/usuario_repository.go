package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type usuarioRepository struct {
	database *gorm.DB
}

func NovoUsuarioRepository(db *gorm.DB) models.UsuarioRepository {
	return &usuarioRepository{
		database: db,
	}
}

func (ur *usuarioRepository) CriarUsuario(usuario *models.Usuario) error {
	err := ur.database.Create(&usuario).Error
	return err
}

func (ur *usuarioRepository) BuscarTodosUsuarios() (*[]models.Usuario, error) {
	var usuarios *[]models.Usuario
	err := ur.database.Select("id", "nome", "telefone", "email", "usuario_perfil_id").Preload("UsuarioPerfil").Find(&usuarios).Error
	return usuarios, err
}

func (ur *usuarioRepository) BuscarUsuarioPorId(id uint) (*models.Usuario, error) {
	var usuario *models.Usuario
	err := ur.database.Select("id", "nome", "telefone", "email", "usuario_perfil_id").Preload("UsuarioPerfil").First(&usuario, id).Error
	return usuario, err
}

func (ur *usuarioRepository) BuscarUsuarioPorEmail(email string) (*models.Usuario, error) {
	var usuario *models.Usuario
	err := ur.database.Where(models.Usuario{Email: email}).Preload("UsuarioPerfil").First(&usuario).Error
	return usuario, err
}

func (ur *usuarioRepository) EditarUsuario(usuario *models.Usuario, id uint) error {
	err := ur.database.Model(&usuario).Where("id = ?", id).Updates(usuario).Error
	return err
}

func (ur *usuarioRepository) DeletarUsuario(id uint) error {
	var usuario *models.Usuario
	err := ur.database.Delete(&usuario, id).Error
	return err
}

func (ur *usuarioRepository) ExisteUsuarioPorId(id uint) error {
	var usuario *models.Usuario
	err := ur.database.First(&usuario, id).Error
	return err
}
