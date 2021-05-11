package admin

import (
	"backend/internal/domain"
	"backend/internal/types"
	"backend/pkg/ginx"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UserController struct {
	logic domain.UserLogicFace
}

func NewUserController(lg domain.UserLogicFace) *UserController {
	return &UserController{logic: lg}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var register types.Register
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ginx.Fail(ctx, err)
		return
	}
	if err := uc.logic.UserRegister(&register); err != nil {
		ginx.Fail(ctx, err)
		return
	}
	ginx.OK(ctx, nil)
	return
}

func (uc *UserController) Login(ctx *gin.Context) {
	var login types.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ginx.Fail(ctx, err)
		return
	}
	token, err := uc.logic.UserLogin(&login)
	if err != nil {
		ginx.Fail(ctx, err)
		return
	}
	ginx.OK(ctx, types.Token{Token: token})
	return
}

func (uc *UserController) Logout(ctx *gin.Context) {
	tokenString, exists := ctx.Get("Token")
	if exists {
		s, ok := tokenString.(string)
		if ok {
			if logout := uc.logic.UserLogout(s); !logout {
				ginx.Fail(ctx, fmt.Errorf("LogOut Failed"))
				return
			}
			ginx.OK(ctx, nil)
			return
		}
		ginx.Fail(ctx, fmt.Errorf("LogOut Failed"))
		return
	}
	ginx.OK(ctx, nil)
	return
}

func (uc *UserController) ModifyPassword(ctx *gin.Context) {
	log.Println(ctx.Get("UID"))
}

func (uc *UserController) UploadFile(ctx *gin.Context) {

}

func (uc *UserController) DownloadFile(ctx *gin.Context) {

}
