package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	s.handleHello().ServerHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "hello")
}
