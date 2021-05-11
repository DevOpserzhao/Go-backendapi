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

	engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	engine.Use(middleware.Limiter())

	adm := engine.Group("/api/v1")

	{
		adm.POST("/register", uc.Register)
		adm.POST("/login", uc.Login)
		adm.GET("/captcha-id", uc.GenCaptcha)
		adm.GET("/captcha-png", uc.GetCaptcha)
		adm.GET("/captcha-verify", uc.VerifyCaptcha)

	}

	auth := engine.Group("/api/v1").Use(middleware.Auth(tokenFace))
	{
		auth.GET("/password", uc.ModifyPassword)
		auth.POST("/logout", uc.Logout)
		auth.POST("/upload", uc.UploadFile)
		adm.GET("/download", uc.DownloadFile)
	}

	return engine
}
