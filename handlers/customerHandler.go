package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func (h *ServiceHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/customer/get"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var p = vars["phone"]
			if p != "" {
				c := h.Manager.GetCustomer(p)
				if c != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(c)
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

func (h *ServiceHandler) GetCustomerAddresses(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/customer/addresses/get"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var cidStr = vars["cid"]
			cid, ciderr := strconv.ParseInt(cidStr, 10, 64)
			if ciderr == nil {
				al := h.Manager.GetCustomerAdresses(cid)
				if al != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(al)
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
