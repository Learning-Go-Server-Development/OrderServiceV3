package database

import (
	"log"
	"strconv"
	"time"
)

func (d *OrderDB) AddOrder(o *Order) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if o != nil {
		var a []any
		a = append(a, o.CID, o.OrderNumber, time.Now())
		log.Println(a)
		suc, id = d.DB.Insert(insertOrder, a...)
		log.Println("suc in add order", suc)
		log.Println("id in add order", id)
	}
	return suc, id
}

func (d *OrderDB) UpdateOrder(o *Order) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if o != nil {
		var a []any
		a = append(a, o.OrderNumber, time.Now(), o.ID)
		suc = d.DB.Update(updateOrder, a...)
		log.Println("suc in update order", suc)
	}
	return suc
}

func (d *OrderDB) GetOrder(id int64) *Order {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []any
	a = append(a, id)
	row := d.DB.Get(selectOrderByID, a...)
	rtn := d.parseOrderRow(&row.Row)

	return rtn
}

func (d *OrderDB) GetAllOrders(cid int64) *[]Order {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Order{}
	var a []any
	a = append(a, cid)
	rows := d.DB.GetList(selectOrderList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseOrderRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *OrderDB) DeleteOrder(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)
	rtn := d.DB.Delete(deleteOrder, a...)
	return rtn
}

func (d *OrderDB) parseOrderRow(foundRow *[]string) *Order {
	var rtn Order
	log.Println("foundRow in order", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		log.Println("id err in get order", err)
		if err == nil {
			cid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			if err == nil {
				eTime, _ := time.Parse(timeFormat, (*foundRow)[3])
				uTime, err := time.Parse(timeFormat, (*foundRow)[4])
				if err == nil {
					rtn.DateUpdated = uTime
				}
				rtn.ID = id
				rtn.CID = cid
				rtn.OrderNumber = (*foundRow)[2]
				rtn.DateEntered = eTime
			}
		}
	}
	return &rtn
}
