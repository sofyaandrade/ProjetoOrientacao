package usecase

import (
	"projeto-back/models"
	"projeto-back/security"
)

type refreshTokenUsecase struct {
	usuarioRepository models.UsuarioRepository
}

func NovoRefreshTokenUsecase(usuarioRepository models.UsuarioRepository) models.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		usuarioRepository: usuarioRepository,
	}
}

func (rtu *refreshTokenUsecase) BuscarUsuarioPorId(id uint) (*models.Usuario, error) {
	return rtu.usuarioRepository.BuscarUsuarioPorId(id)
}

func (rtu *refreshTokenUsecase) CriarTokenAcesso(usuario *models.Usuario) (accessToken string, err error) {
	return security.CriarTokenAcesso(usuario)
}

func (rtu *refreshTokenUsecase) CriarTokenAtualizacao(usuario *models.Usuario) (refreshToken string, err error) {
	return security.CriarTokenAtualizacao(usuario)
}
