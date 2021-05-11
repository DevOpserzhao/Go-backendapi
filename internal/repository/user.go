package repository

import (
	"backend/internal/domain"
	"backend/internal/pkg/e"
	"backend/internal/types"
	"backend/pkg/db"
	"backend/pkg/encrypt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

var _ domain.UserRepositoryFace = (*UserRepository)(nil)

type UserRepository struct {
	db *db.DataBase
}

func NewUserRepository(db *db.DataBase) *UserRepository {
	return &UserRepository{db: db}
}

func NewUserRepositoryFace(ur *UserRepository) domain.UserRepositoryFace {
	var urf domain.UserRepositoryFace = ur
	return urf
}

var ExistedUserName = errors.New("此用户名已经存在")
var ExistedEmail = errors.New("此邮箱已经被使用")

func (ur *UserRepository) CheckExistedAccountAndRegister(reg *types.Register) error {
	// CheckExistedAccount
	var user domain.User
	err := ur.db.Storage.Where("username = ?", reg.UserName).First(&user).Debug().Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return e.ErrorMySQLQuery
	}
	if err != gorm.ErrRecordNotFound {
		return ExistedUserName
	}
	err = ur.db.Storage.Where("email = ?", reg.Email).First(&user).Debug().Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err != gorm.ErrRecordNotFound {
		return ExistedEmail
	}
	// Register Account
	user.UserName = reg.UserName
	user.Email = reg.Email
	eb := encrypt.Bcrypt{
		Salt:     reg.UserName,
		Password: reg.Password,
	}
	bcryptPassword, _ := eb.BcryptPassword()
	user.Password = bcryptPassword
	if err := ur.db.Storage.Create(&user).Debug().Error; err != nil {
		return e.ErrorMySQLQuery
	}
	return nil
}

var NonExistedUserName = errors.New("此用户名不存在, 请先注册!")
var NonExistedEmail = errors.New("此邮箱不存在, 请先注册!")

func (ur *UserRepository) ChekLogin(login *types.Login) (*domain.User, error) {
	var user domain.User
	if strings.ContainsAny(login.Account, "@") {
		err := ur.db.Storage.Where("email = ?", login.Account).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, e.ErrorMySQLQuery
		}
		if err == gorm.ErrRecordNotFound {
			return nil, NonExistedEmail
		}
		return &user, nil
	}
	err := ur.db.Storage.Where("username = ?", login.Account).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, e.ErrorMySQLQuery
	}
	if err == gorm.ErrRecordNotFound {
		return nil, NonExistedUserName
	}
	return &user, nil
}
