package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud/config"
	"go-crud/model"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Returns all the products with details such as Name, Product ID, Specs.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var product model.GetProductDetails
	var response model.GetProductsDetailsResponse
	var arrProducts []model.GetProductDetails
	page := r.URL.Query().Get("page")
	page_int, _ := strconv.Atoi(page)
	db := config.Connect()
	defer db.Close()
	// page_int = 1
	// Added pagination in the query.
	// 20 records at max should be displayed
	rows, err := db.Query("SELECT productid, name, specs from product limit 20 offset ?", (page_int-1)*20)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&product.Productid, &product.Name, &product.Specs)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrProducts = append(arrProducts, product)
		}
	}

	response.Status = 400
	response.Message = "Success"
	response.Data = arrProducts

	w.Header().Set("Content-Type", "application/json")
	// To prevent CORS errors.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// An endpoint which will add items in the database.
func InsertProduct(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var response model.ProductDetails
	var post model.Product
	// var value_converted string
	// var description string
	db := config.Connect()
	defer db.Close()
	json.Unmarshal(reqBody, &post)
	final_specs := string(post.Specs)
	// for k, v := range post.Specs {
	// 	fmt.Println("k:", k, "v:", v)
	// 	value_converted = fmt.Sprint(v)
	// 	description += string(k) + ":" + string(value_converted) + ","
	// }
	// final_specs := description[0 : len(description)-1]

	// err := r.ParseMultipartForm(4096)
	// if err != nil {
	// 	panic(err)
	// }
	// name := r.FormValue("name")
	// specs := r.FormValue("specs")
	// sku := r.FormValue("sku")
	// category := r.FormValue("category")
	// price := r.FormValue("price")
	// productid := r.FormValue("productid")

	// inserts item details into product table
	_, err := db.Exec("INSERT INTO product(name,specs,sku,category,price,productid) VALUES(?,?,?,?,?,?)", post.Name, final_specs, post.Sku, post.Category, post.Price, post.Productid)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into category table
	_, err = db.Exec("INSERT INTO category(name,productid) VALUES(?,?)", post.Name, post.Productid)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into inventory table
	_, err = db.Exec("INSERT INTO inventory(product,quantity,productid) VALUES(?,?,?)", post.Name, 34, post.Productid)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// Get details of a particular product by the ID.
func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	var response model.ProductDetails
	var arrProducts []model.Product

	id := r.URL.Query().Get("id")
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from product where id=" + id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Specs, &product.Sku, &product.Category, &product.Price, &product.Productid)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrProducts = append(arrProducts, product)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// add an item to the cart and save it to the database.
func AddItemsToCart(w http.ResponseWriter, r *http.Request) {

	var product model.Inventory
	var response model.InventoryResponse
	var arrProducts []model.Inventory
	var prod model.Product
	var unit_price_of_product float32
	var total_price_details float32

	name := r.URL.Query().Get("name")
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", name)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&prod.Id, &prod.Name, &prod.Specs, &prod.Sku, &prod.Category, &prod.Price, &prod.Productid)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			unit_price_of_product = prod.Price
		}
	}

	// user sends a request with the product name and quantity
	quantity_request := r.URL.Query().Get("quantity")
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(quantity_request)
	fmt.Println(reflect.TypeOf(quantity_request))
	// var exists bool

	// Make a call to the inventory table to get the total quantity of that item based on the product name
	rows, err = db.Query("SELECT id,product,quantity,productid from inventory where product=?", name)
	fmt.Println(rows)

	if err != nil {
		log.Print(err)
	}
	fmt.Println("Completed first query")
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrProducts = append(arrProducts, product)
		}
	}
	fmt.Println(arrProducts)
	fmt.Println("Completed second query")
	quantity_int, err := strconv.Atoi(quantity_request)

	fmt.Println(quantity_int)
	fmt.Println(reflect.TypeOf(quantity_int))

	if err != nil {
		log.Fatal(err.Error())
	}

	quantity_from_store := arrProducts[0].Quantity

	fmt.Println(quantity_from_store)
	fmt.Println(reflect.TypeOf(quantity_from_store))

	if quantity_int <= quantity_from_store {

		// The net quantity remaining after the user adds the item in the cart.
		total_price := unit_price_of_product * float32(quantity_int)
		fmt.Printf("total price: %v", total_price)
		net_quantity_remain := quantity_from_store - quantity_int

		// checks if the record is present in the cart or not.
		// res, err1 := db.Query("select exists(select * from cart where productid=?)", arrProducts[0].Productid)

		// if err1 != nil {
		// 	log.Fatal(err.Error())
		// }

		// for res.Next() {
		// 	err = res.Scan(&exists)
		// 	if err != nil && err != sql.ErrNoRows {
		// 		log.Fatal(err.Error())
		// 	} else {
		// 		fmt.Println(exists)
		// 	}
		// }

		// if err1 != nil {
		// 	log.Fatal(err.Error())
		// }
		// if the cart doesnt have the shop details, then insert a new record.
		// if !exists {
		_, err := db.Query("INSERT INTO cart(product,quantity,productid,price) VALUES(?,?,?,?)", arrProducts[0].Product, quantity_int, arrProducts[0].Productid, total_price)

		if err != nil {
			log.Fatal(err.Error())
		}
		total_price_returned, err := db.Query("SELECT sum(price) from cart")

		if err != nil {
			log.Fatal(err.Error())
		}

		for total_price_returned.Next() {
			err = total_price_returned.Scan(&total_price_details)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				fmt.Printf("Total price %v", total_price_details)
			}
		}

		// } else {
		// 	_, err = db.Query("UPDATE cart SET quantity=? where productid=?", quantity_int, arrProducts[0].Productid)
		// 	if err != nil {
		// 		log.Fatal(err.Error())
		// 	}
		// }
		// update the inventory table with the net amount for that product.
		rows, err = db.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arrProducts[0].Productid)

		if err != nil {
			log.Fatal(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				arrProducts = append(arrProducts, product)
			}
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts
	response.Price = total_price_details

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// adds multiple items to the cart
func AddItemtoCart(w http.ResponseWriter, r *http.Request) {

	// read the list of items from the user.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post model.ProductRequestDetails2
	var arrResponse []model.ProductRequestDetails
	var arrProducts []model.Inventory
	var response model.InventoryResponse
	var product model.Inventory
	var prod model.Product
	var unit_price_of_product float32
	var total_price_details float32
	json.Unmarshal(reqBody, &post)

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")
	// defer db1.Close()
	if err != nil {
		log.Fatal(err)
	}

	// add the values from the request in a list.
	for index := range post.Response {
		// fmt.Println(post.Response[index].Product, post.Response[index].Quantity_from_request)
		arrResponse = append(arrResponse, post.Response[index])
	}
	fmt.Println(arrResponse)
	for index := range arrResponse {

		time.Sleep(1 * time.Second)

		rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", arrResponse[index].Product)

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&prod.Id, &prod.Name, &prod.Specs, &prod.Sku, &prod.Category, &prod.Price, &prod.Productid)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				unit_price_of_product = prod.Price
			}
		}
		rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", arrResponse[index].Product)

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				arrProducts = append(arrProducts, product)
			}
		}
		quantity_request := int(arrResponse[index].Quantity_from_request)
		fmt.Printf("Quantiy from request:%v", quantity_request)
		quantity_int := quantity_request

		if err != nil {
			log.Fatal(err.Error())
		}

		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		if quantity_int <= quantity_from_store {

			total_price := unit_price_of_product * float32(quantity_int)
			fmt.Printf("total price: %v", total_price)
			fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
			fmt.Printf("Quantity from request inside loop :%v", quantity_int)
			net_quantity_remain := quantity_from_store - quantity_int
			fmt.Printf("Net quantity: %v", int(net_quantity_remain))

			_, err := db1.Query("INSERT INTO cart(product,quantity,productid,price) VALUES(?,?,?,?)", arrProducts[index].Product, quantity_int, arrProducts[index].Productid, total_price)

			if err != nil {
				log.Fatal(err.Error())
			}

			total_price_returned, err := db1.Query("SELECT sum(price) from cart")
			fmt.Printf("Product id %v", arrProducts[index].Productid)
			if err != nil {
				log.Fatal(err.Error())
			}

			for total_price_returned.Next() {
				err = total_price_returned.Scan(&total_price_details)
				if err != nil {
					log.Fatal(err.Error())
				} else {
					fmt.Printf("Total price %v", total_price_details)
				}
			}
			rows, err = db1.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arrProducts[index].Productid)

			if err != nil {
				log.Fatal(err.Error())
			}

			for rows.Next() {
				err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
				if err != nil {
					log.Fatal(err.Error())
				}

			}

		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts
	response.Price = total_price_details

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
