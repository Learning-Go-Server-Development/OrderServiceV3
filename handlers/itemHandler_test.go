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

func TestServiceHandler_AddItem(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	var sec security.OAuth2Security
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
			json:  io.NopCloser(bytes.NewBufferString(`{"orderId":1, "productId": 5}`)),
			suc:   true,
			want:  1,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.AddItemRes = &manager.ResponseID{
				ID:      tt.want,
				Success: tt.want2,
			}
			// aJSON := io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`))

			r, _ := http.NewRequest("POST", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			r.Header.Set("Authorization", authHeader)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.AddItem(w, r)

			var res manager.ResponseID
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_UpdateItem(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	var sec security.OAuth2Security
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
		res   manager.Response
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
			json:  io.NopCloser(bytes.NewBufferString(`{"id":1, "orderId":1, "productId": 5}`)),
			suc:   true,
			want:  1,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.UpdateItemRes = &manager.Response{
				Success: tt.want2,
			}
			// aJSON := io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`))

			r, _ := http.NewRequest("PUT", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			r.Header.Set("Authorization", authHeader)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.UpdateItem(w, r)

			var res manager.Response
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetItems(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	var sec security.OAuth2Security
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
		i     *[]manager.Item
		iid   int64
		len   int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			id:   "1",
			i: &[]manager.Item{
				{
					ID:        3,
					OrderID:   1,
					ProductID: 55,
				},
			},
			iid: 1,
			len: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.MockItems = tt.i
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			r.Header.Set("Authorization", authHeader)
			vars := map[string]string{
				"oid": tt.id,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.GetItems(w, r)

			var res []manager.Item
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || len(res) != tt.len {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_DeleteItem(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler

	//OAuth2 JWT Security---------
	var sec security.OAuth2Security
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
			mm.DeleteItemRes = &manager.Response{
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
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.DeleteItem(w, r)
			var res manager.Response
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}
