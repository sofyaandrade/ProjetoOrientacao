package models

type RefreshTokenRequest struct {
	RefreshToken string `form:"RefreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}
type RefreshTokenUsecase interface {
	BuscarUsuarioPorId(id uint) (*Usuario, error)
	CriarTokenAcesso(usuario *Usuario) (accessToken string, err error)
	CriarTokenAtualizacao(usuario *Usuario) (refreshToken string, err error)
}
