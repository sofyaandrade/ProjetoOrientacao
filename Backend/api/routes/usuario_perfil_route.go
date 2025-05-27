package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsuarioPerfilRoute(db *gorm.DB, group *gin.RouterGroup) {
	upr := repository.NovoUsuarioPerfilRepository(db)
	ur := repository.NovoUsuarioRepository(db)
	upc := &controllers.UsuarioPerfilController{
		UsuarioPerfilUsecase: usecase.NovoUsuarioPerfilUsecase(upr),
		UsuarioUsecase:       usecase.NovoUsuarioUsecase(ur),
	}
	group.POST("/", upc.NovoUsuarioPerfil)
	group.GET("/", upc.BuscarTodosUsuariosPerfis)
	group.GET("/:id/", upc.BuscarUsuarioPerfilPorId)
	group.GET("/usuario/:id/", upc.BuscarUsuarioPerfilPorIdUsuario)
	group.PATCH("/:id/", upc.EditarUsuarioPerfil)
	group.DELETE("/:id/", upc.DeletarUsuarioPerfil)
}
