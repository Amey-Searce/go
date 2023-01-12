package config

import (
	"database/sql"
	"fmt"

	"github.com/qustavo/dotsql"
)

func Alter() string {

	db := Connect()
	failure_flag := 0
	dot, err := dotsql.LoadFromFile("./config/alter_table.sql")
	if err != nil {
		fmt.Println("In alter 1")
		failure_flag = 1
		fmt.Println(err)
	}

	_, err = dot.Exec(db, "alter-product-details")
	if err != nil {
		fmt.Println("In alter 2")
		failure_flag = 1
		fmt.Println(err)
	}

	if failure_flag == 0 {
		return "Success"
	} else {
		return "Issue with one or more alter query."
	}
}
func Insert() string {

	db := Connect()
	dot, err := dotsql.LoadFromFile("C:\\Users\\searce\\Desktop\\go-training\\shopping-cart-api\\go\\shopping-cart\\config\\insert_product_details.sql")
	if err != nil {
		return "Insert failure"
	}

	_, err = dot.Exec(db, "insert-product-details")
	if err != nil {
		err_return := fmt.Sprint(err)
		return err_return
	}
	return "Success"
}

func Create() string {

	db := Connect()
	dot, err := dotsql.LoadFromFile("C:\\Users\\searce\\Desktop\\go-training\\shopping-cart-api\\go\\shopping-cart\\config\\create_schema.sql")
	if err != nil {
		return "Create failure"
	}
	_, err = dot.Exec(db, "create-product-table")
	if err != nil {
		return "Create failure"
	}

	// Run queries
	_, err = dot.Exec(db, "create-product-table")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating thr Product Table if not created.")
	_, err = dot.Exec(db, "create-category-table")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating thr Category Table if not created.")
	_, err = dot.Exec(db, "create-cart-table")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating thr Cart Table if not created.")
	_, err = dot.Exec(db, "create-inventory-table")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating thr Inventory Table if not created.")

	return "Success"
}

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "student"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// Loads queries from file
	return db
}
