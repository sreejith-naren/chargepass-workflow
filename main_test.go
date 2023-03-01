package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestGetAlbums(t *testing.T) {
    r := SetUpRouter()
    r.GET("/albums", GetAlbums)
    req, _ := http.NewRequest("GET", "/albums", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var albums []Album
    json.Unmarshal(w.Body.Bytes(), &albums)

    assert.Equal(t, http.StatusCreated, w.Code)
    assert.NotEmpty(t, albums)
}