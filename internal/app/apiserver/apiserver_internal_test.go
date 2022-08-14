package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleTest(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	s.handleTest().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Ok")
}
