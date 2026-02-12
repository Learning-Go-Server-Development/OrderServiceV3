package manager_test

import (
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
)

func TestServiceManager_AddOrder(t *testing.T) {
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

	///////var odb database.OrderDB
	////// odb.DB = m
	var ss manager.ServiceManager
	////////s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	odb.AddOrderID = 1
	odb.AddOrderSuc = true
	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		o    *manager.Order
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			o: &manager.Order{
				OrderNumber: "MOD-123545-4",
				CustomerID:  12345,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			s := ss.New()
			got := s.AddOrder(tt.o)
			// TODO: update the condition below to compare got with tt.want.
			if got.Success != tt.want {
				t.Errorf("AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_UpdateOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// // mm := &gdb.MyDBMock{}
	// // mm.MockTestRow = &gdb.DbRow{
	// // 	Row: []string{"0"},
	// // }
	// // mm.MockConnectSuccess = true
	// // mm.MockInsertID1 = 2
	// // mm.MockInsertSuccess1 = true
	// // m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB
	odb.UpdateOrderSuc = true
	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		o    *manager.Order
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			o: &manager.Order{
				ID:          15,
				OrderNumber: "MOD-123545-55",
				CustomerID:  3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.UpdateOrder(tt.o)
			// TODO: update the condition below to compare got with tt.want.
			if got.Success != true {
				t.Errorf("UpdateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_GetOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// // mm := &gdb.MyDBMock{}
	// // mm.MockTestRow = &gdb.DbRow{
	// // 	Row: []string{"0"},
	// // }
	// // mm.MockConnectSuccess = true
	// // mm.MockInsertID1 = 2
	// // mm.MockInsertSuccess1 = true
	// // m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB

	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id   int64
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			id:   15,
			want: "MOD-123545-55",
		},
		{
			name: "test 2",
			id:   55,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var o database.Order
			o.ID = tt.id
			o.OrderNumber = tt.want
			odb.MockOrder = &o
			s := ss.New()
			got := s.GetOrder(tt.id)
			// TODO: update the condition below to compare got with tt.want.
			if got.OrderNumber != tt.want {
				t.Errorf("GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_GetCurrentOrders(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// // mm := &gdb.MyDBMock{}
	// // mm.MockTestRow = &gdb.DbRow{
	// // 	Row: []string{"0"},
	// // }
	// // mm.MockConnectSuccess = true
	// // mm.MockInsertID1 = 2
	// // mm.MockInsertSuccess1 = true
	// // m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB

	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id        int64
		id2       int64
		cid       int64
		orderNum  string
		orderNum2 string
		want      int
	}{
		// TODO: Add test cases.
		{
			name:      "test 1",
			id:        1,
			id2:       2,
			cid:       12345,
			orderNum:  "MOD-123545-55",
			orderNum2: "",
			want:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var o database.Order
			o.ID = tt.id
			o.OrderNumber = tt.orderNum

			var o2 database.Order
			o2.ID = tt.id2
			o2.OrderNumber = tt.orderNum2
			var morders = []database.Order{o, o2}
			odb.MockOrderList = &morders
			//odb.MockOrder = &o

			s := ss.New()
			got := s.GetCurrentOrders(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			if len(*got) != tt.want {
				t.Errorf("GetCurrentOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_GetPastOrders(t *testing.T) {
	var ss manager.ServiceManager

	//------ live testing-----

	var gpx px.GoProxy
	ss.OrderServiceHost = "http://localhost:3001/rs"

	//-------live testing-------

	//----- mock testing-----

	// 	var w1 http.Response
	// 	w1.Body = io.NopCloser(bytes.NewBufferString(`[{"id": 12345555,"cid": 12345,"oid": "OD-1255878","products": [
	// 12345,258444]}]`))
	// 	var gpx px.MockGoProxy
	// 	gpx.MockDoSuccess1 = true
	// 	gpx.MockRespCode = 200
	// 	gpx.MockResp = &w1

	//--- mock testing-----

	ss.Proxy = &gpx
	ss.OrderServiceHost = "http://localhost:3001/rs"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cid  int64
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			cid:  12345,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.GetPastOrders(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			if len(*got) != tt.want {
				t.Errorf("GetPastOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_DeleteCurrentOrder(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	// // mm := &gdb.MyDBMock{}
	// // mm.MockTestRow = &gdb.DbRow{
	// // 	Row: []string{"0"},
	// // }
	// // mm.MockConnectSuccess = true
	// // mm.MockInsertID1 = 2
	// // mm.MockInsertSuccess1 = true
	// // m := mm.New()
	// m.Connect()

	//var odb database.OrderDB
	//odb.DB = m
	var ss manager.ServiceManager
	//s.DB = odb.New()

	//------ using mocked OrderDB-------
	var odb database.MockOrderDB

	ss.DB = &odb
	//------ using mocked OrderDB-------

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		id   int64
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			id:   11,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			odb.DeleteOrderSuc = tt.want
			s := ss.New()
			got := s.DeleteCurrentOrder(tt.id)
			// TODO: update the condition below to compare got with tt.want.
			if (*got).Success != tt.want {
				t.Errorf("DeleteCurrentOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
