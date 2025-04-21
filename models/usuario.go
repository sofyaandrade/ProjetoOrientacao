package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome            string `json:"nome"`
	Telefone        string `json:"telefone"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	UsuarioPerfilID uint   `json:"usuario_perfil_id"`
	UsuarioPerfil   UsuarioPerfil
}

type UsuarioRepository interface {
	CriarUsuario(usuarios *Usuario) error
	BuscarTodosUsuarios() (*[]Usuario, error)
	BuscarUsuarioPorId(id uint) (*Usuario, error)
	BuscarUsuarioPorEmail(email string) (*Usuario, error)
	EditarUsuario(usuario *Usuario, id uint) error
	DeletarUsuario(id uint) error
	ExisteUsuarioPorId(id uint) error
}

type UsuarioUsecase interface {
	CriarUsuario(usuarios *Usuario) error
	BuscarTodosUsuarios() (*[]Usuario, error)
	BuscarUsuarioPorId(id uint) (*Usuario, error)
	BuscarUsuarioPorEmail(email string) (*Usuario, error)
	EditarUsuario(usuario *Usuario, id uint) error
	DeletarUsuario(id uint) error
	ExisteUsuarioPorId(id uint) error
}
