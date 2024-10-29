package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	"github.com/AthThobari/simple_music_catalog_go/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	secretKey:= configs.Get().Service.SecretKey
	return func(c *gin.Context) {
		header:= c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
 }
 } 