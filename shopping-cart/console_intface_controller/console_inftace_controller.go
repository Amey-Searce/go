package console_intface_controller

import (
	"database/sql"
	"fmt"
	"go-crud/config"
	"go-crud/model"
	"log"
	"time"
)

// Returns all the products with details such as Name, Product ID, Specs.
func GetProducts() {
	var product model.GetProductDetails
	var arrProducts []model.GetProductDetails
	var page_number int

	fmt.Println("Enter page number:")
	fmt.Scanln(&page_number)
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
}

// // An endpoint which will add items in the database.
// func InsertProduct(w http.ResponseWriter, r *http.Request) {

// 	var response model.ProductDetails

// 	db := config.Connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}
// 	name := r.FormValue("name")
// 	specs := r.FormValue("specs")
// 	sku := r.FormValue("sku")
// 	category := r.FormValue("category")
// 	price := r.FormValue("price")
// 	productid := r.FormValue("productid")

// 	// inserts item details into product table
// 	_, err = db.Exec("INSERT INTO product(name,specs,sku,category,price,productid) VALUES(?,?,?,?,?,?)", name, specs, sku, category, price, productid)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	// inserts item details into category table
// 	_, err = db.Exec("INSERT INTO category(name,productid) VALUES(?,?)", name, productid)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	// inserts item details into inventory table
// 	_, err = db.Exec("INSERT INTO inventory(product,quantity,productid) VALUES(?,?,?)", name, 34, productid)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	response.Status = 200
// 	response.Message = "Insert data successfully"
// 	fmt.Print("Insert data to database")

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

// // Get details of a particular product by the ID.
// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	var product model.Product
// 	var response model.ProductDetails
// 	var arrProducts []model.Product

// 	id := r.URL.Query().Get("id")
// 	db := config.Connect()
// 	defer db.Close()
// 	rows, err := db.Query("SELECT id,name,specs,sku,category,price,productid from product where id=" + id)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&product.Id, &product.Name, &product.Specs, &product.Sku, &product.Category, &product.Price, &product.Productid)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			arrProducts = append(arrProducts, product)
// 		}
// 	}

// 	response.Status = 200
// 	response.Message = "Success"
// 	response.Data = arrProducts

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

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
func AddItemtoCart() {

	// read the list of items from the user.
	// var arrResponse []model.ProductRequestDetails
	var arrProducts []model.Inventory
	var response model.InventoryResponse
	var product model.Inventory
	var prod model.Product
	var unit_price_of_product float32
	var total_price_details float32
	var productdetailsreq model.ShopDetailsReq
	var quantity int
	var name_product string
	var option string
	var arr_product []model.ShopDetailsReq

	db1, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/student")
	// defer db1.Close()
	if err != nil {
		log.Fatal(err)
	}

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

	for index := range arr_product {

		time.Sleep(1 * time.Second)

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

		quantity_from_store := arrProducts[index].Quantity
		fmt.Printf("Quantity from store: %v", quantity_from_store)

		if quantity_int <= quantity_from_store {

			total_price := unit_price_of_product * float32(quantity_int)
			fmt.Printf("total price: %v", total_price)
			fmt.Printf("Quantity from store inside loop :%v", quantity_from_store)
			fmt.Printf("Quantity from request inside loop :%v", quantity_int)
			net_quantity_remain := quantity_from_store - quantity_int
			fmt.Printf("Net quantity: %v", int(net_quantity_remain))

			_, err := db1.Query("INSERT INTO cart(product,quantity,productid,price) VALUES(?,?,?,?)", arr_product[index].Name, quantity_int, arrProducts[index].Productid, total_price)

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

	fmt.Println(response)

}
