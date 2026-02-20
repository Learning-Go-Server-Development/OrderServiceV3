package handlers_test

import (
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

var authHeader = "Bearer 1eNq8lU+TqygUxb/RK/+89BuXia0ESs34DwmbqQDpDoom01ZU/PRTZNJTvZiuZPHK/eVwzu/eC0eNTgxwuZMIljO0Ewl72GUr7sMX2FwI9pH346hRuieqZ07YwPryC7Zq5toUiCEl2Dq8Qhn5aBYVNEI1A0qxLnvbu0hRgK8CqBPbxqbmJKrU1GyZFdoCvOuoTQaWwx62+Ke5NC64Gxf7KZ7huPNHeahCC9bnKZn5lNTrVVKvf77loxTA+6DVqjySjYL1We6rpKYkmaEc5d6ZLrRaWbm5qy515KOOt+FVmFBt0h8qfBWhV/MWN6buUIU91yY4GhjAuGzUDOsLO2r0es+SM2dSeWn0AjvOR1m2Xk9z+AI7PFNyy/9Zu2FgGoSDG4wRup9ZxUXcw+Z+dxD2DHgurTJcNo3cSTjwLh4OJFMs92wOspMA5VdNlPnwJSmQF217md7z3TyUKrjxKYIeKuO/NHxt7mBtWHzWptXUM1coGuI8v5/ZFYH1nW/DNHLRHDnKopU9f9X6N1Op3/Lp13+Zwhu7m8+4KKfIRzdmRkeQRHEzWP+bX93PNGPynV/DWXqaa2/gbaa49E4UpEvwsZ/jYw0U4J6STC3CST/NaclZch6zSmzuZgOrsI4coYT/NetNx447q/+te/f6/nDvmIsaSuBAHWwt0L/pif7VgiQWc+z7rGfavG2UxEt5HHf5kzPWxcPexZq3+LpIP4vH/eRbNNAtrpfxs35yvtAcuZsTdxd5t6bHu+hpCsweZooBbNF8CVbw2T/w2/edEOvHR6/aFb9WaRakASmD+O9RqjPv/vzjOHjnzV/HdT1VV/ZRpOU/AQAA//8Phg61"

func TestServiceHandler_GetProduct(t *testing.T) {
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
			r.Header.Set("Authorization", authHeader)
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
