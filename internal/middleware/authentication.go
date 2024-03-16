package middleware

import (
	"net/http"
	"strings"

	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/internal/utils"
	"github.com/luthfirahman/user-svc/internal/utils/constant"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService config.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponse(constant.MESSAGE_UNAUTHORIZED_USER, nil, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponse(constant.MESSAGE_UNAUTHORIZED_USER, nil, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponse(constant.MESSAGE_UNAUTHORIZED_USER, nil, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		if !token.Valid {
			response := utils.BuildResponse(constant.MESSAGE_UNAUTHORIZED_USER, nil, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildResponse(err.Error(), nil, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
		ctx.Set("token", authHeader)
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
