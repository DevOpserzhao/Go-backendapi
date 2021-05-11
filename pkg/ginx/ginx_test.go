package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	var e interface{} = New()
	if ans, ok := e.(*gin.Engine); !ok {
		t.Errorf("ans: %T, But want %T", ans, e)
	}
}

func testGetString(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func TestRun2GetString(t *testing.T) {

	e := New()
	e.GET("/testRun", testGetString)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/testRun", nil)

	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
