package middleware

import (
	"avito-shop-service/internal/config"
	cErrors "avito-shop-service/internal/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthWithJWT(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			_ = ctx.Error(cErrors.ErrAuthorization)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, cErrors.ErrUnexpectedSigningMethod
			}
			return []byte(cfg.JwtSecret), nil
		})

		if err != nil {
			_ = ctx.Error(cErrors.ErrAuthorization)
			return
		}
		if !parsedToken.Valid {
			_ = ctx.Error(cErrors.ErrTokenNotValid)
			return
		}

		parsedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			_ = ctx.Error(cErrors.ErrTokenNotValid)
			return
		}

		ctx.Set("user_id", int(parsedClaims["user_id"].(float64)))
		ctx.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if err == nil {
				return
			}

			var apiErr *cErrors.Error
			if errors.As(err, &apiErr) {
				c.JSON(apiErr.StatusCode, gin.H{"errors": apiErr.Message})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"errors": cErrors.ErrInternalServer.Error()})
		}
	}
}
