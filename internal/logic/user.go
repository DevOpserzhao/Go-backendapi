package logic

import (
	"backend/internal/domain"
	"backend/internal/pkg/e"
	"backend/internal/types"
	"backend/pkg/cache"
	"backend/pkg/logx"
	"backend/pkg/token"
	"backend/pkg/tools"
)

var _ domain.UserLogicFace = (*UserLogic)(nil)

type UserLogic struct {
	ur    domain.UserRepositoryFace
	cache cache.RedisFace
	token token.JsonWebTokenFace
	log   *logx.ZapLogger
}

func NewUserLogic(ur domain.UserRepositoryFace, cache cache.RedisFace, token token.JsonWebTokenFace,
	log *logx.ZapLogger) *UserLogic {
	return &UserLogic{ur: ur, cache: cache, token: token, log: log}
}

func NewUserLogicFace(ul *UserLogic) domain.UserLogicFace {
	var ulf domain.UserLogicFace = ul
	return ulf
}

func (ul *UserLogic) UserRegister(reg *types.Register) error {
	if err := ul.ur.CheckExistedAccountAndRegister(reg); err != nil {
		return err
	}
	// TODO 异步发送验证邮件
	return nil
}

func (ul *UserLogic) UserLogin(login *types.Login) (string, error) {
	// 用户名和邮箱均可登录
	user, err := ul.ur.ChekLogin(login)
	if err != nil {
		return e.EmptyString, err
	}
	ul.log.Logger.Info(user.UserName)
	// 签发JWT
	return ul.token.Sign(tools.UintToString(user.ID)), nil
}

func (ul *UserLogic) UserLogout(tokenString string) bool {
	if !ul.cache.SIsMember("BlackLogoutList", tokenString) {
		if err := ul.cache.SAdd("BlackLogoutList", tokenString); err != nil {
			return false
		}
		return true
	}
	return true
}
