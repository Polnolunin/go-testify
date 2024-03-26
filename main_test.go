package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)
	response := responseRecorder.Result()
	require.NotNil(t, req, "Erorr creating request")
	require.Equal(t, 200, response.StatusCode, "Unexpected status code")
	require.NotEmpty(t, response.Body, "Response body is empty")
}

func TestMainHandlerWhereIsTheWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=Unexist", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)
	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode, "Unexepected status code")
	body, _ := io.ReadAll(response.Body)
	require.Contains(t, string(body), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)
	response := responseRecorder.Result()

	body, _ := io.ReadAll(response.Body)
	list := strings.Split(string(body), ",")
	assert.Len(t, list, totalCount, "Unexpected number of cafes")
}
