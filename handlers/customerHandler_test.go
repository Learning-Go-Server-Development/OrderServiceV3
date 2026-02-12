package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mux "github.com/GolangToolKits/grrt"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/delegate"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/handlers"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
)

func TestServiceHandler_GetCustomer(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler
	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want2 bool
		phone string
		c     *manager.Customer
		cid   int64
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			code:  200,
			phone: "867-5309",
			c: &manager.Customer{
				ID:          1,
				FirstName:   "Tommy",
				LastName:    "TuTone",
				PhoneNumber: "867-5309",
			},
			cid: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.MockCustomer = tt.c
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			vars := map[string]string{
				"phone": tt.phone,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.GetCustomer(w, r)

			var res manager.Customer
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.ID != tt.cid {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetCustomerAddresses(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler
	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want2 bool
		cid   string
		o     *[]delegate.Address
		id    int64
		len   int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			cid:  "1",
			o: &[]delegate.Address{{
				ID:      1,
				CID:     5,
				Street:  "125 Robins st",
				City:    "Atlanta",
				State:   "GA",
				ZipCode: "12345",
			},
			},
			id:  1,
			len: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.MockCustomerAddresses = tt.o
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			vars := map[string]string{
				"cid": tt.cid,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.GetCustomerAddresses(w, r)

			var res []delegate.Address
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || len(res) != tt.len {
				t.Fail()
			}
		})
	}
}
