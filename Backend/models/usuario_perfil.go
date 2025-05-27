package models

type UsuarioPerfil struct {
	ID        uint
	Descricao string    `json:"descricao"`
	Usuarios  []Usuario `gorm:"foreignkey:usuario_perfil_id"`
}

type UsuarioPerfilRepository interface {
	CriarUsuarioPerfil(usuarioPerfil *UsuarioPerfil) error
	BuscarTodosUsuariosPerfis() (*[]UsuarioPerfil, error)
	BuscarUsuarioPerfilPorId(id uint) (*UsuarioPerfil, error)
	EditarUsuarioPerfil(usuarioPerfil *UsuarioPerfil, id uint) error
	DeletarUsuarioPerfil(id uint) error
	ExisteUsuarioPerfilPorId(id uint) error
}

type UsuarioPerfilUsecase interface {
	CriarUsuarioPerfil(usuarioPerfil *UsuarioPerfil) error
	BuscarTodosUsuariosPerfis() (*[]UsuarioPerfil, error)
	BuscarUsuarioPerfilPorId(id uint) (*UsuarioPerfil, error)
	EditarUsuarioPerfil(usuarioPerfil *UsuarioPerfil, id uint) error
	DeletarUsuarioPerfil(id uint) error
	ExisteUsuarioPerfilPorId(id uint) error
}
