package middleware

import (
	"net/http"
	"strings"

	"bitbucket.org/nsjostrom/machinable/auth"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

// AppUserJwtAuthzMiddleware authorizes the JWT in the Authorization header for application users
func AppUserJwtAuthzMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {

			tokenString := strings.Split(values[0], " ")[1]
			token, err := jwt.Parse(tokenString, auth.TokenLookup)

			if err == nil {
				// token is valid, get claims and perform authorization
				claims := token.Claims.(jwt.MapClaims)

				projects, ok := claims["projects"].(map[string]interface{})
				if !ok {
					respondWithError(http.StatusUnauthorized, "improperly formatted access token", c)
					return
				}

				user, ok := claims["user"].(map[string]interface{})
				if !ok {
					respondWithError(http.StatusUnauthorized, "improperly formatted access token", c)
					return
				}

				userType, ok := user["type"].(string)
				if !ok || userType != "app" {
					respondWithError(http.StatusUnauthorized, "invalid access token", c)
					return
				}

				userIsActive, ok := user["active"].(bool)
				if !ok || !userIsActive {
					respondWithError(http.StatusUnauthorized, "user is not active, please confirm your account", c)
					return
				}

				// inject claims into context
				c.Set("projects", projects)
				c.Set("user_id", user["id"])
				c.Set("username", user["username"])

				c.Next()
				return
			}

			respondWithError(http.StatusUnauthorized, "invalid access token", c)
			return
		}

		respondWithError(http.StatusUnauthorized, "access token required", c)
		return
	}
}

// ValidateRefreshToken validates the refresh token
func ValidateRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {

			tokenString := strings.Split(values[0], " ")[1]
			token, err := jwt.Parse(tokenString, auth.TokenLookup)

			if err == nil {
				// token is valid, validate it's a refresh token
				claims := token.Claims.(jwt.MapClaims)

				sessionID, ok := claims["session_id"].(string)
				if !ok {
					respondWithError(http.StatusUnauthorized, "invalid refresh token", c)
					return
				}

				userID, ok := claims["user_id"].(string)
				if !ok {
					respondWithError(http.StatusUnauthorized, "invalid refresh token", c)
					return
				}

				// inject claims into context
				c.Set("session_id", sessionID)
				c.Set("user_id", userID)

				c.Next()
				return
			}

			respondWithError(http.StatusUnauthorized, "invalid refresh token", c)
			return
		}

		respondWithError(http.StatusUnauthorized, "refresh token required", c)
		return
	}
}