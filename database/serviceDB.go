package database

import (
	"fmt"
	"log"
	"strconv"

	gdb "github.com/GolangToolKits/go-mysql"
)

const (
	timeFormat     = "2006-01-02 15:04:05"
	dateOnlyFormat = "2006-01-02"
)

type OrderDB struct {
	DB gdb.Database
}

func (d *OrderDB) testConnection() bool {
	log.Println("in testConnection")
	var rtn = false
	var a []any
	log.Println("d.DB: ", fmt.Sprintln(d.DB))
	rowPtr := d.DB.Test("select count(*) from orders ", a...)
	log.Println("rowPtr", *rowPtr)
	log.Println("after testConnection test", *rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		log.Print("Records found during test ")
		log.Println("Records found during test :", int64Val)
		if err != nil {
			log.Println(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

//---implement this method when all
//---methods of the Database interface
//---have been implemented

func (d *OrderDB) New() Database {
	return d
}
