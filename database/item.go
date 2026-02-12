package database

import (
	"log"
	"strconv"
)

func (d *OrderDB) AddItem(i *Item) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if i != nil {
		var a []any
		a = append(a, i.OrderID, i.ProductID)
		log.Println(a)
		suc, id = d.DB.Insert(insertItem, a...)
		log.Println("suc in add item", suc)
		log.Println("id in add item", id)
	}
	return suc, id
}

func (d *OrderDB) UpdateItem(i *Item) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if i != nil {
		var a []any
		a = append(a, i.ProductID, i.ID)
		suc = d.DB.Update(updateItem, a...)
		log.Println("suc in update item", suc)
	}
	return suc
}

func (d *OrderDB) GetItems(oid int64) *[]Item {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Item{}
	var a []any
	a = append(a, oid)
	rows := d.DB.GetList(selectItemList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseItemRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *OrderDB) DeleteItem(iid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, iid)
	rtn := d.DB.Delete(deleteItem, a...)
	return rtn
}

func (d *OrderDB) parseItemRow(foundRow *[]string) *Item {
	var rtn Item
	log.Println("foundRow in item", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		log.Println("id err in get item", err)
		if err == nil {
			oid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			if err == nil {
				pid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
				if err == nil {
					rtn.ID = id
					rtn.OrderID = oid
					rtn.ProductID = pid
				}
			}
		}
	}
	return &rtn
}
