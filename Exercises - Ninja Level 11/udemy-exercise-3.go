package main

import "fmt"

type customErr struct {
	name_error string
}

func (e customErr) Error() string {
	return "This is the error: " + e.name_error
}

func foo(err_value error) error {
	return err_value
}

func main() {

	error_obj := customErr{
		name_error: "AttributeError",
	}
	fmt.Println((error_obj))

}
