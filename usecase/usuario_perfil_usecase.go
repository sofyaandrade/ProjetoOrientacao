package usecase

import (
	"projeto-back/models"
)

type usuarioPerfilUsecase struct {
	usuarioPerfilRepository models.UsuarioPerfilRepository
}

func NovoUsuarioPerfilUsecase(usuarioPerfilRepository models.UsuarioPerfilRepository) models.UsuarioPerfilUsecase {
	return &usuarioPerfilUsecase{
		usuarioPerfilRepository: usuarioPerfilRepository,
	}
}

func (upu *usuarioPerfilUsecase) CriarUsuarioPerfil(usuarioPerfil *models.UsuarioPerfil) error {
	return upu.usuarioPerfilRepository.CriarUsuarioPerfil(usuarioPerfil)
}

func (upu *usuarioPerfilUsecase) BuscarTodosUsuariosPerfis() (*[]models.UsuarioPerfil, error) {
	return upu.usuarioPerfilRepository.BuscarTodosUsuariosPerfis()
}

func (upu *usuarioPerfilUsecase) BuscarUsuarioPerfilPorId(id uint) (*models.UsuarioPerfil, error) {
	return upu.usuarioPerfilRepository.BuscarUsuarioPerfilPorId(id)
}

func (upu *usuarioPerfilUsecase) EditarUsuarioPerfil(usuarioPerfil *models.UsuarioPerfil, id uint) error {
	return upu.usuarioPerfilRepository.EditarUsuarioPerfil(usuarioPerfil, id)
}

func (upu *usuarioPerfilUsecase) DeletarUsuarioPerfil(id uint) error {
	return upu.usuarioPerfilRepository.DeletarUsuarioPerfil(id)
}

func (upu *usuarioPerfilUsecase) ExisteUsuarioPerfilPorId(id uint) error {
	return upu.usuarioPerfilRepository.ExisteUsuarioPerfilPorId(id)
}
