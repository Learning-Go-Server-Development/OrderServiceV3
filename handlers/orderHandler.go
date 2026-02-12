package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func (h *ServiceHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/order/add"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		ok := h.checkContent(r)
		if !ok {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var o manager.Order
			po, err := h.processBody(r, &o)
			log.Println("bs: ", po)
			log.Println("err: ", err)
			if !po || err != nil {
				http.Error(w, "Trouble parsing body", http.StatusBadRequest)
			} else {
				or := h.Manager.AddOrder(&o)
				log.Println("oo: ", or)
				if or.Success && or.ID != 0 {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(or)
					fmt.Fprint(w, string(resJSON))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (h *ServiceHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/order/update"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		Ok := h.checkContent(r)
		if !Ok {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var o manager.Order
			po, err := h.processBody(r, &o)
			log.Println("bs: ", po)
			log.Println("err: ", err)
			if !po || err != nil {
				http.Error(w, "Trouble parsing body", http.StatusBadRequest)
			} else {
				or := h.Manager.UpdateOrder(&o)
				log.Println("oo: ", or)
				if or.Success {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(or)
					fmt.Fprint(w, string(resJSON))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (h *ServiceHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/order/get"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var oidStr = vars["id"]
			oid, oiderr := strconv.ParseInt(oidStr, 10, 64)
			if oiderr == nil {
				o := h.Manager.GetOrder(oid)
				if o != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(o)
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

func (h *ServiceHandler) GetCurrentOrders(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/orders/current"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var cidStr = vars["cid"]
			cid, ciderr := strconv.ParseInt(cidStr, 10, 64)
			if ciderr == nil {
				os := h.Manager.GetCurrentOrders(cid)
				if os != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(os)
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

func (h *ServiceHandler) GetPastOrders(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/orders/past"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var cidStr = vars["cid"]
			cid, ciderr := strconv.ParseInt(cidStr, 10, 64)
			if ciderr == nil {
				os := h.Manager.GetPastOrders(cid)
				if os != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(os)
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

func (h *ServiceHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/order/delete"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var oidStr = vars["id"]
			oid, oiderr := strconv.ParseInt(oidStr, 10, 64)
			if oiderr == nil {
				or := h.Manager.DeleteCurrentOrder(oid)
				if or != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(or)
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
