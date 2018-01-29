package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMessages(t *testing.T) {
	router := gin.Default()

	messageResource := resource.MessageResource{}
	messageResource.InitRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/messages", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)
}
