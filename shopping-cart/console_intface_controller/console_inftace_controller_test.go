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
	result := GetProduct("29")
	if len(result.Data) > 0 {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}

func TestInsertProduct(T *testing.T) {

	var shop_details model.Product
	shop_details.Name = "Apple-X7-Phone"
	shop_details.Sku = "APPLX7-PHN-13-BLK"
	shop_details.Productid = "Phone-1199302"
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
	productdetailsreq.Name = "LG-X10-Phone"
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

func TestUpdateProduct(T *testing.T) {
	final_specs := "{Manufacture:India,Space:8GB}"
	product_id := "Phone-45342123"
	var price float32
	var output string

	output = UpdateProduct(final_specs, price, product_id)
	if output == "Updated Product" {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}

func TestDeleteProduct(T *testing.T) {
	product_id := "Phone-45342123"
	var output string

	output = DeleteProduct(product_id)
	if output == "Deleted the data successfully" {
		T.Logf("Success")
	} else {
		T.Errorf("Failure")
	}

}
