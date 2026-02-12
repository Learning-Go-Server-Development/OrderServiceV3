package delegate_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/delegate"
)

func TestServiceDelegate_ProcessCustomerAddresses(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		padds *[]delegate.ProxyAddress
		want  int
		want2 string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			padds: &[]delegate.ProxyAddress{
				{
					ID:      1,
					CID:     11,
					Street:  "123 Peachtree st",
					City:    "Atlanta",
					State:   "GA",
					ZipCode: "12345",
				},
				{
					ID:      2,
					CID:     22,
					Street:  "123 Flagler st",
					City:    "Miami",
					State:   "FL",
					ZipCode: "45678",
				},
			},
			want:  2,
			want2: "123 Peachtree st",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var dd delegate.ServiceDelegate
			d := dd.New()
			got := d.ProcessCustomerAddresses(tt.padds)
			// TODO: update the condition below to compare got with tt.want.
			if len(*got) != tt.want && (*got)[0].State != tt.want2 {
				t.Errorf("ProcessCustomerAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
