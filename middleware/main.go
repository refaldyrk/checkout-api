package middleware

import (
	"checkout-api/service"
	"checkout-api/utils/http_response"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(respWriter http_response.IResponseWriter, authService service.IAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			respWriter.HTTPJsonErr(c, 401, "Authorization header is missing", "", nil)
			c.Abort()
			return
		}

		tokenSplit := strings.Split(tokenString, " ")
		if len(tokenSplit) != 2 && tokenSplit[0] != "Bearer" {
			respWriter.HTTPJsonErr(c, 401, "invalid token", "", nil)
			c.Abort()
			return
		}

		token := tokenSplit[1]
		user, err := authService.CheckToken(token)
		if err != nil {
			respWriter.HTTPCustomErr(c, err)
			c.Abort()
			return
		}

		// Pass the claims to subsequent handlers
		c.Set("user_id", user.ID)
		c.Set("username", user.Username)

		c.Next()
	}
}
