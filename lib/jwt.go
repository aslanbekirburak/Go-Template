package jwt

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		Token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		c.Set("token-claims", Token.Claims)

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}
