package main

import (
	"encoding/json"
	"fmt"
	"go-crud/console_intface_controller"
	"go-crud/logging"
	"go-crud/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"go-crud/controller"
)

func main() {
	var option string
	var system_console_option string
	fmt.Println("Enter 1 to use System Console. Enter 2 to expose the api endpoints")
	fmt.Scanln(&option)

	// option 2 will expose the below api endpoints
	if option == "2" {
		logging.InfoLogger.Println("The application is starting")
		router := mux.NewRouter()
		router.HandleFunc("/getproducts", controller.GetProducts).Methods("GET")
		router.HandleFunc("/getproduct", controller.GetProduct).Methods("GET")
		router.HandleFunc("/insertproduct", controller.InsertProduct).Methods("POST")
		router.HandleFunc("/additemtocart", controller.AddItemtoCart).Methods("POST")
		router.HandleFunc("/additemstocart", controller.AddItemtoCart).Methods("POST")
		router.HandleFunc("/updatecart", controller.UpdateCart).Methods("PUT")
		router.HandleFunc("/deleteitem", controller.DeleteProduct).Methods("DELETE")
		router.HandleFunc("/updateproduct", controller.UpdateProductDetails).Methods("PUT")
		http.Handle("/", router)
		fmt.Println("The server is listening on port 1234")
		log.Fatal(http.ListenAndServe("127.0.0.1:1234", router))
	}
	// if option == 1, the user enters the system console mode
	if option == "1" {
		for {
			fmt.Println("\n\n")
			fmt.Println("1. Get all the products")
			fmt.Println("2. Get a particular product")
			fmt.Println("3. Add items to cart ")
			fmt.Println("4. Add item to cart ")
			fmt.Println("5. Insert item to inventory")
			fmt.Println("6. Update Cart")
			fmt.Println("7. Delete product")
			fmt.Println("8. Update Product")
			fmt.Println("9. Exit")
			fmt.Println("Enter your choice:")
			fmt.Scanln(&system_console_option)
			if system_console_option == "1" {
				var pg_number string
				var arr_obj []model.GetProductDetails

				fmt.Println("Enter page number:")
				fmt.Scanln(&pg_number)
				fmt.Println("Item Details \n")
				arr_obj = console_intface_controller.GetProducts(pg_number)
				if len(arr_obj) != 0 {
					for index := range arr_obj {
						fmt.Printf("Item name %v \n", arr_obj[index].Name)
						fmt.Printf("Item specs %v \n", arr_obj[index].Specs)
						fmt.Printf("Item Product Id %v \n", arr_obj[index].Productid)
						fmt.Println("\n")
					}
				} else {
					fmt.Println("No data found")
				}
			}
			if system_console_option == "2" {
				var id string
				var arr_obj model.ProductDetailsConsole
				fmt.Println("Enter the shop id")
				fmt.Scanln(&id)
				if len(id) != 0 {
					arr_obj = console_intface_controller.GetProduct(id)
					if len(arr_obj.Data) != 0 {
						fmt.Printf("Product ID: %v \n", arr_obj.Data[0].Productid)
						fmt.Printf("Product Name: %v \n", arr_obj.Data[0].Name)
						fmt.Printf("Product Specs: %v \n", arr_obj.Data[0].Specs)
						fmt.Printf("Product Sku: %v \n", arr_obj.Data[0].Sku)
						fmt.Printf("Product Category: %v \n", arr_obj.Data[0].Category)
						fmt.Printf("Product Price: %v \n", arr_obj.Data[0].Price)
					} else {
						fmt.Println("No data found")
					}
				} else {
					fmt.Println("Id cannot be empty")
				}
			}
			if system_console_option == "3" {
				var name_product string
				var quantity int
				var arr_obj model.InventoryAdditionalResponse2
				var productdetailsreq model.ShopDetailsReq
				var arr_product []model.ShopDetailsReq

				for {
					fmt.Println("Enter the Product Name")
					fmt.Scanln(&name_product)
					fmt.Println("Enter the quantity")
					fmt.Scanln(&quantity)

					productdetailsreq.Name = name_product
					productdetailsreq.Quantity = quantity
					arr_product = append(arr_product, productdetailsreq)

					fmt.Println("If done adding shop items, press 1")
					fmt.Scanln(&option)
					if option == "1" {
						break
					}

				}
				arr_obj = console_intface_controller.AddItemtoCart(arr_product)
				fmt.Printf("Total Price: %v\n", arr_obj.Price)
				fmt.Printf("Cart ID:  %v\n", arr_obj.CartID)
				fmt.Printf("Cart Details: \n %v\n", arr_obj.Data)
				fmt.Printf("Shortage status: %v\n", arr_obj.ShortageResponse)
				fmt.Printf("Not Found status: %v\n", arr_obj.NotfoundResponse)

			}
			if system_console_option == "4" {
				var name_product string
				var quantity int
				var arr_obj model.InventoryAdditionalResponse2
				var productdetailsreq model.ShopDetailsReq
				var arr_product []model.ShopDetailsReq

				fmt.Println("Enter the Product Name")
				fmt.Scanln(&name_product)
				fmt.Println("Enter the quantity")
				fmt.Scanln(&quantity)

				productdetailsreq.Name = name_product
				productdetailsreq.Quantity = quantity
				arr_product = append(arr_product, productdetailsreq)

				arr_obj = console_intface_controller.AddItemtoCart(arr_product)
				fmt.Printf("Total Price: %v\n", arr_obj.Price)
				fmt.Printf("Cart ID:  %v\n", arr_obj.CartID)
				fmt.Printf("Cart Details: \n %v\n", arr_obj.Data)
				fmt.Printf("Shortage status: %v\n", arr_obj.ShortageResponse)
				fmt.Printf("Not Found status: %v\n", arr_obj.NotfoundResponse)

			}
			if system_console_option == "5" {

				var shop_details model.Product
				var specs_details_map = make(map[string]interface{})
				var key string
				var value string
				var opt string
				var description string
				var final_specs string
				var quantity int
				fmt.Println("Enter the name of the product")
				fmt.Scanln(&shop_details.Name)
				for {
					fmt.Println("Enter the specs of the product")
					fmt.Println("Enter the specs(key) :")
					fmt.Scanln(&key)
					fmt.Println("Enter the specs(value) :")
					fmt.Scanln(&value)
					specs_details_map[key] = value
					fmt.Println("To exit, press 1, else press any key to continue adding key/value specs")
					fmt.Scanln(&opt)
					if opt == "1" {
						break
					} else {
						continue
					}
				}
				for k, v := range specs_details_map {
					// fmt.Println("k:", k, "v:", v)
					value_converted := fmt.Sprint(v)
					description += string(k) + ":" + string(value_converted) + ","
				}
				final_specs = "{" + description[0:len(description)-1] + "}"
				// fmt.Println(final_specs)
				shop_details.Specs = json.RawMessage(final_specs)
				fmt.Println("Enter the sku of the product")
				fmt.Scanln(&shop_details.Sku)
				fmt.Println("Enter the category of the product")
				fmt.Scanln(&shop_details.Category)
				fmt.Println("Enter the price of the product")
				fmt.Scanln(&shop_details.Price)
				fmt.Println("Enter the productid")
				fmt.Scanln(&shop_details.Productid)
				fmt.Println("Enter the quantity")
				fmt.Scanln(&quantity)
				console_intface_controller.InsertProduct(shop_details, quantity)
			}
			if system_console_option == "6" {
				var cart_id string
				var product_id string
				var arr_obj model.InventoryAdditionalResponse2
				var update_cart_arr model.UpdateCartBody
				var product_id_input model.UpdateCartBodyConsole
				fmt.Println("Enter the Cart ID")
				fmt.Scanln(&cart_id)
				update_cart_arr.CartId = cart_id
				for {

					fmt.Println("Enter the Product ID")
					fmt.Scanln(&product_id)
					product_id_input.ProductIdInput = product_id
					update_cart_arr.ProductId = append(update_cart_arr.ProductId, product_id_input)

					fmt.Println("If done adding shop items, press 1")
					fmt.Scanln(&option)
					if option == "1" {
						break
					} else {
						continue
					}

				}
				arr_obj = console_intface_controller.UpdateCart(update_cart_arr)
				fmt.Printf("Total Price: %v\n", arr_obj.Price)
				fmt.Printf("Cart ID:  %v\n", arr_obj.CartID)
				fmt.Printf("Cart Details: \n %v\n", arr_obj.Data)
				fmt.Printf("Shortage status: %v\n", arr_obj.ShortageResponse)
				fmt.Printf("Not Found status: %v\n", arr_obj.NotfoundResponse)

			}
			if system_console_option == "7" {
				var product_id string
				var output string
				fmt.Println("Enter productid:")
				fmt.Scanln(&product_id)
				output = console_intface_controller.DeleteProduct(product_id)
				fmt.Println(output)
			}

			if system_console_option == "8" {

				var shop_details model.Product
				var productid string
				var key string
				var value string
				var price float32
				var description string
				var final_specs string
				var opt string
				var output string
				var specs_details_map = make(map[string]interface{})
				fmt.Println("Enter the product id")
				fmt.Scanln(&productid)
				for {
					fmt.Println("Enter the specs of the product")
					fmt.Println("Enter the specs(key) :")
					fmt.Scanln(&key)
					fmt.Println("Enter the specs(value) :")
					fmt.Scanln(&value)
					specs_details_map[key] = value
					fmt.Println("To exit, press 1, else press any key to continue adding key/value specs")
					fmt.Scanln(&opt)
					if opt == "1" {
						break
					} else {
						continue
					}
				}
				for k, v := range specs_details_map {
					// fmt.Println("k:", k, "v:", v)
					value_converted := fmt.Sprint(v)
					description += string(k) + ":" + string(value_converted) + ","
				}
				final_specs = "{" + description[0:len(description)-1] + "}"
				fmt.Println(final_specs)
				shop_details.Specs = json.RawMessage(final_specs)
				fmt.Println("Enter price:")
				fmt.Scanln(&price)
				output = console_intface_controller.UpdateProduct(final_specs, price, productid)
				fmt.Println(output)
			}

			if system_console_option == "9" {
				fmt.Println("Exiting...")
				break
			}
		}

	}
}
