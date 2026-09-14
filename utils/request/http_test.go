package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello-body"))
	}))
	defer srv.Close()

	body := HttpRequest(srv.URL)
	if string(body) != "hello-body" {
		t.Errorf("HttpRequest = %q, want %q", string(body), "hello-body")
	}
}
