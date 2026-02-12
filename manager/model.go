package manager

import "time"

type Order struct {
	ID          int64     `json:"id"`
	OrderNumber string    `json:"orderNumber"`
	Entered     time.Time `json:"dateEntered"`
	Updated     time.Time `json:"dateUpdated"`
	CustomerID  int64     `json:"customerId"`
}

type Item struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"orderId"`
	ProductID int64 `json:"productId"`
}

type Product struct {
	ID          int64   `json:"id"`
	SKU         string  `json:"sku"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Customer struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type Address struct {
	ID      int64  `json:"id"`
	CID     int64  `json:"cid"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

// ResponseID ResponseID
type ResponseID struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
	Code    int64 `json:"code"`
}

// Response Response
type Response struct {
	Success bool  `json:"success"`
	Code    int64 `json:"code"`
}

type ProxyOrder struct {
	ID       int64   `json:"id"`
	CID      int64   `json:"cid"`
	OID      string  `json:"oid"`
	Products []int64 `json:"products"`
}

type ProxyCustomer struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type ProxyAddress struct {
	ID      int64  `json:"id"`
	CID     int64  `json:"cid"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type ProxyProduct struct {
	ID          int64   `json:"id"`
	SKU         string  `json:"sku"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
