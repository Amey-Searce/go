package console_intface_controller

import (
	"database/sql"
	"fmt"
	"go-crud/config"
	"go-crud/logging"
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
func GetProducts(page_int string) []model.GetProductDetails {

	logging.InfoLogger.Println("Inside GetPrdoucts Function")
	var product model.GetProductDetails
	var arrProducts []model.GetProductDetails
	var error_response model.GeneralResponse
	db := config.Connect()
	defer db.Close()

	if len(page_int) == 0 {

		// To prevent CORS errors.
		logging.ErrorLogger.Println("Bad Request. Page is empty.")
		error_response.Message = "Bad Request. Page is empty."
		fmt.Println(error_response.Message)
		return arrProducts
	}
	page_number, _ := strconv.Atoi(page_int)
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

	return arrProducts
}

// // An endpoint which will add items in the database.
func InsertProduct(shop_details model.Product, quantity int) (result string) {

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
	_, err = db.Exec("INSERT INTO inventory(product,quantity,productid) VALUES(?,?,?)", shop_details.Name, quantity, shop_details.Productid)

	if err != nil {
		log.Print(err)
		return
	}
	// response.Status = 200
	// response.Message = "Insert data successfully"
	fmt.Print("Inserted data to database")
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

	// fmt.Println(response)
	return response

}

// // adds multiple items to the cart
func AddItemtoCart(arr_product []model.ShopDetailsReq) model.InventoryAdditionalResponse2 {

	var arrProducts []model.Inventory
	var response model.InventoryAdditionalResponse2
	var product model.Inventory
	var prod model.Product
	var arrProductsEmpty []model.Inventory
	var FinalOrderList []model.Inventory
	var less_quantity_inventory string
	var unit_price_of_product float32
	var productEmpty model.Inventory
	var total_price_details float32
	var not_found_order int
	var not_found_inventory string
	var not_placed_order int

	var records_received int
	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")

	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	// Every new transaction will have a new Cart ID.
	cart_id_returned := GetCartID()
	cart_id := strconv.Itoa(int(cart_id_returned))
	// fmt.Println(cart_id)

	// fmt.Println(arr_product)
	for index := range arr_product {

		rows, err := db1.Query("SELECT COUNT(*) from product where name=?", arr_product[index].Name)

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(&records_received)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		if records_received != 0 {
			rows, err = db1.Query("SELECT id,name,specs,sku,category,price,productid from product where name=?", arr_product[index].Name)

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
			// fmt.Printf("Quantiy from request:%v", quantity_request)
			quantity_int := quantity_request

			if err != nil {
				log.Fatal(err.Error())
			}
			// fmt.Println(arrProducts)
			quantity_from_store := arrProducts[index].Quantity
			// fmt.Printf("done with this also")
			// fmt.Printf("Quantity from store: %v", quantity_from_store)

			if quantity_int <= quantity_from_store {

				total_price := unit_price_of_product * float32(quantity_int)
				// fmt.Printf("total price: %v", total_price)
				// fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
				// fmt.Printf("Quantity from request inside loop :%v", quantity_int)
				net_quantity_remain := quantity_from_store - quantity_int

				_, err := db1.Query("INSERT INTO cart(product,quantity,productid,price,cartid) VALUES(?,?,?,?,?)", arr_product[index].Name, quantity_int, arrProducts[index].Productid, total_price, cart_id)

				if err != nil {
					log.Fatal(err.Error())
				}

				total_price_returned, err := db1.Query("SELECT sum(price) from cart where cartid=?", cart_id)
				// fmt.Printf("Product id %v", arrProducts[index].Productid)
				if err != nil {
					log.Fatal(err.Error())
				}

				for total_price_returned.Next() {
					err = total_price_returned.Scan(&total_price_details)
					if err != nil {
						log.Fatal(err.Error())
					}
					// else {
					// 	// fmt.Printf("Total price %v", total_price_details)
					// }
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

				rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", arr_product[index].Name)

				if err != nil {
					log.Print(err)
				}

				for rows.Next() {
					err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
					if err != nil {
						logging.ErrorLogger.Println(err)
						return response
					} else {
						product.Quantity = quantity_int
						FinalOrderList = append(FinalOrderList, product)
					}
				}
			} else {
				less_quantity_inventory += "  The product with name: " + arr_product[index].Name + " is less in number in the inventory. Cannot place this item. "
				not_placed_order += 1
			}
		} else {

			// fmt.Println("Inhere ")
			not_found_inventory += "  The product with ID: " + arr_product[index].Name + " is not found in the inventory. Cannot place this item. "
			// fmt.Println(not_found_inventory)
			not_found_order += 1
			arrProducts = append(arrProducts, productEmpty)
		}
	}

	response.Status = 200
	response.Message = "Success"
	if not_placed_order == len(arrProducts) {
		response.Data = arrProductsEmpty
	} else {
		response.Data = FinalOrderList
	}
	if not_found_order != 0 {
		response.NotfoundResponse = not_found_inventory
	} else {
		response.NotfoundResponse = "N/A"
	}
	if len(less_quantity_inventory) != 0 {
		response.ShortageResponse = less_quantity_inventory
	} else {
		response.ShortageResponse = "N/A"
	}
	response.CartID = cart_id
	response.Price = total_price_details

	return response

}

func UpdateCart(arr_product model.UpdateCartBody) model.InventoryAdditionalResponse2 {

	var arrProducts []model.Inventory
	var response model.InventoryAdditionalResponse2
	var product model.Inventory
	var FinalOrderList []model.Inventory
	var product_arr model.Product
	var arrProductsEmpty []model.Inventory
	// var prod model.Cart
	var unit_price_of_product float32
	var total_price_details float32
	var quantity int
	var productEmpty model.Inventory
	var not_found_order int
	var not_found_inventory string
	var not_placed_order int
	var less_quantity_inventory string
	var name_from_request string
	var records_returned int
	var cart_exists int
	// var updated_quantity int

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")

	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	rows, err := db1.Query("SELECT COUNT(*) from cart where cartid=?", arr_product.CartId)
	if err != nil {
		logging.ErrorLogger.Println(err)
		log.Fatal(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&cart_exists)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if cart_exists != 0 {
		fmt.Println(arr_product.ProductId[0])
		for index := range arr_product.ProductId {

			fmt.Println("Enter the quanity:")
			fmt.Scanln(&quantity)

			// get the total quantity available for that item.
			rows, err = db1.Query("SELECT COUNT(*) from product where productid=?", arr_product.ProductId[index].ProductIdInput)

			if err != nil {
				log.Print(err)
			}

			for rows.Next() {
				err = rows.Scan(&records_returned)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
			if records_returned != 0 {
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
				name_from_request = arrProducts[index].Product
				// fmt.Printf("done with this also")
				// fmt.Printf("Quantity from store: %v", quantity_from_store)

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
					// fmt.Printf("total price: %v", total_price)
					// fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
					// fmt.Printf("Quantity from request inside loop :%v", quantity)
					net_quantity_remain := quantity_from_store - quantity
					// fmt.Printf("Net quantity: %v", int(net_quantity_remain))

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
					// fmt.Printf("Product id %v", arrProducts[index].Productid)
					if err != nil {
						log.Fatal(err.Error())
					}

					for total_price_returned.Next() {
						err = total_price_returned.Scan(&total_price_details)
						if err != nil {
							log.Fatal(err.Error())
						}
					}
					_, err = db1.Query("UPDATE inventory SET quantity=? where productid=?", net_quantity_remain, arr_product.ProductId[index].ProductIdInput)

					if err != nil {
						log.Fatal(err.Error())
					}

					rows, err = db1.Query("SELECT id,product,quantity,productid from inventory where product=?", name_from_request)

					if err != nil {
						log.Print(err)
					}

					for rows.Next() {
						err = rows.Scan(&product.Id, &product.Product, &product.Quantity, &product.Productid)
						if err != nil {
							logging.ErrorLogger.Println(err)
							return response
						} else {
							product.Quantity = quantity
							FinalOrderList = append(FinalOrderList, product)
						}
					}

				} else {
					less_quantity_inventory += "  The product with name: " + name_from_request + " is less in number in the inventory. Cannot place this item. "
					not_placed_order += 1
				}
			} else {

				fmt.Println("Inhere ")
				not_found_inventory += "  The product with ID: " + arr_product.ProductId[index].ProductIdInput + " is not found in the inventory. Cannot place this item. "
				fmt.Println(not_found_inventory)
				not_found_order += 1
				arrProducts = append(arrProducts, productEmpty)
			}
		}
	}
	response.Status = 200
	response.Message = "Success"
	if not_placed_order == len(arr_product.ProductId) {
		response.Data = arrProductsEmpty
	} else {
		response.Data = FinalOrderList
	}
	if not_found_order != 0 {
		response.NotfoundResponse = not_found_inventory
	} else {
		response.NotfoundResponse = "N/A"
	}
	if len(less_quantity_inventory) != 0 {
		response.ShortageResponse = less_quantity_inventory
	} else {
		response.ShortageResponse = "N/A"
	}
	response.CartID = arr_product.CartId
	response.Price = total_price_details

	// fmt.Println(response)
	return response

}

func DeleteProduct(product_id string) string {

	var response model.ProductDetails
	var product_id_exists int
	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) from product where productid=?", product_id)
	if err != nil {
		logging.ErrorLogger.Println(err)
		log.Print(err)

	}
	for rows.Next() {
		err = rows.Scan(&product_id_exists)
		if err != nil {
			log.Print(err)

		}
	}

	// Delete item details from product table
	if product_id_exists != 0 {
		_, err = db.Query("Delete from  product where productid=?", product_id)

		if err != nil {
			log.Print(err)

		}

		// inserts item details into category table
		_, err = db.Exec("Delete from category where productid=?", product_id)

		if err != nil {
			log.Print(err)
		}

		// inserts item details into inventory table
		_, err = db.Exec("Delete from inventory where productid=?", product_id)

		if err != nil {
			log.Print(err)

		}
		response.Status = 200
		response.Message = "Deleted data successfully"
		return "Deleted the data successfully"
	} else {
		return "Record Not Found"
	}

}

func UpdateProduct(specs string, price float32, product_id string) string {

	var product_id_exists int
	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) from product where productid=?", product_id)
	if err != nil {
		logging.ErrorLogger.Println(err)
		log.Print(err)

	}
	for rows.Next() {
		err = rows.Scan(&product_id_exists)
		if err != nil {
			log.Print(err)
		}
	}

	if product_id_exists != 0 {
		_, err := db.Exec("UPDATE Product Set specs=?,price=? where productid=?", specs, price, product_id)

		if err != nil {
			log.Print(err)
		}

		return "Updated Product"
	} else {
		return "Record Not Found"
	}

}
