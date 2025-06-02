package routes

import (
	"projeto-back/api/middlewares"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func RouterConfiguracao(db *gorm.DB, router *gin.Engine, enforcer *casbin.Enforcer) {

	secretKey := viper.GetString("TOKEN_SECRET")

	router.Use(middlewares.PermissoesAcessoMiddleware(secretKey, enforcer))

	publicRouter := router.Group("")
	LoginRouter(db, publicRouter)
	RefreshRouter(db, publicRouter)

	protectedUsuariosRouter := router.Group("/usuarios")
	protectedUsuariosRouter.Use(middlewares.JwtAuthMiddleware(secretKey))
	UsuarioRouter(db, protectedUsuariosRouter)

	protectedUsuarioPerfilRouter := router.Group("/usuario-perfil")
	protectedUsuarioPerfilRouter.Use(middlewares.JwtAuthMiddleware(secretKey))
	UsuarioPerfilRoute(db, protectedUsuarioPerfilRouter)

	protectedAlunoRouter := router.Group("/aluno")
	protectedAlunoRouter.Use(middlewares.JwtAuthMiddleware(secretKey))
	AlunoRouter(db, protectedAlunoRouter)

	protectedAvaliacaoRouter := router.Group("/avaliacao")
	protectedAvaliacaoRouter.Use(middlewares.JwtAuthMiddleware(secretKey))
	AvaliacaoRouter(db, protectedAvaliacaoRouter)
}
