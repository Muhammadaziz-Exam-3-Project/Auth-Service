package middleware

import (
	"fmt"
	"net/http"
	"strings"

	t "Github.com/LocalEats/Api-get-way/api/token"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := ctx.Request.URL.Path
		fmt.Println("token : ", token)
		if strings.Contains(url, "swagger") || url == "/user/login" || url == "/user/register" {
			ctx.Next()
			return
		} else if _, err := t.ExtractClaim(token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()

	}
}
