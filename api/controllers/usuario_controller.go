package controllers

import (
	"net/http"

	"projeto-back/models"
	"projeto-back/security"
	"projeto-back/utils"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	UsuarioUsecase models.UsuarioUsecase
}

func (uc *UsuarioController) NovoUsuario(c *gin.Context) {

	var usuarios models.Usuario

	if err := c.ShouldBindJSON(&usuarios); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	usuario, _ := uc.UsuarioUsecase.BuscarUsuarioPorEmail(usuarios.Email)

	if usuario.ID != 0 {
		utils.ErroLog.Println("Usuário já cadastrado")
		c.JSON(http.StatusBadRequest, models.WarningResponse{Warning: "Usuário já cadastrado"})
		return
	}

	hash, err := security.HashSenha(usuarios.Password)
	if err != nil {
		utils.ErroLog.Println("Erro no hash", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	usuarios.Password = string(hash)

	err = uc.UsuarioUsecase.CriarUsuario(&usuarios)
	if err != nil {
		utils.ErroLog.Println("Erro ao criar usuário", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Usuário criado com sucessp")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "lUsuário " + usuarios.Nome + " criado com sucesso"})
}

func (uc *UsuarioController) BuscarTodosUsuarios(c *gin.Context) {

	usuarios, err := uc.UsuarioUsecase.BuscarTodosUsuarios()
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar todos usuários", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if usuarios == nil {
		utils.ErroLog.Println("Não existe usuário cadastrado")
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Não existe usuario cadatsrado"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (uc *UsuarioController) BuscarUsuarioPorId(c *gin.Context) {

	id := c.Params.ByName("id")

	usuario, err := uc.UsuarioUsecase.BuscarUsuarioPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuario == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Usuário " + id + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioController) EditarUsuario(c *gin.Context) {
	var usuario *models.Usuario

	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&usuario); err != nil {
		utils.ErroLog.Println("Erro ao ler JSON", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	usuarioAntigo, err := uc.UsuarioUsecase.BuscarUsuarioPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuario.Password != "" {
		hash, err := security.HashSenha(usuario.Password)

		if err != nil {
			utils.ErroLog.Println("Erro no hash", err.Error())
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		usuario.Password = string(hash)
	}

	err = uc.UsuarioUsecase.EditarUsuario(usuario, utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao editar usuário", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if usuarioAntigo.Nome != usuario.Nome {
		utils.AcaoUsuario.Println("Usuário editado")
	}

	utils.AcaoUsuario.Println("Usuário editado com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Usuário " + id + " editado com sucesso"})

}

func (uc *UsuarioController) DeletarUsuario(c *gin.Context) {
	id := c.Params.ByName("id")

	usuario, err := uc.UsuarioUsecase.BuscarUsuarioPorId(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao buscar usuário por id", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err = uc.UsuarioUsecase.DeletarUsuario(utils.StringToUint(id))
	if err != nil {
		utils.ErroLog.Println("Erro ao deletar usuário", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	utils.AcaoUsuario.Println("Usuário " + usuario.Nome + " deletado com sucesso")
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Usuário " + id + " deletado com sucesso"})
}
