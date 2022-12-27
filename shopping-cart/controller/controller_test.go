package controller

import (
	"bytes"
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
func TestGetProduct(T *testing.T) {
	request, _ := http.NewRequest("GET", "/getproduct?id=16", nil)
	handler := http.HandlerFunc(GetProduct)
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

func TestAddItemsCart(T *testing.T) {

	var json = []byte(`{
			"data": [
				{
					"name": "Nokia-Phone",
					"quantity":2
				}
			]
		}`)
	request, _ := http.NewRequest("POST", "/additemstocart", bytes.NewBuffer(json))
	handler := http.HandlerFunc(AddItemtoCart)
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

func TestDeleteProduct(T *testing.T) {
	request, _ := http.NewRequest("DELETE", "/deleteitem?productid=laptop-23456", nil)
	handler := http.HandlerFunc(DeleteProduct)
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

func TestUpdateProductDetails(T *testing.T) {
	var json = []byte(`{
		"specs":{"memory":"7GB","manufacture":"India","Color":"Black"},
		 "price":36400
		}`)
	request, _ := http.NewRequest("PUT", "/updateproduct?productid=laptop-123456", bytes.NewBuffer(json))
	handler := http.HandlerFunc(UpdateProductDetails)
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

func TestInsertProduct(T *testing.T) {

	var json = []byte(`{"name":"Transcend-Phone",
		"specs":{"memory":"8GB","manufacture":"India","Color":"Black"},
		"category":"Phone",
		"price":36500,
		"productid":"Phone-345993521",
		"sku":"TRANS-PHN-15-BLK"
	}`)
	request, _ := http.NewRequest("POST", "/insertproduct", bytes.NewBuffer(json))
	handler := http.HandlerFunc(InsertProduct)
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
