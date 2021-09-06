package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS
func TestServeFile(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	router := httprouter.New()
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Hello HTTPRouter", string(body))
}
func TestServeFileGoodBye(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	router := httprouter.New()
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	assert.Equal(t, "Good Bye HTTPRouter", string(body))
}
