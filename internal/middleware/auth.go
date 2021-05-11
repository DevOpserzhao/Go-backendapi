package middleware

import (
	"backend/pkg/ginx"
	"backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func Auth(tokenFace token.JsonWebTokenFace) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if len(tokenString) == 0 {
			ginx.AuthFailed(ctx)
			ctx.Abort()
			return
		}
		claims, err := tokenFace.Valid(tokenString)
		if err != nil || claims == nil {
			ginx.AuthFailed(ctx)
			ctx.Abort()
			return
		}
		ctx.Set("UID", claims.Id)
		ctx.Set("Token", tokenString)
		ctx.Next()
	}
}
