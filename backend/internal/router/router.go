package router

import (
	"backend/internal/app/admin"
	"backend/internal/middleware"
	"backend/pkg/ginx"
	"backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(uc *admin.UserController, tokenFace token.JsonWebTokenFace) *gin.Engine {
	engine := ginx.New()

	engine.Use(middleware.Limiter())

	adm := engine.Group("/api/v1")
	{
		adm.POST("/register", uc.Register)
		adm.POST("/login", uc.Login)
	}
	auth := engine.Group("/api/v1").Use(middleware.Auth(tokenFace))
	{
		auth.GET("/password", uc.ModifyPassword)
		auth.POST("/logout", uc.Logout)
	}
	return engine
}
