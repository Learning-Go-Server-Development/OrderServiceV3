package database_test

import (
	"testing"

	gdb "github.com/GolangToolKits/go-mysql"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
)

func TestOrderDB_AddItem(t *testing.T) {
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
	mm.MockInsertSuccess1 = true
	mm.MockInsertID1 = 8
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i     *database.Item
		want  bool
		want2 int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			i: &database.Item{
				OrderID:   10,
				ProductID: 258444,
			},
			want:  true,
			want2: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var id database.OrderDB
			id.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := id.New()
			got, got2 := d.AddItem(tt.i)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("AddItem() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("AddItem() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOrderDB_UpdateItem(t *testing.T) {
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
		i    *database.Item
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			i: &database.Item{
				ID:        1,
				OrderID:   10,
				ProductID: 4561122,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var id database.OrderDB
			id.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := id.New()
			got := d.UpdateItem(tt.i)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("UpdateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDB_GetItems(t *testing.T) {
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
		Rows: [][]string{{"1", "10", "258444"},
			{"2", "10", "258445"}},
	}
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		oid  int64
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			oid:  10,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var id database.OrderDB
			id.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := id.New()
			got := d.GetItems(tt.oid)
			// TODO: update the condition below to compare got with tt.want.
			if len((*got)) != tt.want {
				t.Errorf("GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDB_DeleteItem(t *testing.T) {
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
		iid  int64
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			iid:  2,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var id database.OrderDB
			id.DB = m
			// -----this line not used yet-----
			// -----bacause we havent implemented all-----
			// -----methods of the Database interface-----
			d := id.New()
			got := d.DeleteItem(tt.iid)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("DeleteItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
