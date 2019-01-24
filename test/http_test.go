package test

import (
	"encoding/json"
	"gin-gonic/config"
	"gin-gonic/model"
	"gin-gonic/server"
	"gin-gonic/service"
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	//path = fmt.Sprintf("http://localhost:8888%s", path)
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetExample(t *testing.T) {
	// Build our expected body
	body := &model.ExampleMessage{
		Message: "Successfully to query get example",
	}

	// Grab our router
	engine := gin.Default()
	engine.Use(limit.MaxAllowed(config.Get().MaxConnCount))
	appConfig := config.Get()
	server.CreateRoutes(appConfig, engine)
	service.Add(service.IExampleServiceType, service.NewExampleService())

	// Perform a GET request with that handler.
	w := performRequest(engine, "GET", "/example")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["message"]

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body.Message, value)
}