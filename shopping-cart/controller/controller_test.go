package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetProducts(T *testing.T) {
	request, _ := http.NewRequest("GET", "/getproducts?page=1", nil)
	handler := http.HandlerFunc(GetProducts)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
	handler.ServeHTTP(response, request)
	status := response.Code
	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}
func TestGetProduct(T *testing.T) {
	request, _ := http.NewRequest("GET", "/getproduct?id=40", nil)
	handler := http.HandlerFunc(GetProduct)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
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
					"name": "Gionee-X34-Phone",
					"quantity":2
				}
			]
		}`)
	request, _ := http.NewRequest("POST", "/additemstocart", bytes.NewBuffer(json))
	handler := http.HandlerFunc(AddItemtoCart)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
	handler.ServeHTTP(response, request)
	status := response.Code
	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}

func TestDeleteProduct(T *testing.T) {
	request, _ := http.NewRequest("DELETE", "/deleteitem?productid=laptop-123456", nil)
	handler := http.HandlerFunc(DeleteProduct)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
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
	request, _ := http.NewRequest("PUT", "/updateproduct?productid=laptop-56789", bytes.NewBuffer(json))
	handler := http.HandlerFunc(UpdateProductDetails)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
	handler.ServeHTTP(response, request)
	status := response.Code
	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}

func TestInsertProduct(T *testing.T) {

	var json = []byte(`{"name":"Samsung-S32-Ultra-Phone",
		"specs":{"memory":"8GB","manufacture":"India","Color":"Black"},
		"category":"Phone",
		"price":36500,
		"productid":"Phone-3454439565490",
		"sku":"SAMULTS32-PHN-15-BLK"
	}`)
	request, _ := http.NewRequest("POST", "/insertproduct", bytes.NewBuffer(json))
	handler := http.HandlerFunc(InsertProduct)
	response := httptest.NewRecorder()
	// fmt.Printf("Request:%v", request)
	// fmt.Printf("Response: %v", response)
	handler.ServeHTTP(response, request)
	status := response.Code
	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}

func TestUpdateCart(T *testing.T) {

	var json = []byte(`
	
		{
		"cartid": "1672349751714",
		"data": [
			{
				"productid": "Phone-9830095542",
				"quantity": 3
			}
			]
		}`)

	request, _ := http.NewRequest("POST", "/updatecart", bytes.NewBuffer(json))
	handler := http.HandlerFunc(UpdateCart)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	status := response.Code

	if status == 200 {
		T.Logf("Success")
	} else {
		T.Errorf("Fail")
	}

}
