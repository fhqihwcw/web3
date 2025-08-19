package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleInit(c *gin.Context) {
	c.Set("requestID", uuid.NewString())
	log.Printf("[%s] %s %s", c.GetString("requestID"), c.Request.Method, c.Request.URL)
	auth := c.GetHeader("Authorization")
	url := c.Request.URL.Path
	if url == "/users/register" || url == "/users/login" {
		c.Next()
		return
	}
	//判断jwt是否有效

	if auth == "" {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "auth is empty",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("test"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "token error: " + err.Error(),
		})
		log.Println("JWT parse error:", err)
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("userID", claims["userID"])
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}

	c.Next()
}
