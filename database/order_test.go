package database_test

import (
	"testing"

	gdb "github.com/GolangToolKits/go-mysql"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
)

func TestOrderDB_AddOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true
	mm.MockInsertID1 = 2
	mm.MockInsertSuccess1 = true
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		o     *database.Order
		want  bool
		want2 int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			o: &database.Order{
				CID:         12345,
				OrderNumber: "OD-12555-1",
			},
			want:  true,
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var od database.OrderDB
			od.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := od.New()
			got, got2 := d.AddOrder(tt.o)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("AddOrder() = %v, want %v", got, tt.want)
			}
			if got2 == tt.want2 {
				t.Errorf("AddOrder() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOrderDB_UpdateOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true
	mm.MockUpdateSuccess1 = true
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		o    *database.Order
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			o: &database.Order{

				ID:          12,
				OrderNumber: "OD-12555-1--B",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var od database.OrderDB
			od.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := od.New()
			got := d.UpdateOrder(tt.o)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDB_GetOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true

	mm.MockRow1 = &gdb.DbRow{
		Row: []string{"1", "12345", "OD-12555-1", "2026-02-02 18:32:50", ""},
	}

	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id   int64
		want string // order number
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			id:   12,
			want: "OD-12555-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var od database.OrderDB
			od.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := od.New()
			got := d.GetOrder(tt.id)
			// TODO: update the condition below to compare got with tt.want.
			if (*got).OrderNumber != tt.want {
				t.Errorf("GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDB_GetAllOrders(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true

	mm.MockRows1 = &gdb.DbRows{
		Rows: [][]string{{"1", "12345", "OD-12555-1", "2026-02-02 18:32:50", ""},
			{"2", "12345", "OD-12555-2", "2025-03-01 00:01:14", "2025-03-01 00:01:14"}},
	}

	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cid  int64
		want int // number of orders
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			cid:  12345,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var od database.OrderDB
			od.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := od.New()
			got := d.GetAllOrders(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			if len((*got)) != tt.want {
				t.Errorf("GetAllOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDB_DeleteOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true
	mm.MockDeleteSuccess1 = true
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id   int64
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			id:   12,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var od database.OrderDB
			od.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := od.New()
			got := d.DeleteOrder(tt.id)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("DeleteOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
