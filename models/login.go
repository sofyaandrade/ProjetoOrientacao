package models

type Login struct {
	Email    string
	Password string
}

type RespostaLogin struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type LoginUsecase interface {
	BuscarUsuarioPorEmail(email string) (*Usuario, error)
	CriarTokenAcesso(usuario *Usuario) (accessToken string, err error)
	CriarTokenAtualizacao(usuario *Usuario) (refreshToken string, err error)
	ValidarToken(token string) (tokenValido bool, idToken uint, perfil string, err error)
	ValidarRefreshToken(token string) (tokenValido bool, idToken uint, err error)
}
