package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// type CustomClaims struct {
// 	username string
// 	jwt.RegisteredClaims
// }

func AccessToken(signature string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim := jwt.MapClaims{
			"username": "Test UserA",
			"exp":      jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		}

		// claim := CustomClaims{
		// 	"Test UserA",
		// 	jwt.RegisteredClaims{
		// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		// 	},
		// }

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		ss, err := token.SignedString([]byte(signature))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"eror": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": ss,
		})
	}
}
