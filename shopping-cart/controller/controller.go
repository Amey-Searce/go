package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud/config"
	"go-crud/logging"
	"go-crud/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetCartID() int64 {
	// returns a unique Cart ID
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Returns all the products with details such as Name, Product ID, Specs.
func GetProducts(w http.ResponseWriter, r *http.Request) {

	logging.InfoLogger.Println("Inside GetPrdoucts Function")
	var product model.GetProductDetails
	var response model.GetProductsDetailsResponse
	var error_response model.GeneralResponse
	var arrProducts []model.GetProductDetails
	page := r.URL.Query().Get("page")
	if len(page) == 0 {

		// To prevent CORS errors.
		logging.ErrorLogger.Println("Bad Request. Page is empty.")
		error_response.Status = 400
		error_response.Message = "Bad Request. Page is empty."
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_response)
		return
	}
	page_int, _ := strconv.Atoi(page)
	db := config.Connect()
	defer db.Close()

	// Added pagination in the query.
	// 20 records at max should be displayed
	rows, err := db.Query("SELECT productid, name, specs from product limit 20 offset ?", (page_int-1)*20)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&product.Productid, &product.Name, &product.Specs)
		if err != nil {
			logging.ErrorLogger.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		} else {
			arrProducts = append(arrProducts, product)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts
	logging.InfoLogger.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// An endpoint which will add items in the database.
func InsertProduct(w http.ResponseWriter, r *http.Request) {

	logging.InfoLogger.Println("Inside Insert Product Function")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var response model.GeneralResponse
	var post model.InsertProductDetails
	db := config.Connect()
	defer db.Close()
	json.Unmarshal(reqBody, &post)
	logging.InfoLogger.Println("Request Body: ", post)
	final_specs := string(post.Specs)

	if len(post.Productid) == 0 {

		logging.ErrorLogger.Println("Bad Request. Product ID can't be empty")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Bad Request. Product ID can't be empty")
		return

	}

	// inserts item details into product table
	logging.InfoLogger.Println("Made query to product table")
	_, err := db.Exec("INSERT INTO product(name,specs,sku,category,price,productid) VALUES(?,?,?,?,?,?)", post.Name, final_specs, post.Sku, post.Category, post.Price, post.Productid)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	// inserts item details into category table
	logging.InfoLogger.Println("Made query to category table")
	_, err = db.Exec("INSERT INTO category(name,productid) VALUES(?,?)", post.Name, post.Productid)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	// inserts item details into inventory table
	logging.InfoLogger.Println("Made query to inventory table")
	_, err = db.Exec("INSERT INTO inventory(product,quantity,productid) VALUES(?,?,?)", post.Name, post.Quantity, post.Productid)

	if err != nil {
		log.Print(err)
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	logging.InfoLogger.Println("Details inserted to the database", response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// Get details of a particular product by the ID.
func GetProduct(w http.ResponseWriter, r *http.Request) {

	logging.InfoLogger.Println("Inside Get Particular Product function")
	var product model.ProductConsole
	var response model.ProductDetailsConsole
	var arrProducts []model.ProductConsole

	id := r.URL.Query().Get("id")
	logging.InfoLogger.Println("ID from request:" + id)
	db := config.Connect()
	defer db.Close()

	if len(id) == 0 {

		logging.ErrorLogger.Println("Bad Request. ID can't be empty")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Bad Request. ID can't be empty")
		return

	}

	// search the record by the shop ID
	logging.InfoLogger.Println("Made a database query to product table")
	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from student.product where id=?", id)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		return
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Specs, &product.Sku, &product.Category, &product.Price, &product.Productid)
		if err != nil {
			log.Fatal(err.Error())
		} else {

			logging.InfoLogger.Println("Details retrieved from database:", product.Id, product.Name, product.Specs, product.Sku, product.Category, product.Price, product.Productid)
			arrProducts = append(arrProducts, product)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logging.InfoLogger.Println("Response:", response)
	json.NewEncoder(w).Encode(response)
}

// adds multiple/single items to the cart
func AddItemtoCart(w http.ResponseWriter, r *http.Request) {

	// read the list of items from the user.
	logging.InfoLogger.Println("Inside Add item/s to cart function")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post model.ProductRequestDetails2
	var arrResponse []model.ProductRequestDetails
	var arrProducts []model.Inventory
	var response model.InventoryAdditionalResponse
	var product model.Inventory
	var arrProductsEmpty []model.Inventory
	var FinalOrderList []model.Inventory
	var prod model.Product
	var unit_price_of_product float32
	var total_price_details float32
	var less_quantity_inventory string
	var not_placed_order int

	json.Unmarshal(reqBody, &post)

	logging.InfoLogger.Println("Request Body:", post)

	cart_id_returned := GetCartID()
	cart_id := strconv.Itoa(int(cart_id_returned))

	logging.InfoLogger.Println("Cart ID:", cart_id)

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")
	// defer db1.Close()
	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		return
	}

	// add the values from the request in a list.
	for index := range post.Response {

		arrResponse = append(arrResponse, post.Response[index])
	}
	fmt.Println(arrResponse)
	for index := range arrResponse {

		logging.InfoLogger.Println("Made a query to product table")
		rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", arrResponse[index].Product)

		if err != nil {
			logging.ErrorLogger.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(500)
			return
		}

		for rows.Next() {
			err = rows.Scan(&prod.Id, &prod.Name, &prod.Specs, &prod.Sku, &prod.Category, &prod.Price, &prod.Productid)
			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
				return
			} else {
				unit_price_of_product = prod.Price
				logging.InfoLogger.Println("Unit Price of the Product:", unit_price_of_product)
			}
		}
		logging.InfoLogger.Println("Query made to invenotry table")
		rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", arrResponse[index].Product)

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
				return
			} else {
				arrProducts = append(arrProducts, product)
			}
		}
		quantity_request := int(arrResponse[index].Quantity_from_request)
		logging.InfoLogger.Printf("Quantiy from request:%v", quantity_request)
		quantity_int := quantity_request

		if err != nil {
			logging.ErrorLogger.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(500)
			return
		}

		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		if quantity_int <= quantity_from_store {

			logging.InfoLogger.Println("Quantity requested is less than the inventory quantity")
			total_price := unit_price_of_product * float32(quantity_int)
			logging.InfoLogger.Printf("total price: %v", total_price)
			logging.InfoLogger.Printf("Quantity from store inside loop :%v", quantity_from_store)
			logging.InfoLogger.Printf("Quantity from request inside loop :%v", quantity_int)
			net_quantity_remain := quantity_from_store - quantity_int
			logging.InfoLogger.Printf("Net quantity: %v", int(net_quantity_remain))

			_, err := db1.Query("INSERT INTO cart(product,quantity,productid,price,cartid) VALUES(?,?,?,?,?)", arrProducts[index].Product, quantity_int, arrProducts[index].Productid, total_price, cart_id)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			total_price_returned, err := db1.Query("SELECT sum(price) from cart where cartid=?", cart_id)
			fmt.Printf("Product id %v", arrProducts[index].Productid)
			if err != nil {
				log.Fatal(err.Error())
			}

			for total_price_returned.Next() {
				err = total_price_returned.Scan(&total_price_details)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
				} else {
					fmt.Printf("Total price %v", total_price_details)
				}
			}
			rows, err = db1.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arrProducts[index].Productid)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			for rows.Next() {
				err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
				}

			}

			rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", arrResponse[index].Product)

			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
					return
				} else {
					product.Quantity = quantity_int
					FinalOrderList = append(FinalOrderList, product)
				}
			}

		} else {
			less_quantity_inventory += "  The product with name: " + arrResponse[index].Product + " is less in number in the inventory. Cannot place this item. "
			not_placed_order += 1
		}
	}

	response.Status = 200
	response.Message = "Success"
	if not_placed_order == len(arrProducts) {
		response.Data = arrProductsEmpty
	} else {
		response.Data = FinalOrderList
	}
	response.Price = total_price_details
	response.CartID = cart_id
	if len(less_quantity_inventory) != 0 {
		response.ShortageResponse = less_quantity_inventory
	} else {
		response.ShortageResponse = "No shortage in the products placed for cart ID: " + cart_id
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

// update the cart based on the product ID and cart ID
func UpdateCart(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var arrProducts []model.Inventory
	var arrProductsEmpty []model.Inventory
	var response model.InventoryAdditionalResponse
	var product model.Inventory
	// var post model.UpdateCartBodyApi
	var FinalOrderList []model.Inventory
	var update_data model.UpdateCartBodyApiData
	var product_arr model.Product
	var unit_price_of_product float32
	var total_price_details float32
	var less_quantity_inventory string
	var not_placed_order int

	json.Unmarshal(reqBody, &update_data)

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
	}
	defer db1.Close()

	for index := range update_data.Data {
		rows, err := db1.Query("SELECT id,product,quantity,productid from inventory where productid=?", update_data.Data[index].ProductId)

		if err != nil {
			logging.ErrorLogger.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(500)
		}

		for rows.Next() {
			err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			} else {
				arrProducts = append(arrProducts, product)
			}
		}

		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("done with this also")
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		// if the quantity requested is feasible to be ordered.
		if int(update_data.Data[index].Quantity) <= quantity_from_store {

			// get the unit price of the product
			rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where productid=?", update_data.Data[index].ProductId)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			for rows.Next() {
				err = rows.Scan(&product_arr.Id, &product_arr.Name, &product_arr.Specs, &product_arr.Sku, &product_arr.Category, &product_arr.Price, &product_arr.Productid)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
				} else {
					unit_price_of_product = product_arr.Price
				}
			}

			total_price := unit_price_of_product * float32(update_data.Data[index].Quantity)
			fmt.Printf("total price: %v", total_price)
			fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
			fmt.Printf("Quantity from request inside loop :%v", update_data.Data[index].Quantity)
			net_quantity_remain := quantity_from_store - update_data.Data[index].Quantity
			fmt.Printf("Net quantity: %v", int(net_quantity_remain))

			// update query to change quantity based on the product and cart id.
			_, err = db1.Query("Update cart set quantity=? where productid=? and cartid=?", update_data.Data[index].Quantity, update_data.Data[index].ProductId, update_data.CartId)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			// update the cart with the updated price.
			_, err = db1.Query("Update cart set price=? where productid=? and cartid=?", total_price, update_data.Data[index].ProductId, update_data.CartId)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			// return the total price for all items belonging to that cart.
			total_price_returned, err := db1.Query("SELECT sum(price) from cart where cartid=?", update_data.CartId)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			for total_price_returned.Next() {
				err = total_price_returned.Scan(&total_price_details)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
				} else {
					fmt.Printf("Total price %v", total_price_details)
				}
			}
			_, err = db1.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, update_data.Data[index].ProductId)

			if err != nil {
				logging.ErrorLogger.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(500)
			}

			rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where productid=?", update_data.Data[index].ProductId)

			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
				if err != nil {
					logging.ErrorLogger.Println(err)
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.WriteHeader(500)
					return
				} else {
					product.Quantity = update_data.Data[index].Quantity
					FinalOrderList = append(FinalOrderList, product)
				}
			}

		} else {
			less_quantity_inventory += "  The product with ID: " + update_data.Data[index].ProductId + " is less in number in the invenotry. Cannot place this item. "
			not_placed_order += 1
		}

	}

	response.Status = 200
	response.Message = "Success"
	if not_placed_order == len(arrProducts) {
		response.Data = arrProductsEmpty
	} else {
		response.Data = FinalOrderList
	}
	if len(less_quantity_inventory) != 0 {
		response.ShortageResponse = less_quantity_inventory
	} else {
		response.ShortageResponse = "No shortage in the products placed for cart ID: " + update_data.CartId
	}
	response.CartID = update_data.CartId
	response.Price = total_price_details
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

