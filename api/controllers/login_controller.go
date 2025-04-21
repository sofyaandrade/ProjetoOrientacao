package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"projeto-back/models"
	"projeto-back/security"
	"projeto-back/utils"
)

type LoginController struct {
	LoginUsecase models.LoginUsecase
}

// Login godoc
//
//	@Summary		Realiza Login
//	@Description	Rota para realizar login do usuário
//	@Tags			Autenticação
//	@Accept			json
//	@Produce		json
//	@Param			login	body		models.Login	true	"Modelo do login"
//	@Success		200		{object}	models.RespostaLogin
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/login/ [post]
func (lc *LoginController) Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	usuario, err := lc.LoginUsecase.BuscarUsuarioPorEmail(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao buscar usuário por email", err.Error())
		return
	}

	err = security.VerificaSenha(usuario.Password, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro credencial inválida", err.Error())
		return
	}

	accessToken, err := lc.LoginUsecase.CriarTokenAcesso(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao criar token acesso", err.Error())
		return
	}

	refreshToken, err := lc.LoginUsecase.CriarTokenAtualizacao(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao criar token de atualização", err.Error())
		return
	}

	respostaLogin := models.RespostaLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, respostaLogin)
}
