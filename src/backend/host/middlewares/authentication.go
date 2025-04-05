package middlewares

import (
	"errors"
	"hoangcaophi/english-note-app/src/backend/global"
	"hoangcaophi/english-note-app/src/backend/shared"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	UserName string `json:"username"`
	UserId   string `json:"sub"`
	jwt.StandardClaims
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		claims := &AuthClaims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(global.Config.Authentication.Jwt.SecretKey), nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		id, _ := shared.StringToBsonBinary(claims.UserId)
		c.Set("userId", id)
		c.Next()
	}
}
