package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/soccer-manager/auth"
	"github.com/nvzard/soccer-manager/model"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("user", model.UserAuth{
			ID:     claims.ID,
			Email:  claims.Email,
			TeamID: claims.TeamID,
		})
		context.Next()
	}
}
