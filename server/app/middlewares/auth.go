package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks for a valid JWT token in the authorization header of an incoming request,
// and sets the user ID from the token to the gin context.
//
// It takes a secret key to validate the token.
//
// If the request does not have a valid token, it returns a 401 Unauthorized status. If the token is not valid, it returns
// a 401 Unauthorized status with a JSON response containing the error message. If the token is valid, it sets the
// user ID to the gin context and calls the next handler in the chain.
//
// Example usage:
//
// router.GET("/api/protected", AuthMiddleware("secretKey"), protectedHandler)
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c)
		if err != nil {
			handleAuthError(c, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := validateToken(token, secretKey)
		if err != nil {
			handleAuthError(c, http.StatusUnauthorized, "invalid token")
			return
		}

		userID, err := extractUserID(claims)
		if err != nil {
			handleAuthError(c, http.StatusUnauthorized, "invalid token claims")
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
		return "", errors.New("invalid token format")
	}

	return bearerToken[1], nil
}

func validateToken(tokenStr, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func extractUserID(claims jwt.MapClaims) (uint, error) {
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token claims")
	}
	return uint(userID), nil
}

func handleAuthError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
	c.Abort()
}
