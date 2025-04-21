package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RefreshTokenController struct {
	RefreshTokenUsecase models.RefreshTokenUsecase
	LoginUsecase        models.LoginUsecase
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {

	var refreshToken = viper.GetString("TOKEN_REFRESH")

	var request models.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	_, id, err := rtc.LoginUsecase.ValidarRefreshToken(request.RefreshToken)
	if err != nil {
		utils.ErroLog.Println("Usuário não encontrado", err.Error())
		c.JSON(http.StatusUnauthorized, models.WarningResponse{Warning: "Usuário não encontrado"})
		return
	}

	if id == 0 {
		utils.ErroLog.Println("Usuário não encontrado", err.Error())
		c.JSON(http.StatusUnauthorized, models.WarningResponse{Warning: "Usuário não encontrado"})
		return
	}

	usuario, err := rtc.RefreshTokenUsecase.BuscarUsuarioPorId(id)
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário por id", err.Error())
		c.JSON(http.StatusUnauthorized, models.WarningResponse{Warning: "Usuário não encontrado"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CriarTokenAcesso(usuario)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar token acesso", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	refreshToken, err = rtc.RefreshTokenUsecase.CriarTokenAtualizacao(usuario)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar token atualização acesso", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	refreshTokenResponse := models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
