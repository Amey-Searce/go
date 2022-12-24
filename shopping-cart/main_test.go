package main

import (
	"fmt"
	"go-crud/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getproducts", controller.GetProducts).Methods("GET")
	return router
}

func TestCreateEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getproducts?page=1", nil)
	response := httptest.NewRecorder()
	fmt.Printf("Request:%v", request)
	fmt.Printf("Response: %v", response)
	// r := Router().HandleFunc("/getproducts", controller.GetProducts).Methods("GET")
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
