package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginRouter(db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NovoUsuarioRepository(db)
	lc := &controllers.LoginController{
		LoginUsecase: usecase.NovoLoginUsecase(ur),
	}
	group.POST("/login/", lc.Login)
}
