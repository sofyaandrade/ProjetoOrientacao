package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AvaliacaoRouter(db *gorm.DB, group *gin.RouterGroup) {
	ar := repository.NovaAvaliacaoRepository(db)
	ac := &controllers.AvaliacaoController{
		AvaliacaoUsecase: usecase.NovaAvaliacaoUsecase(ar),
	}
	group.POST("/", ac.NovaAvaliacao)
	group.GET("/", ac.BuscarTodasAvaliacoes)
	group.GET("/:id/", ac.BuscarAvaliacaoPorId)
	group.GET("/cliente/:id/", ac.BuscarAvaliacoesPorCliente)
	group.PATCH("/:id/", ac.EditarAvaliacao)
	group.DELETE("/:id/", ac.DeletarAvaliacao)
}
