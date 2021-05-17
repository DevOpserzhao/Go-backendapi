package tests

import (
	"backend/internal/app/admin"
	mockDomain "backend/internal/mock"
	"backend/internal/router"
	"backend/internal/types"
	"backend/pkg/ginx"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUserRegister(t *testing.T) {
	tests := []struct {
		Name  string
		Param *types.Register
		Want  ginx.Response
	}{
		{
			Name: "Success",
			Param: &types.Register{
				UserName:   "right",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "931883201@qq.com",
			},
			Want: ginx.Response{
				Code: 0,
				Msg:  "Success",
				Data: nil,
			},
		},
		{
			Name: "Failed",
			Param: &types.Register{
				UserName:   "Mike",
				Password:   "123456",
				ConfirmPwd: "123456",
				Email:      "931883202@qq.com",
			},
			Want: ginx.Response{
				Code: 1,
				Msg:  "此用户名已经存在",
				Data: nil,
			},
		},
	}

	ctrl := gomock.NewController(t)

	logic := mockDomain.NewMockUserLogicFace(ctrl)
	controller := admin.NewUserController(logic)

	// Success
	logic.EXPECT().UserRegister(gomock.Eq(tests[0].Param)).Return(nil)
	// Failed
	logic.EXPECT().UserRegister(gomock.Eq(tests[1].Param)).Return(fmt.Errorf("此用户名已经存在"))

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			r := router.SetUpRouter(controller, nil)
			w := httptest.NewRecorder()
			bs, _ := json.Marshal(tt.Param)
			reader := bytes.NewReader(bs)
			req, _ := http.NewRequest("POST", "/api/v1/register", reader)
			r.ServeHTTP(w, req)

			var resp ginx.Response
			_ = json.Unmarshal(w.Body.Bytes(), &resp)

			t.Log(resp)
			t.Log(tt.Want)

			if resp.Code != tt.Want.Code {
				t.Errorf("Got Code = %d, But Want Code = %d", resp.Code, tt.Want.Code)
			}
			if resp.Data != tt.Want.Data {
				t.Errorf("Got Data = %v, But Want Data = %v", resp.Data, tt.Want.Data)
			}
			if resp.Msg != tt.Want.Msg {
				t.Errorf("Got Msg = %s, But Want Msg = %s", resp.Msg, tt.Want.Msg)
			}
		})
	}
}

func TestUserLogin(t *testing.T) {

	var NonExistedUserName = errors.New("此用户名不存在, 请先注册!")
	var NonExistedEmail = errors.New("此邮箱不存在, 请先注册!")
	var EmptyString = ""
	tests := []struct {
		Name  string
		Param *types.Login
		Want  ginx.Response
	}{
		{
			Name: "UserName Success",
			Param: &types.Login{
				Account:  "Mike",
				Password: "123456",
			},
			Want: ginx.Response{
				Code: 0,
				Msg:  "Success",
				Data: map[string]string{
					"token": "token123",
				},
			},
		},
		{
			Name: "Email Success",
			Param: &types.Login{
				Account:  "123@qq.com",
				Password: "123456",
			},
			Want: ginx.Response{
				Code: 0,
				Msg:  "Success",
				Data: map[string]string{
					"token": "token123",
				},
			},
		},
		{
			Name: "UserName Failed",
			Param: &types.Login{
				Account:  "123@qq.com",
				Password: "123456",
			},
			Want: ginx.Response{
				Code: 1,
				Msg:  "此用户名不存在, 请先注册!",
				Data: map[string]string{},
			},
		},
		{
			Name: "Email Failed",
			Param: &types.Login{
				Account:  "123@qq.com",
				Password: "123456",
			},
			Want: ginx.Response{
				Code: 1,
				Msg:  "此邮箱不存在, 请先注册!",
				Data: map[string]string{},
			},
		},
	}

	ctrl := gomock.NewController(t)

	logic := mockDomain.NewMockUserLogicFace(ctrl)
	controller := admin.NewUserController(logic)

	// Success
	logic.EXPECT().UserLogin(gomock.Eq(tests[0].Param)).Return("token123", nil)
	logic.EXPECT().UserLogin(gomock.Eq(tests[1].Param)).Return("token123", nil)

	// Failed
	logic.EXPECT().UserLogin(gomock.Eq(tests[2].Param)).Return(EmptyString, NonExistedUserName)
	logic.EXPECT().UserLogin(gomock.Eq(tests[3].Param)).Return(EmptyString, NonExistedEmail)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			r := router.SetUpRouter(controller, nil)
			w := httptest.NewRecorder()
			bs, _ := json.Marshal(tt.Param)
			reader := bytes.NewReader(bs)
			req, _ := http.NewRequest("POST", "/api/v1/login", reader)
			r.ServeHTTP(w, req)

			var resp ginx.Response
			_ = json.Unmarshal(w.Body.Bytes(), &resp)

			t.Log(resp)
			t.Log(tt.Want)

			if resp.Code != tt.Want.Code {
				t.Errorf("Got Code = %d, But Want Code = %d", resp.Code, tt.Want.Code)
			}
			if reflect.DeepEqual(resp.Data, tt.Want.Data) {
				t.Errorf("Got Data = %v, But Want Data = %v", resp.Data, tt.Want.Data)
			}
			if resp.Msg != tt.Want.Msg {
				t.Errorf("Got Msg = %s, But Want Msg = %s", resp.Msg, tt.Want.Msg)
			}
		})
	}
}
