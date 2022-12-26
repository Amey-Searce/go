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
		router.HandleFunc("/additemtocart", controller.AddItemsToCart).Methods("GET")
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

			fmt.Println("1. Get all the products")
			fmt.Println("2. Get a particular product")
			fmt.Println("3. Add to cart ")
			fmt.Println("4. Insert item to inventory")
			fmt.Println("5. Update Cart")
			fmt.Println("6. Exit")
			fmt.Println("Enter your choice:")
			fmt.Scanln(&system_console_option)
			if system_console_option == "1" {
				var pg_number int

				fmt.Println("Enter page number:")
				fmt.Scanln(&pg_number)
				console_intface_controller.GetProducts(pg_number)
			}
			if system_console_option == "2" {
				var id string

				fmt.Println("Enter the shop id")
				fmt.Scanln(&id)
				console_intface_controller.GetProduct(id)
			}
			if system_console_option == "3" {
				var name_product string
				var quantity int
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
				console_intface_controller.AddItemtoCart(arr_product)
			}
			if system_console_option == "4" {

				var shop_details model.Product
				var specs_details_map = make(map[string]interface{})
				var key string
				var value string
				var opt string
				var description string
				var final_specs string
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
					fmt.Println("k:", k, "v:", v)
					value_converted := fmt.Sprint(v)
					description += string(k) + ":" + string(value_converted) + ","
				}
				final_specs = "{" + description[0:len(description)-1] + "}"
				fmt.Println(final_specs)
				shop_details.Specs = json.RawMessage(final_specs)
				fmt.Println("Enter the sku of the product")
				fmt.Scanln(&shop_details.Sku)
				fmt.Println("Enter the category of the product")
				fmt.Scanln(&shop_details.Category)
				fmt.Println("Enter the price of the product")
				fmt.Scanln(&shop_details.Price)
				fmt.Println("Enter the productid")
				fmt.Scanln(&shop_details.Productid)
				console_intface_controller.InsertProduct(shop_details)
			}
			if system_console_option == "5" {
				var cart_id string
				var product_id string
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
				console_intface_controller.UpdateCart(update_cart_arr)
			}
			if system_console_option == "6" {
				fmt.Println("Exiting...")
				break
			}
		}

	}
}
