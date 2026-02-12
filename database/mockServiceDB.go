package database

import (
	gdb "github.com/GolangToolKits/go-mysql"
)

type MockOrderDB struct {
	DB gdb.Database

	// Order mocking
	AddOrderID  int64
	AddOrderSuc bool

	UpdateOrderSuc bool

	MockOrder *Order

	MockOrderList *[]Order

	DeleteOrderSuc bool

	//Item mocking
	AddItemID  int64
	AddItemSuc bool

	UpdateItemSuc bool

	MockItemList *[]Item

	DeleteItemSuc bool
}

func (d *MockOrderDB) AddOrder(o *Order) (bool, int64) {
	return d.AddOrderSuc, d.AddOrderID
}

func (d *MockOrderDB) UpdateOrder(o *Order) bool {
	return d.UpdateOrderSuc
}

func (d *MockOrderDB) GetOrder(id int64) *Order {
	return d.MockOrder
}

func (d *MockOrderDB) GetAllOrders(cid int64) *[]Order {
	return d.MockOrderList
}

func (d *MockOrderDB) DeleteOrder(id int64) bool {
	return d.DeleteOrderSuc
}

func (d *MockOrderDB) AddItem(i *Item) (bool, int64) {
	return d.AddItemSuc, d.AddItemID
}

func (d *MockOrderDB) UpdateItem(i *Item) bool {
	return d.UpdateItemSuc
}

func (d *MockOrderDB) GetItems(oid int64) *[]Item {
	return d.MockItemList
}

func (d *MockOrderDB) DeleteItem(iid int64) bool {
	return d.DeleteItemSuc
}

func (d *MockOrderDB) New() Database {
	return d
}
