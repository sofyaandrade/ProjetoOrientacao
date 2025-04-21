package middlewares

import (
	"projeto-back/security"

	"github.com/gin-gonic/gin"
)

// comentado para não atrapalhar o teste das demais rotas, se não pede autenticação em todas
func JwtAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := SplitJWTToken(c)
		if len(t) == 2 {
			authToken := t[1]
			authorized, userID, _, err := security.ValidarToken(authToken)
			if err != nil {
				// c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
				return
			}
			if authorized {
				c.Set("id", userID)
				c.Next()
				return
			}
			// c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
			return
		}
		// c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Não autorizado"})
	}
}
