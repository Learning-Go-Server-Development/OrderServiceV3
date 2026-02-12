package handlers

import "net/http"

type Handler interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	GetCurrentOrders(w http.ResponseWriter, r *http.Request)
	GetPastOrders(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)

	AddItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)

	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerAddresses(w http.ResponseWriter, r *http.Request)

	GetProduct(w http.ResponseWriter, r *http.Request)
}
