package console_intface_controller

import (
	"encoding/json"
	"fmt"
	"go-crud/model"
	"testing"
)

func TestGetProducts(T *testing.T) {
	result := GetProducts("1")
	if len(result) >= 0 {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}

func TestGetProduct(T *testing.T) {
	result := GetProduct("5")
	if len(result.Data) > 0 {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}

func TestInsertProduct(T *testing.T) {

	var shop_details model.Product
	shop_details.Name = "Motorola-Phone"
	shop_details.Sku = "MOTO-PHN-13-BLK"
	shop_details.Productid = "Phone-10102"
	shop_details.Category = "Phone"
	shop_details.Price = 99000
	shop_details.Specs = json.RawMessage("{size: 13-inch}")

	result := InsertProduct(shop_details, 34)
	print(result)
	if len(result) >= 0 {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}
func TestAddItemtoCart(T *testing.T) {

	var productdetailsreq model.ShopDetailsReq
	var arr_product []model.ShopDetailsReq
	productdetailsreq.Name = "Lenovo-laptop"
	productdetailsreq.Quantity = 3
	arr_product = append(arr_product, productdetailsreq)

	result := AddItemtoCart(arr_product)
	fmt.Println(result)
	if result.Price >= 0 {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}
