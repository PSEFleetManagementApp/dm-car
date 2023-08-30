package controller

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

// Creates all resources to mock a running HTTP server
func CreateMockEcho(method string, path string, body io.Reader) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	request := httptest.NewRequest(method, path, body)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	return context, request, recorder
}
