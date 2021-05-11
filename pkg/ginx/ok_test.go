package ginx

import (
	"backend/pkg/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func testOKJSON(ctx *gin.Context) {
	OK(ctx, []int{1, 2, 3})
}

func TestOK(t *testing.T) {
	e := New()
	e.GET("/testOK", testOKJSON)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/testOK", nil)

	e.ServeHTTP(w, r)

	want := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: 0,
		Msg:  "Success",
		Data: []int{1, 2, 3},
	}

	var ans = struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{}
	if err := json.Unmarshal(w.Body.Bytes(), &ans); err != nil {
		t.Errorf("err: %v", err.Error())
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want.Code, ans.Code)
	assert.Equal(t, want.Msg, ans.Msg)
	v1, ok1 := want.Data.([]int)
	v2, ok2 := ans.Data.([]int)
	if ok1 && ok2 {
		assert.Equal(t, v1, v2)
	}
}

func Authorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sign := ctx.Request.Header.Get("SESSION_ID")
		log.Println(ctx.Request.URL)
		if tools.EmptyString(sign) {
			AuthFailed(ctx)
			ctx.Abort()
		}
		t := time.Now()
		ctx.Next()
		latency := time.Since(t)
		log.Println(latency)
	}
}

func Auth(ctx *gin.Context) {
	time.Sleep(time.Second)
	OK(ctx, nil)
}

func TestAuthFailed(t *testing.T) {
	e := New()
	e.POST("/auth", Authorized(), Auth)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/auth", nil)
	//r.Header.Set("SESSION_ID", "123456")
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}


