package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type validMsgModel struct {
	Name string `validate:"required" msg:"名称必填"`
}

func TestOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	OK(c, gin.H{"key": "value"})

	if w.Code != http.StatusOK {
		t.Errorf("OK status = %d, want 200", w.Code)
	}
	if !strings.Contains(w.Body.String(), "key") {
		t.Errorf("OK body = %q, want contain \"key\"", w.Body.String())
	}
}

func TestError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error(c, "something wrong", http.StatusBadRequest)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Error status = %d, want 400", w.Code)
	}
	if !strings.Contains(w.Body.String(), "something wrong") {
		t.Errorf("Error body = %q, want contain \"something wrong\"", w.Body.String())
	}
	if !c.IsAborted() {
		t.Errorf("Error should abort the context")
	}
}

func TestValidMsg(t *testing.T) {
	v := validator.New()
	err := v.Struct(&validMsgModel{})
	if err == nil {
		t.Fatal("expected validation error")
	}
	if got := ValidMsg(err, &validMsgModel{}); got != "名称必填" {
		t.Errorf("ValidMsg = %q, want %q", got, "名称必填")
	}
}

func TestValidMsgFallback(t *testing.T) {
	if got := ValidMsg(errors.New("plain error"), &validMsgModel{}); got != "plain error" {
		t.Errorf("ValidMsg fallback = %q, want %q", got, "plain error")
	}
}
