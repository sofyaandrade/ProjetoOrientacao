package usecase

import (
	"projeto-back/models"
	"projeto-back/security"
)

type loginUsecase struct {
	usuarioRepository models.UsuarioRepository
}

func NovoLoginUsecase(usuarioRepository models.UsuarioRepository) models.LoginUsecase {
	return &loginUsecase{
		usuarioRepository: usuarioRepository,
	}
}

func (lu *loginUsecase) BuscarUsuarioPorEmail(email string) (*models.Usuario, error) {
	return lu.usuarioRepository.BuscarUsuarioPorEmail(email)
}

func (lu *loginUsecase) CriarTokenAcesso(usuarios *models.Usuario) (accessToken string, err error) {
	return security.CriarTokenAcesso(usuarios)
}

func (lu *loginUsecase) CriarTokenAtualizacao(usuarios *models.Usuario) (accessToken string, err error) {
	return security.CriarTokenAtualizacao(usuarios)
}

func (lu *loginUsecase) ValidarToken(token string) (tokenValido bool, idToken uint, perfil string, err error) {
	isTokenValido, idToken, perfil, err := security.ValidarToken(token)
	return isTokenValido, idToken, perfil, err
}

func (lu *loginUsecase) ValidarRefreshToken(token string) (tokenValido bool, idToken uint, err error) {
	isTokenValido, idToken, err := security.ValidarTokenRefresh(token)
	return isTokenValido, idToken, err
}
