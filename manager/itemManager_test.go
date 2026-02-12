package manager_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
)

func TestServiceManager_AddItem(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// mm := &gdb.MyDBMock{}
	// mm.MockTestRow = &gdb.DbRow{
	// 	Row: []string{"0"},
	// }
	// mm.MockConnectSuccess = true
	// mm.MockInsertID1 = 2
	// mm.MockInsertSuccess1 = true
	// m := mm.New()
	//m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	odb.AddItemID = 1
	odb.AddItemSuc = true
	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i    *manager.Item
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			i: &manager.Item{
				OrderID:   15,
				ProductID: 66558,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.AddItem(tt.i)
			// TODO: update the condition below to compare got with tt.want.
			if got.Success != tt.want {
				t.Errorf("AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_UpdateItem(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// mm := &gdb.MyDBMock{}
	// mm.MockTestRow = &gdb.DbRow{
	// 	Row: []string{"0"},
	// }
	// mm.MockConnectSuccess = true
	// mm.MockInsertID1 = 2
	// mm.MockInsertSuccess1 = true
	// m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	odb.UpdateItemSuc = true
	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i    *manager.Item
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			i: &manager.Item{
				ID:        5,
				OrderID:   15,
				ProductID: 474748,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.UpdateItem(tt.i)
			// TODO: update the condition below to compare got with tt.want.
			if got.Success != tt.want {
				t.Errorf("UpdateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_GetItems(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// mm := &gdb.MyDBMock{}
	// mm.MockTestRow = &gdb.DbRow{
	// 	Row: []string{"0"},
	// }
	// mm.MockConnectSuccess = true
	// mm.MockInsertID1 = 2
	// mm.MockInsertSuccess1 = true
	// m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	var i = database.Item{
		ID:        1,
		OrderID:   2,
		ProductID: 3,
	}
	odb.MockItemList = &[]database.Item{i}

	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		oid  int64
		want int
	}{
		// TODO: Add test cases.
		{
			oid:  15,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.GetItems(tt.oid)
			// TODO: update the condition below to compare got with tt.want.
			if len(*got) != tt.want {
				t.Errorf("GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_DeleteItem(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// mm := &gdb.MyDBMock{}
	// mm.MockTestRow = &gdb.DbRow{
	// 	Row: []string{"0"},
	// }
	// mm.MockConnectSuccess = true
	// mm.MockInsertID1 = 2
	// mm.MockInsertSuccess1 = true
	// m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	odb.DeleteItemSuc = true
	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		iid  int64
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			iid:  7,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.DeleteItem(tt.iid)
			// TODO: update the condition below to compare got with tt.want.
			if got.Success != tt.want {
				t.Errorf("DeleteItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
