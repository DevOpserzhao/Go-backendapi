package tests

import (
	"backend/internal/logic"
	mockDomain "backend/internal/mock"
	"backend/internal/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

func TestUserRegisterLogic(t *testing.T) {
	var ExistedUserName = errors.New("此用户名已经存在")
	var ExistedEmail = errors.New("此邮箱已经被使用")
	var ErrMySQLQuery = errors.New("数据库错误")
	tests := []struct {
		Name  string
		Param *types.Register
		Want  error
	}{
		{
			Name: "Success",
			Param: &types.Register{
				UserName:   "God Yao",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "123@qq.com",
			},
			Want: nil,
		},
		{
			Name: "UserName Failed",
			Param: &types.Register{
				UserName:   "Mike",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "456@qq.com",
			},
			Want: ExistedUserName,
		},
		{
			Name: "Email Failed",
			Param: &types.Register{
				UserName:   "Mike",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "789@qq.com",
			},
			Want: ExistedEmail,
		},
		{
			Name: "MySQL Failed",
			Param: &types.Register{
				UserName:   "Mike",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "789@qq.com",
			},
			Want: ErrMySQLQuery,
		},
	}
	ctrl := gomock.NewController(t)
	repo := mockDomain.NewMockUserRepositoryFace(ctrl)

	repo.EXPECT().CheckExistedAccountAndRegister(tests[0].Param).Return(nil)
	repo.EXPECT().CheckExistedAccountAndRegister(tests[1].Param).Return(ExistedUserName)
	repo.EXPECT().CheckExistedAccountAndRegister(tests[2].Param).Return(ExistedEmail)
	repo.EXPECT().CheckExistedAccountAndRegister(tests[3].Param).Return(ErrMySQLQuery)

	lg := logic.NewUserLogic(repo, nil, nil, nil)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := lg.UserRegister(tt.Param)
			t.Log(got)
			t.Log(tt.Want)
			if got != tt.Want {
				t.Errorf("Got: %v, Want %v", got, tt.Want)
			}
		})
	}
}
