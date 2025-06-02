package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlunoRouter(db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NovoAlunoRepository(db)
	cc := &controllers.AlunoController{
		AlunoUsecase: usecase.NovoAlunoUsecase(cr),
	}
	group.POST("/", cc.NovoAluno)
	group.GET("/", cc.BuscarTodosAlunos)
	group.GET("/:id/", cc.BuscarAlunoPorId)
	group.PATCH("/:id/", cc.EditarAluno)
	group.DELETE("/:id/", cc.DeletarAluno)
}
