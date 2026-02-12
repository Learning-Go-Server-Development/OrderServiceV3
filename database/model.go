package database

import "time"

type Order struct {
	ID          int64     `json:"id"`
	CID         int64     `json:"cid"`
	OrderNumber string    `json:"orderNumber"`
	DateEntered time.Time `json:"dateEntered"`
	DateUpdated time.Time `json:"dateUpdated"`
}

type Item struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"orderId"`
	ProductID int64 `json:"productId"`
}
