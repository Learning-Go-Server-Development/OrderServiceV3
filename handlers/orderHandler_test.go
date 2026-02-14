package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	mux "github.com/GolangToolKits/grrt"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/handlers"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/security"
)

func TestServiceHandler_AddOrder(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		res   manager.ResponseID
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want  int64
		want2 bool
	}{
		// TODO: Add test cases.
		{
			code:  200,
			ctype: "application/json",
			json:  io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`)),
			suc:   true,
			want:  1,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.AddOrderRes = &manager.ResponseID{
				ID:      tt.want,
				Success: tt.want2,
			}
			// aJSON := io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`))

			r, _ := http.NewRequest("POST", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			r.Header.Set("Authorization", authHeader)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.

			h := hh.New()
			h.AddOrder(w, r)
			var res manager.ResponseID
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_UpdateOrder(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		res   manager.ResponseID
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want2 bool
	}{
		// TODO: Add test cases.
		{
			code:  200,
			ctype: "application/json",
			json:  io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`)),
			suc:   true,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.UpdateOrderRes = &manager.Response{
				Success: tt.want2,
			}
			r, _ := http.NewRequest("POST", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			r.Header.Set("Authorization", authHeader)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.

			// TODO: construct the receiver type.

			h := hh.New()
			h.UpdateOrder(w, r)
			var res manager.Response
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetOrder(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

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
		id    string
		o     *manager.Order
		oid   int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			id:   "1",
			o: &manager.Order{
				ID:          1,
				OrderNumber: "OD-12345",
				CustomerID:  55,
			},
			oid: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			// var h handlers.ServiceHandler
			mm.MockOrder = tt.o
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			r.Header.Set("Authorization", authHeader)
			vars := map[string]string{
				"id": tt.id,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()

			h := hh.New()
			h.GetOrder(w, r)

			var res manager.Order
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.ID != tt.oid {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetCurrentOrders(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		json  io.ReadCloser
		code  int
		suc   bool
		want2 bool
		cid   string
		os    *[]manager.Order
		id    int64
		len   int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			cid:  "1",
			os: &[]manager.Order{
				{
					ID:          1,
					OrderNumber: "OD-12345",
					CustomerID:  55,
				},
			},
			id:  11,
			len: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			// var h handlers.ServiceHandler
			mm.MockOrders = tt.os
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			r.Header.Set("Authorization", authHeader)
			vars := map[string]string{
				"cid": tt.cid,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()

			h := hh.New()
			h.GetCurrentOrders(w, r)

			var res []manager.Order
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || len(res) != tt.len {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetPastOrders(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		json  io.ReadCloser
		code  int
		suc   bool
		want2 bool
		cid   string
		os    *[]manager.Order
		id    int64
		len   int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			cid:  "1",
			os: &[]manager.Order{
				{
					ID:          1,
					OrderNumber: "OD-12345",
					CustomerID:  55,
				},
			},
			id:  11,
			len: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			mm.MockOrders = tt.os
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			r.Header.Set("Authorization", authHeader)
			vars := map[string]string{
				"cid": tt.cid,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()

			h := hh.New()
			h.GetPastOrders(w, r)
			var res []manager.Order
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || len(res) != tt.len {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_DeleteOrder(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	// var sec security.OAuth2Security
	var sec security.MockOAuth2Security
	sec.MockValid = true
	sec.ValidationHost = "http://www.goauth2.com"
	sec.Proxy = &px.GoProxy{}
	sec.ClientID = 52
	hh.Security = sec.New()
	//-----------------------------

	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		code  int
		suc   bool
		want2 bool
		id    string
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			code:  200,
			id:    "1",
			suc:   true,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.DeleteOrderRes = &manager.Response{
				Success: tt.want2,
			}
			r, _ := http.NewRequest("DELETE", "/ffllist", nil)
			r.Header.Set("Authorization", authHeader)
			vars := map[string]string{
				"id": tt.id,
			}
			r = mux.SetURLVars(r, vars)
			//r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.DeleteOrder(w, r)
			var res manager.Response
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}
