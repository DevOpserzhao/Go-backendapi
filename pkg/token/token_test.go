package token

import (
	"backend/pkg/tools"
	"testing"
)

func TestJsonWebToken_Sign(t *testing.T) {
	config := &JsonWebTokenConfig{
		ExpireTime: 60 * 60,
		Secret:     "123",
		Audience:   "User",
	}
	token := New(config).Sign(tools.IntToString(100))
	t.Logf("Token: %s", token)
}

func TestJsonWebToken_Valid(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJhdWQiOiJVc2VyIiwiZXhwIjoxNjE4ODgzMDMyLCJqdGkiOiIxMDAiLCJpYXQiOjE2MTg4Nzk0MzIsIml" +
		"zcyI6IkdvLWFkbWluIiwibmJmIjoxNjE4ODc5NDMyLCJzdWIiOiJKc29uIFdlYiBUb2tlbiJ9" +
		".xlrA7aoKdh64m6vI4Cnyq51fzEEq3hDDxE0H2nVtAI8"
	config := &JsonWebTokenConfig{
		ExpireTime: 60 * 60,
		Secret:     "123",
		Audience:   "User",
	}
	if _, err := New(config).Valid(token); err != nil {
		t.Error("Token Error:", err.Error())
	}
}

func TestJsonWebToken_Refresh(t *testing.T) {
	config := &JsonWebTokenConfig{
		ExpireTime: 60 * 60,
		Secret:     "123",
		Audience:   "User",
	}
	token := New(config).Refresh(tools.IntToString(100))
	t.Logf("Token: %s", token)
}
