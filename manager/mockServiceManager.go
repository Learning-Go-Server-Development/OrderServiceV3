package manager

import (
	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/delegate"
)

type MockServiceManager struct {
	DB               database.Database
	Proxy            px.Proxy
	Delegate         delegate.Delegate
	OrderServiceHost string

	AddOrderRes    *ResponseID
	UpdateOrderRes *Response
	MockOrder      *Order
	MockOrders     *[]Order
	DeleteOrderRes *Response

	AddItemRes    *ResponseID
	UpdateItemRes *Response
	MockItems     *[]Item
	DeleteItemRes *Response

	MockCustomer          *Customer
	MockCustomerAddresses *[]delegate.Address

	MockProduct *Product
}

func (s *MockServiceManager) AddOrder(o *Order) *ResponseID {
	return s.AddOrderRes
}

func (s *MockServiceManager) UpdateOrder(o *Order) *Response {
	return s.UpdateOrderRes
}

func (s *MockServiceManager) GetOrder(id int64) *Order {
	return s.MockOrder
}

func (s *MockServiceManager) GetCurrentOrders(cid int64) *[]Order {
	return s.MockOrders
}

func (s *MockServiceManager) GetPastOrders(cid int64) *[]Order {
	return s.MockOrders
}

func (s *MockServiceManager) DeleteCurrentOrder(id int64) *Response {
	return s.DeleteOrderRes
}

func (s *MockServiceManager) AddItem(i *Item) *ResponseID {
	return s.AddItemRes
}

func (s *MockServiceManager) UpdateItem(i *Item) *Response {
	return s.UpdateItemRes
}

func (s *MockServiceManager) GetItems(oid int64) *[]Item {
	return s.MockItems
}

func (s *MockServiceManager) DeleteItem(iid int64) *Response {
	return s.DeleteItemRes
}

func (s *MockServiceManager) GetCustomer(phone string) *Customer {
	return s.MockCustomer
}

func (s *MockServiceManager) GetCustomerAddresses(cid int64) *[]delegate.Address {
	return s.MockCustomerAddresses
}

func (s *MockServiceManager) GetProduct(sku string) *Product {
	return s.MockProduct
}

func (s *MockServiceManager) New() Manager {
	return s
}
