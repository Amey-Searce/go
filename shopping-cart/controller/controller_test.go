package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetProducts(T *testing.T) {
	request, _ := http.NewRequest("GET", "/getproducts?page=1", nil)
	handler := http.HandlerFunc(GetProducts)
	response := httptest.NewRecorder()
	fmt.Printf("Request:%v", request)
	fmt.Printf("Response: %v", response)
	handler.ServeHTTP(response, request)
	status := response.Code
	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}
