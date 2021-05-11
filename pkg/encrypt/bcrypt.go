package encrypt

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	Salt     string
	Password string
}

var BcryptErr = errors.New("密码加密错误")

func (b *Bcrypt) BcryptPassword() (string, error) {
	bcryptBytes, err := bcrypt.GenerateFromPassword([]byte(b.Salt+b.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", BcryptErr
	}
	return string(bcryptBytes), nil
}

func (b *Bcrypt) CheckBcryptPassword(bcryptPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(b.Salt+b.Password)) == nil
}
