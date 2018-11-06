package router

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwt "github.com/iceEvening/me/service/middleware/jwt"
)

func generateToken(c *gin.Context, ID uint) (string, error) {
	j := &jwt.JWT{
		SigningKey: []byte("iceEvening"),
	}
	claims := jwt.CustomClaims{
		ID: ID,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 3600),
			ExpiresAt: int64(time.Now().Unix() + 36000),
			Issuer:    "iceEvening",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
