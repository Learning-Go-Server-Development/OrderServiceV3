package manager_test

import (
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
)

func TestServiceManager_GetProduct(t *testing.T) {
	var ss manager.ServiceManager

	var gpx px.GoProxy
	ss.Proxy = &gpx
	ss.OrderServiceHost = "http://localhost:3001/rs"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sku  string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			sku:  "123444",
			want: "System76 Laptop",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			s := ss.New()
			got := s.GetProduct(tt.sku)
			// TODO: update the condition below to compare got with tt.want.
			if got.Description != tt.want {
				t.Errorf("GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
