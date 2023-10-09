package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* Defines a map for storing encrypted users and passwords. "Hard coded" information for test only,
* for prod we recommend database storage.
 */
var users = map[string]string{
	"dc": "$2a$10$jXRtMagkTkC3lxO3ynj.FeEXxtirLMvlbsl1947WSDr5A4OislB2a",
}

// Basic authentication middleware with bcrypt
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		username, password, _ := c.Request.BasicAuth()
		hashedPassword, found := users[username]

		if !found || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
