package database

type Database interface {
	AddOrder(o *Order) (bool, int64)
	UpdateOrder(o *Order) bool
	GetOrder(id int64) *Order
	GetAllOrders(cid int64) *[]Order
	DeleteOrder(id int64) bool

	AddItem(i *Item) (bool, int64)
	UpdateItem(i *Item) bool
	GetItems(oid int64) *[]Item
	DeleteItem(iid int64) bool
}
