package model

import "encoding/json"

// type Product struct {
// 	Id        int     `form:"id" json:"id"`
// 	Name      string  `form:"name" json:"name"`
// 	Sku       string  `form:"sku" json:"sku"`
// 	Specs     string  `form:"specs" json:"specs"`
// 	Category  string  `form:"category" json:"category"`
// 	Price     float32 `form:"price" json:"price"`
// 	Productid string  `form:"productid" json:"productid"`
// }
// Specs     map[string]interface{} `json:"specs"`
type Product struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Sku       string          `json:"sku"`
	Specs     json.RawMessage `json:"specs"`
	Category  string          `json:"category"`
	Price     float32         `json:"price"`
	Productid string          `json:"productid"`
}
type ProductConsole struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Sku       string  `json:"sku"`
	Specs     string  `json:"specs"`
	Category  string  `json:"category"`
	Price     float32 `json:"price"`
	Productid string  `json:"productid"`
}

type GetProductDetails struct {
	Name      string
	Specs     string
	Productid string
}

type Category struct {
	Id        int
	Name      string
	Productid string
}

type Inventory struct {
	Id        int
	Product   string
	Quantity  int
	Productid string
}

type Cart struct {
	Id        int
	Product   string
	Quantity  int
	Productid string
	Price     float32
	Cartid    string
}

type ProductDetails struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Product
}
type GeneralResponse struct {
	Status  int
	Message string
}
type ProductDetailsConsole struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []ProductConsole
}

type ProductRequestDetails struct {
	Quantity_from_request int    `json:"quantity"`
	Product               string `json:"name"`
}
type ProductRequestDetails2 struct {
	Response []ProductRequestDetails `json:"data"`
}

type GetProductsDetailsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []GetProductDetails
}

type InventoryResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Price   float32 `json:"price"`
	CartID  string  `json:"cartid"`
	Data    []Inventory
}
type InventoryAdditionalResponse struct {
	Status           int     `json:"status"`
	Message          string  `json:"message"`
	Price            float32 `json:"price"`
	CartID           string  `json:"cartid"`
	Data             []Inventory
	ShortageResponse string
}

type ShopDetailsReq struct {
	Quantity int
	Name     string
}

type UpdateCartBody struct {
	ProductId []UpdateCartBodyConsole
	CartId    string
}

type UpdateCartBodyConsole struct {
	ProductIdInput string `json:"productid"`
}

type UpdateCartBodyApiData struct {
	CartId string              `json:"cartid"`
	Data   []UpdateCartBodyApi `json:"data"`
}
type UpdateCartBodyApi struct {
	ProductId string `json:"productid"`
	Quantity  int    `json:"quantity"`
}
type UpdateProduct struct {
	Specs json.RawMessage `json:"specs"`
	Price float32         `json:"price"`
}

type InsertProductDetails struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Sku       string          `json:"sku"`
	Specs     json.RawMessage `json:"specs"`
	Category  string          `json:"category"`
	Price     float32         `json:"price"`
	Productid string          `json:"productid"`
	Quantity  int             `json:"quantity"`
}
