package config

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(T *testing.T) {

	response := Connect()
	_, err := response.Query("Select * from cart")
	if err == nil {
		T.Logf("Successfully Connected")
	} else {
		T.Errorf("Db not coneected")
	}
}
