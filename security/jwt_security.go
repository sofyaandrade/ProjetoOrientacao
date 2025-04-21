package security

import (
	"errors"
	"fmt"
	"time"

	"projeto-back/models"
	"projeto-back/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func CriarTokenAcesso(usuario *models.Usuario) (accessToken string, err error) {
	expiracao := viper.GetFloat64("TOKEN_SECRET_EXPIRY_HOUR")
	secretAccesKey := viper.GetString("TOKEN_SECRET")
	emissor := viper.GetString("TOKEN_EMISSOR")

	claims := &models.JwtCustomClaims{
		ID:   usuario.ID,
		Role: usuario.UsuarioPerfil.Descricao,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(utils.SolicitarNovoTempoEmMinutos(expiracao)),
			Issuer:    emissor,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretAccesKey))
	if err != nil {
		return "", err
	}
	return t, err
}

func CriarTokenAtualizacao(usuario *models.Usuario) (refreshToken string, err error) {
	secretRefreshKey := viper.GetString("TOKEN_REFRESH")
	expiracao := viper.GetFloat64("TOKEN_REFRESH_EXPIRY_HOUR")

	claimsRefresh := &models.JwtCustomRefreshClaims{
		ID: usuario.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(utils.SolicitarNovoTempoEmMinutos(expiracao)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secretRefreshKey))
	if err != nil {
		return "", err
	}
	return rt, err
}

func ValidarToken(accessToken string) (bool, uint, string, error) {
	secretAccesKey := viper.GetString("TOKEN_SECRET")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("log.metodo.assinar.inesperado", token.Header["alg"])
		}
		return []byte(secretAccesKey), nil
	})

	if err != nil {
		return false, 0, "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	switch {
	case token.Valid:
		return true, uint(claims["Id"].(float64)), claims["Role"].(string), err
	case errors.Is(err, jwt.ErrTokenMalformed):
		return false, 0, "", jwt.ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return false, 0, "", jwt.ErrTokenSignatureInvalid
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return false, 0, "", jwt.ErrTokenExpired
	default:
		return false, 0, "", err
	}
}

func ValidarTokenRefresh(refreshToken string) (bool, uint, error) {
	secretRefreshKey := viper.GetString("TOKEN_REFRESH")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("log.metodo.assinar.inesperado", token.Header["alg"])
		}
		return []byte(secretRefreshKey), nil
	})

	if err != nil {
		return false, 0, err
	}

	claims := token.Claims.(jwt.MapClaims)

	switch {
	case token.Valid:
		return true, uint(claims["Id"].(float64)), err
	case errors.Is(err, jwt.ErrTokenMalformed):
	default:
		return false, 0, err
	}
	return false, 0, err
}