// Delete product details from the database
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	logging.InfoLogger.Println("Inside delete item details")
	product_id := r.URL.Query().Get("productid")
	logging.InfoLogger.Println("Product ID", product_id)
	var response model.GeneralResponse
	db := config.Connect()
	defer db.Close()

	if len(product_id) == 0 {

		logging.ErrorLogger.Println("Bad Request. Product ID can't be empty")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Bad Request. Product ID can't be empty")
		return

	}

	// Delete item details from product table
	logging.InfoLogger.Println("Delete query to product table")
	_, err := db.Query("Delete from  product where productid=?", product_id)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		return
	}

	// inserts item details into category table
	logging.InfoLogger.Println("Delete query to category table")
	_, err = db.Exec("Delete from category where productid=?", product_id)

	if err != nil {
		logging.ErrorLogger.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(500)
		return
	}

	// inserts item details into inventory table
	logging.InfoLogger.Println("Delete query to inventory table")
	_, err = db.Exec("Delete from inventory where productid=?", product_id)

	if err != nil {

		return
	}
	response.Status = 200
	response.Message = "Deleted data successfully"
	logging.InfoLogger.Println("Deleted record succesfully")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// Update Details of a product
func UpdateProductDetails(w http.ResponseWriter, r *http.Request) {

	logging.InfoLogger.Println("Inside update item details")
	var response model.GeneralResponse
	var update_data model.UpdateProduct
	reqBody, _ := ioutil.ReadAll(r.Body)
	product_id := r.URL.Query().Get("productid")
	logging.InfoLogger.Println("Product ID", product_id)
	db := config.Connect()
	defer db.Close()
	json.Unmarshal(reqBody, &update_data)

	if len(product_id) == 0 {

		logging.ErrorLogger.Println("Bad Request. Product ID can't be empty")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Bad Request. Product ID can't be empty")
		return

	}

	logging.InfoLogger.Println("Update Query to Product table")
	_, err := db.Query("UPDATE product SET specs=?, price=? where productid=?", update_data.Specs, update_data.Price, product_id)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	response.Status = 200
	response.Message = "Updated Succesfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logging.InfoLogger.Println("Updated Successfully")
	json.NewEncoder(w).Encode(response)

}
