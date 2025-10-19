package test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"test/portal/domain"
	routedelivery "test/portal/internal/route/delivery/http"
	"test/portal/internal/route/repository/yaml"
	"test/portal/internal/route/usecase"

	"test/portal/pkg/httputils"
	"test/portal/pkg/validations"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouteTestSuite struct {
	suite.Suite
	repo     domain.RouteItemRepository
	usecase  domain.RouteItemUsecase
	validate *validator.Validate
	ctx      context.Context
	handler  *RouteHandler
}

type RouteHandler struct {
	usecase  domain.RouteItemUsecase
	validate *validator.Validate
}

func (suite *RouteTestSuite) SetupTest() {

	// Clean or recreate yaml file
	os.Remove("./.data/routes.yaml")
	file, err := os.Create("./.data/routes.yaml")
	if err != nil {
		log.Fatal("Failed to create test yaml file:", err)
	}
	file.Close()

	suite.repo = yaml.NewRouteYamlRepository()
	suite.usecase = usecase.NewRouteUsecase(suite.repo)
	suite.validate = validator.New()
	if err := suite.validate.RegisterValidation("is_valid_name", validations.IsValidName); err != nil {
		log.Println("Failed initiate validator is_valid_name", err)
	}
	if err := suite.validate.RegisterValidation("is_valid_path", validations.IsValidPath); err != nil {
		log.Println("Failed initiate validator is_valid_path", err)
	}
	if err := suite.validate.RegisterValidation("is_valid_host", validations.IsValidHostName); err != nil {
		log.Println("Failed initiate validator is_valid_host", err)
	}
	if err := suite.validate.RegisterValidation("is_valid_backend_url", validations.IsValidBackendUrl); err != nil {
		log.Println("Failed initiate validator is_valid_backend_url", err)
	}

	suite.ctx = context.Background()

	suite.handler = &RouteHandler{
		usecase:  suite.usecase,
		validate: suite.validate,
	}
}
func (suite *RouteTestSuite) TestCreateRoute() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// Test data
	isEnabled := true
	route := domain.RouteItem{
		Name:    "test-route",
		Host:    "test.example.com",
		Path:    "/test",
		Backend: "http://localhost:8080",
		Enabled: &isEnabled,
	}

	payload, err := json.Marshal(route)
	assert.NoError(suite.T(), err)

	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodPost,
		Path:    "/routes",
		Payload: bytes.NewBuffer(payload),
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.Create(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Success", responseBody["message"])
	assert.Equal(suite.T(), float64(200), responseBody["status"])
}

func (suite *RouteTestSuite) TestGetAllRoutes() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodGet,
		Path:    "/routes",
		Payload: nil,
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.GetAll(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Success", responseBody["message"])
	assert.Equal(suite.T(), float64(200), responseBody["status"])
}

func (suite *RouteTestSuite) TestGetOneRoute() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// First create a route to get
	isEnabled := true
	route := domain.RouteItem{
		Name:    "get-test-route",
		Host:    "get.example.com",
		Path:    "/get-test",
		Backend: "http://localhost:8081",
		Enabled: &isEnabled,
	}
	_, err := suite.repo.Create(suite.ctx, route)
	assert.NoError(suite.T(), err)

	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodGet,
		Path:    "/routes/get-test-route",
		Payload: nil,
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.GetOne(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Success", responseBody["message"])
	assert.Equal(suite.T(), float64(200), responseBody["status"])
}

func (suite *RouteTestSuite) TestUpdateRoute() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// First create a route to update
	isEnabled := true
	originalRoute := domain.RouteItem{
		Name:    "update-test-route",
		Host:    "update.example.com",
		Path:    "/update-test",
		Backend: "http://localhost:8082",
		Enabled: &isEnabled,
	}
	_, err := suite.repo.Create(suite.ctx, originalRoute)
	assert.NoError(suite.T(), err)

	// Update the route
	isEnabled = false
	updatedRoute := domain.RouteItem{
		Name:    "update-test-route",     // Keep same name
		Host:    "updated.example.com",   // Change host
		Path:    "/update-test",          // Change path
		Backend: "http://localhost:8083", // Change backend
		Enabled: &isEnabled,              // Change enable status
	}

	payload, err := json.Marshal(updatedRoute)
	assert.NoError(suite.T(), err)

	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodPut,
		Path:    "/routes/update-test-route",
		Payload: bytes.NewBuffer(payload),
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.Update(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Success", responseBody["message"])
	assert.Equal(suite.T(), float64(200), responseBody["status"])
}

func (suite *RouteTestSuite) TestDeleteRoute() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// First create a route to delete
	isEnabled := true
	route := domain.RouteItem{
		Name:    "delete-test-route",
		Host:    "delete.example.com",
		Path:    "/delete-test",
		Backend: "http://localhost:8084",
		Enabled: &isEnabled,
	}
	_, err := suite.repo.Create(suite.ctx, route)
	assert.NoError(suite.T(), err)

	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodDelete,
		Path:    "/routes/delete-test-route",
		Payload: nil,
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.Delete(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Success", responseBody["message"])
	assert.Equal(suite.T(), float64(200), responseBody["status"])
}

func (suite *RouteTestSuite) TestCreateRoute_InvalidData() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// Test with invalid data (missing required fields)
	invalidRoute := map[string]interface{}{
		"name": "test", // Too short (min 3)
		// Missing required fields: host, path, backend, enable
	}

	payload, err := json.Marshal(invalidRoute)
	assert.NoError(suite.T(), err)

	// Create HTTP request
	config := httputils.HTTPTestConfig{
		Method:  http.MethodPost,
		Path:    "/routes",
		Payload: bytes.NewBuffer(payload),
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.Create(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions - should return bad request
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
}

func (suite *RouteTestSuite) TestGetOneRoute_NotFound() {
	delivery := routedelivery.NewTestRouteDelivery(suite.ctx, suite.validate, suite.usecase)
	// Create HTTP request for non-existent route
	config := httputils.HTTPTestConfig{
		Method:  http.MethodGet,
		Path:    "/routes/non-existent-route",
		Payload: nil,
		HandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			delivery.GetOne(suite.ctx, w, r)
		}),
	}

	response := httputils.HTTPTestRequest(suite.T(), config)

	// Assertions - should return not found or internal server error
	assert.True(suite.T(), response.Code == http.StatusNotFound || response.Code == http.StatusInternalServerError)
}

func TestRouteTestSuite(t *testing.T) {
	suite.Run(t, new(RouteTestSuite))
}
