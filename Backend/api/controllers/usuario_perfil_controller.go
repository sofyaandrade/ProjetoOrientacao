package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
)

type UsuarioPerfilController struct {
	UsuarioPerfilUsecase models.UsuarioPerfilUsecase
	UsuarioUsecase       models.UsuarioUsecase
}

func (upc *UsuarioPerfilController) NovoUsuarioPerfil(c *gin.Context) {

	var usuarioPerfil models.UsuarioPerfil

	if err := c.ShouldBindJSON(&usuarioPerfil); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err := upc.UsuarioPerfilUsecase.CriarUsuarioPerfil(&usuarioPerfil)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar usuário perfil", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("lPerfil de usuário" + usuarioPerfil.Descricao + " criado")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Perfil usuário " + usuarioPerfil.Descricao + " criado com sucesso"})
}

func (upc *UsuarioPerfilController) BuscarTodosUsuariosPerfis(c *gin.Context) {

	usuariosPerfis, err := upc.UsuarioPerfilUsecase.BuscarTodosUsuariosPerfis()
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar todos os usuários perfil", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if usuariosPerfis == nil {
		utils.ErroLog.Println("Não existe perfil cadastrado")
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Não existe perfil cadastrado"})
		return
	}

	c.JSON(http.StatusOK, usuariosPerfis)
}

func (upc *UsuarioPerfilController) BuscarUsuarioPerfilPorId(c *gin.Context) {

	id := c.Params.ByName("id")

	usuarioPerfil, err := upc.UsuarioPerfilUsecase.BuscarUsuarioPerfilPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário perfil por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuarioPerfil == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Perfil usuário" + id + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuarioPerfil)
}

func (upc *UsuarioPerfilController) BuscarUsuarioPerfilPorIdUsuario(c *gin.Context) {

	id := c.Params.ByName("id")

	usuario, err := upc.UsuarioUsecase.BuscarUsuarioPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário perfil por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuario == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Usuário " + id + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario.UsuarioPerfil.Descricao)
}

func (upc *UsuarioPerfilController) EditarUsuarioPerfil(c *gin.Context) {
	var usuarioPerfil *models.UsuarioPerfil

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&usuarioPerfil); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	usuarioPerfilAntigo, err := upc.UsuarioPerfilUsecase.BuscarUsuarioPerfilPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário perfil por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = upc.UsuarioPerfilUsecase.EditarUsuarioPerfil(usuarioPerfil, utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao editar usuário perfil", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuarioPerfilAntigo.Descricao != usuarioPerfil.Descricao {
		utils.AcaoUsuario.Println("Perfil usuário " + usuarioPerfilAntigo.Descricao + " renomeado para: " + usuarioPerfil.Descricao)
	}

	utils.AcaoUsuario.Println("Perfil usuário" + usuarioPerfil.Descricao + " editado")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Perfil usuário" + id + " editado com sucesso"})

}

func (upc *UsuarioPerfilController) DeletarUsuarioPerfil(c *gin.Context) {
	id := c.Params.ByName("id")

	usuarioPerfil, err := upc.UsuarioPerfilUsecase.BuscarUsuarioPerfilPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = upc.UsuarioPerfilUsecase.DeletarUsuarioPerfil(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao deletar usuário perfil", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Perfil usuário" + usuarioPerfil.Descricao + " deletado")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Perfil usuário" + id + " deletado com sucesso"})
}
