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

func DeleteProduct(product_id string) {

	var response model.ProductDetails
	db := config.Connect()
	defer db.Close()

	// Delete item details from product table
	_, err := db.Query("Delete from  product where productid=?", product_id)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into category table
	_, err = db.Exec("Delete from category where productid=?", product_id)

	if err != nil {
		log.Print(err)
		return
	}

	// inserts item details into inventory table
	_, err = db.Exec("Delete from inventory where productid=?", product_id)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Deleted data successfully"
	fmt.Print("Insert data to database")

}

func UpdateProduct(specs string, price float32, product_id string) {

	db := config.Connect()
	defer db.Close()
	_, err := db.Exec("UPDATE Product Set specs=?,price=? where productid=?", specs, price, product_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Updated Product")

}
