package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome            string `json:"Nome"`
	Telefone        int    `json:"Telefone"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	UsuarioPerfilID uint   `json:"UsuarioPerfilID"`
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
