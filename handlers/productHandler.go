package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/GolangToolKits/grrt"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func (h *ServiceHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/product"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var sku = vars["sku"]
			if sku != "" {
				p := h.Manager.GetProduct(sku)
				if p != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(p)
					fmt.Fprint(w, string(resJSON))
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
