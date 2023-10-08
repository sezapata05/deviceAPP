package routehandlerpackage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		rr := httptest.NewRecorder()

		router := gin.New()
		router.Group("/ping")

		NewHandler(&Config{
			R: router,
		})

		request, err := http.NewRequest(http.MethodGet, "/ping", nil)
		assert.NoError(t, err)

		request.SetBasicAuth("dc", "dc-nearshore29")

		router.ServeHTTP(rr, request)

		respBody := "pong"

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.String())
	})
}
