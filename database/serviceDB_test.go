package database

import (
	"testing"

	gdb "github.com/GolangToolKits/go-mysql"
)

func TestOrderDB_testConnection(t *testing.T) {
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
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var d OrderDB
			d.DB = m
			//d.DB.Connect()
			got := d.testConnection()
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("testConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
