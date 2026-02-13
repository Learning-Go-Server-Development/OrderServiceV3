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

var authHeader = "Bearer eNq8lEuT6jYQhX9RbvkxcK+Xw8NCim3ih2ShHZKYsWwZqHEwln99SoRJzSJTsLjlfffpcz61+mBQxYFQW4UgHqGbKNjBYzYTSziHzZmSJQp+HAxKd1R33AsbWJ9/wlaPwtgC2aeUOPsVVNESjbKEVqjmQGt+zN52PtIMkIsEuuKb2NZUskxtzYY7oSvBu4napOc57GBLXuzQuBB+vGqcuMB+rK5qX4YOrE9DMorrtoj9bSHMW35VEgQfrJzhA11oWJ/UrkxqRpMRqqvaecOZlTMnt7NqbKIlOoo2vEgbqk26fUkuMgxq0ZLG1u3LsBPGBkc9B4TgRo+wPvODQat7lpx7g86x1Vu7cX5VuA06lsM5PJKR0Vv+z9oFB0MvPdIQgtC9ZxYXcQeb++x12HEQ+KzMCG4atVWwF8e439NM8zxwBcgqCfBXTZQt4TwpUBBtOpXe8908YL2+8SnWHdTWP7Z8XeERY1l81qbl0HFfahaSPL/3bItX7zvflmnkozHyAsMAMZEXNmwJv2S4abjx0el+L6f0OU40noKP/zQfn1wZCB2WT8EIP7tLDSsHPdEuvTy/S1LLSXZp95DTziejBMHfjMKeecT5mvNfHWze8uHnfzrh7UbcGMYFHqIluvm0GSVNtFDwm5n63tMMifmGpc2mAiNMUEuaONxztVBBxUBm7G1jNJ7K43WbP/TYizbTd4ZGtOQyzZu+P3xTsUE925B6Gj+vD/1wHzWMojHyF5Xw0wn+4np4/i9mmgMy0d2Cz96tioH/50Sp8wMmVbTAOM1fuo/jr70Z+8Xsj1O4KKo/67Wr5x+/+r/ozlxWi9M/AQAA///uFg4d"

func TestServiceHandler_GetProduct(t *testing.T) {
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
