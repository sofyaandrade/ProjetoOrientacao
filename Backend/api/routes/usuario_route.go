package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsuarioRouter(db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NovoUsuarioRepository(db)
	uc := &controllers.UsuarioController{
		UsuarioUsecase: usecase.NovoUsuarioUsecase(ur),
	}
	group.POST("/", uc.NovoUsuario)
	group.GET("/", uc.BuscarTodosUsuarios)
	group.GET("/:id/", uc.BuscarUsuarioPorId)
	group.PATCH("/:id/", uc.EditarUsuario)
	group.DELETE("/:id/", uc.DeletarUsuario)
}
