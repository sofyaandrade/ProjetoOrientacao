package usecase

import (
	"projeto-back/models"
)

type usuariosUsecase struct {
	usuarioRepository models.UsuarioRepository
}

func NovoUsuarioUsecase(usuarioRepository models.UsuarioRepository) models.UsuarioUsecase {
	return &usuariosUsecase{
		usuarioRepository: usuarioRepository,
	}
}

func (uu *usuariosUsecase) CriarUsuario(usuario *models.Usuario) error {
	return uu.usuarioRepository.CriarUsuario(usuario)
}

func (uu *usuariosUsecase) BuscarTodosUsuarios() (*[]models.Usuario, error) {
	return uu.usuarioRepository.BuscarTodosUsuarios()
}

func (uu *usuariosUsecase) BuscarUsuarioPorId(id uint) (*models.Usuario, error) {
	return uu.usuarioRepository.BuscarUsuarioPorId(id)
}

func (uu *usuariosUsecase) BuscarUsuarioPorEmail(email string) (*models.Usuario, error) {
	return uu.usuarioRepository.BuscarUsuarioPorEmail(email)
}

func (uu *usuariosUsecase) EditarUsuario(usuario *models.Usuario, id uint) error {
	return uu.usuarioRepository.EditarUsuario(usuario, id)
}

func (uu *usuariosUsecase) DeletarUsuario(id uint) error {
	return uu.usuarioRepository.DeletarUsuario(id)
}

func (uu *usuariosUsecase) ExisteUsuarioPorId(id uint) error {
	return uu.usuarioRepository.ExisteUsuarioPorId(id)
}
