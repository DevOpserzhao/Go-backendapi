package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"time"
)

var _ JsonWebTokenFace = (*JsonWebToken)(nil)

type JsonWebTokenFace interface {
	Sign(string) string
	Valid(string) (*jwt.StandardClaims, error)
	Refresh(string) string
}

type JsonWebTokenConfig struct {
	ExpireTime int64
	Secret     string
	Audience   string
}

type JsonWebToken struct {
	ExpireTime int64
	Secret     string
	Audience   string
}

func New(config *JsonWebTokenConfig) *JsonWebToken {
	return &JsonWebToken{ExpireTime: config.ExpireTime, Secret: config.Secret, Audience: config.Audience}
}

func NewJsonWebTokenFace(token *JsonWebToken) JsonWebTokenFace {
	return JsonWebTokenFace(token)
}

func (j *JsonWebToken) Sign(id string) string {
	now := time.Now().Unix()
	claims := jwt.StandardClaims{
		Audience:  j.Audience,
		ExpiresAt: now + j.ExpireTime,
		Id:        id,
		IssuedAt:  now,
		Issuer:    "Go-backendapi",
		NotBefore: now,
		Subject:   "Json Web Token",
	}
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		log.Println(err.Error())
	}
	return signedString
}

var InvalidToken = errors.New("Invalid Token")

func (j *JsonWebToken) Valid(tokenString string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil || jwtToken == nil {
		return nil, InvalidToken
	}
	if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, InvalidToken
	}
	if err = jwtToken.Claims.Valid(); err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if !ok || !jwtToken.Valid {
		return nil, InvalidToken
	}
	return claims, nil
}

const EmptyString = ""

func (j *JsonWebToken) Refresh(tokenString string) string {
	jwtToken, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Secret), nil
		})
	if err != nil || jwtToken == nil {
		return EmptyString
	}
	if err = jwtToken.Claims.Valid(); err != nil {
		return EmptyString
	}
	if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
		return EmptyString
	}
	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if !ok || !jwtToken.Valid {
		return EmptyString
	}
	return j.Sign(claims.Id)
}
