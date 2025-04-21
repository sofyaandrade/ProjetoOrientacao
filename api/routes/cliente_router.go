package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClienteRouter(db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NovoClienteRepository(db)
	cc := &controllers.ClienteController{
		ClienteUsecase: usecase.NovoClienteUsecase(cr),
	}
	group.POST("/", cc.NovoCliente)
	group.GET("/", cc.BuscarTodosClientes)
	group.GET("/:id/", cc.BuscarClientePorId)
	group.PATCH("/:id/", cc.EditarCliente)
	group.DELETE("/:id/", cc.DeletarCliente)
}
