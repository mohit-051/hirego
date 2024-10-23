package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mohit-051/hirego/internal/helper"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Inside the Auth middleware", c.GetHeader("Authorization"))
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		accessToken := parts[1]

		fmt.Println("Access token: ", accessToken)
		userName, err := helper.VerifyToken(accessToken)
		if err != nil {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		fmt.Println("User email: ", userName)

		// Set the user email in the context for later use
		c.Set("username", userName)

		c.Next()
	}
}
