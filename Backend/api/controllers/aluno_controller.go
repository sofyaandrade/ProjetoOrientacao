package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
)

type AlunoController struct {
	AlunoUsecase models.AlunoUsecase
}

func (cc *AlunoController) NovoAluno(c *gin.Context) {

	var alunos models.Aluno

	if err := c.ShouldBindJSON(&alunos); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		return
	}

	aluno, _ := cc.AlunoUsecase.BuscarAlunoPorEmail(alunos.Email)

	if aluno.ID != 0 {
		utils.ErroLog.Println("Aluno ja cadastrado", aluno.ID)
		c.JSON(http.StatusBadRequest, models.WarningResponse{Warning: "Aluno já cadastrado"})
		return
	}

	err := cc.AlunoUsecase.CriarAluno(&alunos)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar cliete", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Aluno cadastrado com sucesso", aluno.Nome)
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Aluno" + alunos.Nome + " criado com sucesso"})
}

func (cc *AlunoController) BuscarTodosAlunos(c *gin.Context) {

	alunos, err := cc.AlunoUsecase.BuscarTodosAlunos()
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar todos os clietes", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if alunos == nil {
		utils.ErroLog.Println("Não existe usuário cadastrado", alunos)
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Não existe usuário cadastrado"})
		return
	}

	c.JSON(http.StatusOK, alunos)
}

func (cc *AlunoController) BuscarAlunoPorId(c *gin.Context) {

	id := c.Params.ByName("id")

	aluno, err := cc.AlunoUsecase.BuscarAlunoPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o cliete", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if aluno == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Aluno" + id + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func (cc *AlunoController) EditarAluno(c *gin.Context) {
	var aluno *models.Aluno

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&aluno); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	alunoAntigo, err := cc.AlunoUsecase.BuscarAlunoPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o aluno", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = cc.AlunoUsecase.EditarAluno(aluno, utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao editar o aluno", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if alunoAntigo.Nome != aluno.Nome {
		utils.AcaoUsuario.Println("Aluno editado de: ", alunoAntigo, "para: ", aluno.Nome)
	}

	utils.AcaoUsuario.Println("Aluno criado com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Aluno" + id + " editado com sucesso"})

}

func (cc *AlunoController) DeletarAluno(c *gin.Context) {
	id := c.Params.ByName("id")

	aluno, err := cc.AlunoUsecase.BuscarAlunoPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o aluno por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = cc.AlunoUsecase.DeletarAluno(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao deletar o aluno", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Aluno criado com sucesso", aluno.Nome)
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Aluno " + id + " deletado com sucesso"})
}
