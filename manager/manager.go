package manager

import "github.com/Learning-Go-Server-Development/OrderServiceV3/delegate"

type Manager interface {
	AddOrder(o *Order) *ResponseID
	UpdateOrder(o *Order) *Response
	GetOrder(id int64) *Order
	GetCurrentOrders(cid int64) *[]Order
	GetPastOrders(cid int64) *[]Order
	DeleteCurrentOrder(id int64) *Response

	AddItem(i *Item) *ResponseID
	UpdateItem(i *Item) *Response
	GetItems(oid int64) *[]Item
	DeleteItem(iid int64) *Response

	GetProduct(sku string) *Product

	GetCustomer(phone string) *Customer

	GetCustomerAdresses(cid int64) *[]delegate.Address
}
