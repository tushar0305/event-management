package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/utils"
)

func Authenticate(context *gin.Context) {
    token := context.Request.Header.Get("Authorization")
    if token == "" {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized!"})
        context.Abort()
        return
    }

    // Remove "Bearer " prefix if present
    token = strings.TrimPrefix(token, "Bearer ")

    _, err := utils.VerifyToken(token)
    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized!"})
        context.Abort()
        return
    }

    context.Next()
}