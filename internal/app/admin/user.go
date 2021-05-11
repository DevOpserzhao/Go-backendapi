package admin

import (
	"backend/internal/domain"
	"backend/internal/types"
	"backend/pkg/ginx"
	"backend/pkg/tools"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"time"
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

func (uc *UserController) GenCaptcha(ctx *gin.Context) {
	captchaID := captcha.New()
	ginx.OK(ctx, struct {
		CaptchaID string `form:"captcha_id"`
	}{
		CaptchaID: captchaID,
	})
	return
}

func (uc *UserController) GetCaptcha(ctx *gin.Context) {
	var c types.Captcha
	if err := ctx.ShouldBindQuery(&c); err != nil {
		ginx.Fail(ctx, errors.New("验证码id错误"))
		return
	}
	if err := captcha.WriteImage(ctx.Writer, c.CaptchaID, 86, 40); err != nil {
		ginx.Fail(ctx, errors.New("生成验证码错误"))
		return
	}
}

func (uc *UserController) VerifyCaptcha(ctx *gin.Context) {
	var c types.Captcha
	if err := ctx.ShouldBindQuery(&c); err != nil {
		ginx.Fail(ctx, errors.New("captcha or code lost"))
		return
	}
	if !captcha.VerifyString(c.CaptchaID, c.Code) {
		ginx.Fail(ctx, errors.New("验证码错误"))
		return
	}
	ginx.OK(ctx, nil)
	return
}

func (uc *UserController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ginx.Fail(ctx, err)
		return
	}
	if file != nil {
		dst := tools.JoinStrings("../resources/files/", time.Now().Format(tools.TimeFileFormat), filepath.Base(file.Filename))
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ginx.Fail(ctx, err)
			return
		}
		ginx.OK(ctx, nil)
	}
}

func (uc *UserController) DownloadFile(ctx *gin.Context) {
	ctx.Header("Content-Disposition", "attachment; filename="+ctx.Query("f"))
	ctx.File(tools.JoinStrings("../resources/files/", ctx.Query("f")))
	return
}

func (uc *UserController) ModifyPassword(ctx *gin.Context) {
	log.Println(ctx.Get("UID"))
}
