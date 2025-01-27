package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// func Protect(tokenString string) error {
// 	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
// 		}
// 		return []byte("==signature=="), nil
// 	})
// 	return err
// }

// Middleware closure function
func Protect(signature []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		auth := ctx.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return signature, nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claim, ok := token.Claims.(jwt.MapClaims); ok {
			userName := claim["username"]
			ctx.Set("username", userName)

		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}

}
