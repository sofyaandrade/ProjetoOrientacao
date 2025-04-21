package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	ClienteUsecase models.ClienteUsecase
}

func (cc *ClienteController) NovoCliente(c *gin.Context) {

	var clientes models.Cliente

	if err := c.ShouldBindJSON(&clientes); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		return
	}

	cliente, _ := cc.ClienteUsecase.BuscarClientePorEmail(clientes.Email)

	if cliente.ID != 0 {
		utils.ErroLog.Println("Cliente ja cadastrado", cliente.ID)
		c.JSON(http.StatusBadRequest, models.WarningResponse{Warning: "Cliente já cadastrado"})
		return
	}

	err := cc.ClienteUsecase.CriarCliente(&clientes)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar cliete", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Cliente cadastrado com sucesso", cliente.Nome)
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Cliente" + clientes.Nome + " criado com sucesso"})
}

func (cc *ClienteController) BuscarTodosClientes(c *gin.Context) {

	clientes, err := cc.ClienteUsecase.BuscarTodosClientes()
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar todos os clietes", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if clientes == nil {
		utils.ErroLog.Println("Não existe usuário cadastrado", clientes)
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Não existe usuário cadastrado"})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

func (cc *ClienteController) BuscarClientePorId(c *gin.Context) {

	id := c.Params.ByName("id")

	cliente, err := cc.ClienteUsecase.BuscarClientePorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o cliete", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if cliente == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Cliente" + id + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (cc *ClienteController) EditarCliente(c *gin.Context) {
	var cliente *models.Cliente

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&cliente); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	clienteAntigo, err := cc.ClienteUsecase.BuscarClientePorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o cliente", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = cc.ClienteUsecase.EditarCliente(cliente, utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao editar o cliente", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if clienteAntigo.Nome != cliente.Nome {
		utils.AcaoUsuario.Println("CLiente editado de: ", clienteAntigo, "para: ", cliente.Nome)
	}

	utils.AcaoUsuario.Println("Cliente criado com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Cliente" + id + " editado com sucesso"})

}

func (cc *ClienteController) DeletarCliente(c *gin.Context) {
	id := c.Params.ByName("id")

	cliente, err := cc.ClienteUsecase.BuscarClientePorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar o cliente por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = cc.ClienteUsecase.DeletarCliente(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao deletar o cliente", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Cliente criado com sucesso", cliente.Nome)
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Cliente " + id + " deletado com sucesso"})
}
