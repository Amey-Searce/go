package console_intface_controller

import (
	"database/sql"
	"fmt"
	"go-crud/config"
	"go-crud/model"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetCartID() int64 {
	// returns a unique Cart ID
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Returns all the products with details such as Name, Product ID, Specs.
func GetProducts(page_number int) []model.GetProductDetails {
	var product model.GetProductDetails
	var arrProducts []model.GetProductDetails
	db := config.Connect()
	defer db.Close()
	// Added pagination in the query.
	// 20 records at max should be displayed
	rows, err := db.Query("SELECT productid, name, specs from product limit 20 offset ?", (page_number-1)*20)

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

	fmt.Println(arrProducts)
	return arrProducts
}

// // An endpoint which will add items in the database.
func InsertProduct(shop_details model.Product) (result string) {

	db := config.Connect()
	defer db.Close()

	// inserts item details into product table
	_, err := db.Exec("INSERT INTO product(name,specs,sku,category,price,productid) VALUES(?,?,?,?,?,?)", shop_details.Name, shop_details.Specs, shop_details.Sku, shop_details.Category, shop_details.Price, shop_details.Productid)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into category table
	_, err = db.Exec("INSERT INTO category(name,productid) VALUES(?,?)", shop_details.Name, shop_details.Productid)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into inventory table
	_, err = db.Exec("INSERT INTO inventory(product,quantity,productid) VALUES(?,?,?)", shop_details.Name, 34, shop_details.Productid)

	if err != nil {
		log.Print(err)
		return
	}
	// response.Status = 200
	// response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")
	return "Inserted Successfully"

}

// Get details of a particular product by the ID.
func GetProduct(id string) model.ProductDetailsConsole {
	var product model.ProductConsole
	var response model.ProductDetailsConsole
	var arrProducts []model.ProductConsole

	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from product where id=?", id)

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

	fmt.Println(response)
	return response

}

// // add an item to the cart and save it to the database.
// func AddItemsToCart(w http.ResponseWriter, r *http.Request) {

// 	var product model.Inventory
// 	var response model.InventoryResponse
// 	var arrProducts []model.Inventory
// 	var prod model.Product
// 	var unit_price_of_product float32
// 	var total_price_details float32

// 	name := r.URL.Query().Get("name")
// 	db := config.Connect()
// 	defer db.Close()
// 	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", name)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&prod.Id, &prod.Name, &prod.Specs, &prod.Sku, &prod.Category, &prod.Price, &prod.Productid)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			unit_price_of_product = prod.Price
// 		}
// 	}

// 	// user sends a request with the product name and quantity
// 	quantity_request := r.URL.Query().Get("quantity")
// 	fmt.Println(name)
// 	fmt.Println(reflect.TypeOf(name))
// 	fmt.Println(quantity_request)
// 	fmt.Println(reflect.TypeOf(quantity_request))
// 	// var exists bool

// 	// Make a call to the inventory table to get the total quantity of that item based on the product name
// 	rows, err = db.Query("SELECT id,product,quantity,productid from inventory where product=?", name)
// 	fmt.Println(rows)

// 	if err != nil {
// 		log.Print(err)
// 	}
// 	fmt.Println("Completed first query")
// 	for rows.Next() {
// 		err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			arrProducts = append(arrProducts, product)
// 		}
// 	}
// 	fmt.Println(arrProducts)
// 	fmt.Println("Completed second query")
// 	quantity_int, err := strconv.Atoi(quantity_request)

// 	fmt.Println(quantity_int)
// 	fmt.Println(reflect.TypeOf(quantity_int))

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	quantity_from_store := arrProducts[0].Quantity

// 	fmt.Println(quantity_from_store)
// 	fmt.Println(reflect.TypeOf(quantity_from_store))

// 	if quantity_int <= quantity_from_store {

// 		// The net quantity remaining after the user adds the item in the cart.
// 		total_price := unit_price_of_product * float32(quantity_int)
// 		fmt.Printf("total price: %v", total_price)
// 		net_quantity_remain := quantity_from_store - quantity_int

// 		// checks if the record is present in the cart or not.
// 		// res, err1 := db.Query("select exists(select * from cart where productid=?)", arrProducts[0].Productid)

// 		// if err1 != nil {
// 		// 	log.Fatal(err.Error())
// 		// }

// 		// for res.Next() {
// 		// 	err = res.Scan(&exists)
// 		// 	if err != nil && err != sql.ErrNoRows {
// 		// 		log.Fatal(err.Error())
// 		// 	} else {
// 		// 		fmt.Println(exists)
// 		// 	}
// 		// }

// 		// if err1 != nil {
// 		// 	log.Fatal(err.Error())
// 		// }
// 		// if the cart doesnt have the shop details, then insert a new record.
// 		// if !exists {
// 		_, err := db.Query("INSERT INTO cart(product,quantity,productid,price) VALUES(?,?,?,?)", arrProducts[0].Product, quantity_int, arrProducts[0].Productid, total_price)

// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		total_price_returned, err := db.Query("SELECT sum(price) from cart")

// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}

// 		for total_price_returned.Next() {
// 			err = total_price_returned.Scan(&total_price_details)
// 			if err != nil {
// 				log.Fatal(err.Error())
// 			} else {
// 				fmt.Printf("Total price %v", total_price_details)
// 			}
// 		}

// 		// } else {
// 		// 	_, err = db.Query("UPDATE cart SET quantity=? where productid=?", quantity_int, arrProducts[0].Productid)
// 		// 	if err != nil {
// 		// 		log.Fatal(err.Error())
// 		// 	}
// 		// }
// 		// update the inventory table with the net amount for that product.
// 		rows, err = db.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arrProducts[0].Productid)

// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		for rows.Next() {
// 			err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
// 			if err != nil {
// 				log.Fatal(err.Error())
// 			} else {
// 				arrProducts = append(arrProducts, product)
// 			}
// 		}
// 	}

// 	response.Status = 200
// 	response.Message = "Success"
// 	response.Data = arrProducts
// 	response.Price = total_price_details

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

// // adds multiple items to the cart
func AddItemtoCart(arr_product []model.ShopDetailsReq) model.InventoryResponse {

	var arrProducts []model.Inventory
	var response model.InventoryResponse
	var product model.Inventory
	var prod model.Product
	var unit_price_of_product float32
	var total_price_details float32
	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")

	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	// Every new transaction will have a new Cart ID.
	cart_id_returned := GetCartID()
	cart_id := strconv.Itoa(int(cart_id_returned))
	fmt.Println(cart_id)

	fmt.Println(arr_product)
	for index := range arr_product {

		rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", arr_product[index].Name)

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
		rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", arr_product[index].Name)

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
		quantity_request := int(arr_product[index].Quantity)
		fmt.Printf("Quantiy from request:%v", quantity_request)
		quantity_int := quantity_request

		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(arrProducts)
		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("done with this also")
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		if quantity_int <= quantity_from_store {

			total_price := unit_price_of_product * float32(quantity_int)
			fmt.Printf("total price: %v", total_price)
			fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
			fmt.Printf("Quantity from request inside loop :%v", quantity_int)
			net_quantity_remain := quantity_from_store - quantity_int
			fmt.Printf("Net quantity: %v", int(net_quantity_remain))
			_, err := db1.Query("INSERT INTO cart(product,quantity,productid,price,cartid) VALUES(?,?,?,?,?)", arr_product[index].Name, quantity_int, arrProducts[index].Productid, total_price, cart_id)

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
	response.CartID = cart_id
	response.Price = total_price_details

	fmt.Println(response)
	return response

}

func UpdateCart(arr_product model.UpdateCartBody) model.InventoryResponse {

	var arrProducts []model.Inventory
	var response model.InventoryResponse
	var product model.Inventory
	var product_arr model.Product
	// var prod model.Cart
	var unit_price_of_product float32
	var total_price_details float32
	var quantity int
	// var updated_quantity int

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")

	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	fmt.Println(arr_product.ProductId[0])
	for index := range arr_product.ProductId {

		fmt.Println("Enter the quanity:")
		fmt.Scanln(&quantity)

		// get the unit price of the product
		// rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where productid=?", arr_product[index].ProductId)

		// if err != nil {
		// 	log.Print(err)
		// }

		// for rows.Next() {
		// 	err = rows.Scan(&product_arr.Id, &product_arr.Name, &product_arr.Specs, &product_arr.Sku, &product_arr.Category, &product_arr.Price, &product_arr.Productid)
		// 	if err != nil {
		// 		log.Fatal(err.Error())
		// 	} else {
		// 		unit_price_of_product = prod.Price
		// 	}
		// }
		// update query to change quanity based on the product and cart id.
		// rows, err = db1.Query("Update cart set quantity=? where productid=? and cartid=?", quanity, arr_product[index].ProductId, arr_product[index].CartId)

		// if err != nil {
		// 	log.Print(err)
		// }

		// for rows.Next() {
		// 	err = rows.Scan(&prod.Id, &prod.Product, &prod.Quantity, &prod.Productid, &prod.Price, &prod.Cartid)
		// 	if err != nil {
		// 		log.Fatal(err.Error())
		// 	} else {
		// 		updated_quantity = prod.Quantity
		// 	}
		// }

		// get the total quantity available for that item.
		rows, err := db1.Query("SELECT id,product,quantity,productid from inventory where productid=?", arr_product.ProductId[index].ProductIdInput)

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

		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("done with this also")
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		// if the quantity requested is feasible to be ordered.
		if int(quantity) <= quantity_from_store {

			// get the unit price of the product
			rows, err := db1.Query("SELECT id,name,specs,sku,category,price,productid from product where productid=?", arr_product.ProductId[index].ProductIdInput)

			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				err = rows.Scan(&product_arr.Id, &product_arr.Name, &product_arr.Specs, &product_arr.Sku, &product_arr.Category, &product_arr.Price, &product_arr.Productid)
				if err != nil {
					log.Fatal(err.Error())
				} else {
					unit_price_of_product = product_arr.Price
				}
			}

			total_price := unit_price_of_product * float32(quantity)
			fmt.Printf("total price: %v", total_price)
			fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
			fmt.Printf("Quantity from request inside loop :%v", quantity)
			net_quantity_remain := quantity_from_store - quantity
			fmt.Printf("Net quantity: %v", int(net_quantity_remain))

			// update query to change quantity based on the product and cart id.
			_, err = db1.Query("Update cart set quantity=? where productid=? and cartid=?", quantity, arr_product.ProductId[index].ProductIdInput, arr_product.CartId)

			if err != nil {
				log.Print(err)
			}

			// update the cart with the updated price.
			_, err = db1.Query("Update cart set price=? where productid=? and cartid=?", total_price, arr_product.ProductId[index].ProductIdInput, arr_product.CartId)

			if err != nil {
				log.Print(err)
			}

			// return the total price for all items belonging to that cart.
			total_price_returned, err := db1.Query("SELECT sum(price) from cart where cartid=?", arr_product.CartId)
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
			_, err = db1.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arr_product.ProductId[index].ProductIdInput)

			if err != nil {
				log.Fatal(err.Error())
			}

		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProducts
	response.CartID = arr_product.CartId
	response.Price = total_price_details

	fmt.Println(response)
	return response

}
