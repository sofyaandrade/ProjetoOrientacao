package routes

import (
	"projeto-back/api/controllers"
	"projeto-back/repository"
	"projeto-back/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RefreshRouter(db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NovoUsuarioRepository(db)
	rtc := &controllers.RefreshTokenController{
		RefreshTokenUsecase: usecase.NovoRefreshTokenUsecase(ur),
		LoginUsecase:        usecase.NovoLoginUsecase(ur),
	}
	group.POST("/refresh/", rtc.RefreshToken)
}
