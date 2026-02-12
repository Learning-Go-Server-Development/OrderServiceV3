package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mux "github.com/GolangToolKits/grrt"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/handlers"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
)

func TestServiceHandler_GetProduct(t *testing.T) {
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
		sku   string
		c     *manager.Product
		id    int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			code: 200,
			sku:  "8675309",
			c: &manager.Product{
				ID:          1,
				SKU:         "8675309",
				Description: "System76 Computer",
				Price:       2800.59,
			},
			id: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.MockProduct = tt.c
			r, _ := http.NewRequest("GET", "/ffllist", nil)
			vars := map[string]string{
				"sku": tt.sku,
			}
			r = mux.SetURLVars(r, vars)
			// r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			h := hh.New()
			h.GetProduct(w, r)

			var res manager.Product
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.ID != tt.id {
				t.Fail()
			}
		})
	}
}
