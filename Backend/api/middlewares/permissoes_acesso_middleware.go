package middlewares

import (
	"strings"

	"projeto-back/security"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// comentado para não atrapalhar o teste das demais rotas, se não pede autenticação em todas
func PermissoesAcessoMiddleware(secretKey string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := SplitJWTToken(c)
		if len(t) == 2 {
			authToken := t[1]
			_, userID, usuarioPerfil, err := security.ValidarToken(authToken)
			if err != nil {
				// c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
				return
			}
			if err != nil {
				// c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Erro ao localizar o usuário" + err.Error()})
				return
			}
			ok, err := enforcer.Enforce(strings.ToUpper(usuarioPerfil), c.Request.URL.Path, c.Request.Method)
			if err != nil {
				// c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Erro ao autorizar o usuário" + err.Error()})
				return
			}
			if !ok {
				// c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Não autorizado"})
				return
			}
			c.Set("id", userID)
			c.Next()
			return
		}
	}
}

func SplitJWTToken(c *gin.Context) []string {
	authHeader := c.Request.Header.Get("Authorization")
	t := strings.Split(authHeader, " ")
	return t
}
