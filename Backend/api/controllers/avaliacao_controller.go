package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
)

type AvaliacaoController struct {
	AvaliacaoUsecase models.AvaliacaoUsecase
}

func (ac *AvaliacaoController) NovaAvaliacao(c *gin.Context) {

	var avaliacao models.Avaliacao

	if err := c.ShouldBindJSON(&avaliacao); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		return
	}

	err := ac.AvaliacaoUsecase.CriarAvaliacao(&avaliacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao criar a avaliacao", err.Error())
		return
	}

	utils.AcaoUsuario.Println("Avaliação criada com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Avaliacao criada com sucesso"})
}

func (ac *AvaliacaoController) BuscarTodasAvaliacoes(c *gin.Context) {

	avaliacao, err := ac.AvaliacaoUsecase.BuscarTodasAvaliacoes()
	if err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if avaliacao == nil {
		utils.ErroLog.Println("Nao existem avaliações cadastradas", err.Error())
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Não existe avaliacoes cadastradas"})
		return
	}

	c.JSON(http.StatusOK, avaliacao)
}

func (ac *AvaliacaoController) BuscarAvaliacaoPorId(c *gin.Context) {

	id := c.Params.ByName("id")

	cliente, err := ac.AvaliacaoUsecase.BuscarAvaliacaoPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar avaliacao por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if cliente == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Avaliação" + id + "não encontrada"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (ac *AvaliacaoController) BuscarAvaliacoesPorCliente(c *gin.Context) {
	id := c.Params.ByName("id")

	avaliacoes, err := ac.AvaliacaoUsecase.BuscarAvaliacoesPorCliente(utils.StringToUint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, avaliacoes)
}

func (ac *AvaliacaoController) EditarAvaliacao(c *gin.Context) {
	var avaliacao *models.Avaliacao

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&avaliacao); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	avaliacaoAntiga, err := ac.AvaliacaoUsecase.BuscarAvaliacaoPorId(utils.StringToUint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao buscar avaliação por id", err.Error())
		return
	}

	err = ac.AvaliacaoUsecase.EditarAvaliacao(avaliacao, utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao editar avaliação", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if avaliacaoAntiga != avaliacao {
		utils.AcaoUsuario.Println("Avaliação editada de: ", avaliacaoAntiga, "para: ", avaliacao)
	}

	utils.AcaoUsuario.Println("Avaliação editada com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Avaliação" + id + " editada com sucesso"})

}

func (ac *AvaliacaoController) DeletarAvaliacao(c *gin.Context) {
	id := c.Params.ByName("id")

	avaliacao, err := ac.AvaliacaoUsecase.BuscarAvaliacaoPorId(utils.StringToUint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		return
	}

	err = ac.AvaliacaoUsecase.DeletarAvaliacao(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao deletar avaliação", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Avaliação deletada com sucesso", avaliacao.ID)
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Avaliação" + id + " deletada com sucesso"})
}
