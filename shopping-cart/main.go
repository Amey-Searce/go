package main

import (
	"fmt"
	"go-crud/console_intface_controller"
	"go-crud/logging"
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

	// if option == 2, then the api endpoints are exposed.
	if option == "2" {
		logging.InfoLogger.Println("The application is starting")
		router := mux.NewRouter()
		router.HandleFunc("/getproducts", controller.GetProducts).Methods("GET")
		router.HandleFunc("/getproduct", controller.GetProduct).Methods("GET")
		router.HandleFunc("/insertproduct", controller.InsertProduct).Methods("POST")
		router.HandleFunc("/additemtocart", controller.AddItemsToCart).Methods("GET")
		router.HandleFunc("/additemstocart", controller.AddItemtoCart).Methods("POST")
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
			fmt.Println("5. Exit")
			fmt.Println("Enter your choice:")
			fmt.Scanln(&system_console_option)
			if system_console_option == "1" {
				console_intface_controller.GetProducts()
			}
			if system_console_option == "3" {
				console_intface_controller.AddItemtoCart()
			}
			if system_console_option == "5" {
				fmt.Println("Exiting...")
				break
			}
		}

	}
}
