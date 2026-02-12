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

func (h *ServiceHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/item/add"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		Ok := h.checkContent(r)
		if !Ok {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var i manager.Item
			pi, err := h.processBody(r, &i)
			log.Println("bs: ", pi)
			log.Println("err: ", err)
			if !pi || err != nil {
				http.Error(w, "Trouble parsing body", http.StatusBadRequest)
			} else {
				ir := h.Manager.AddItem(&i)
				log.Println("oo: ", ir)
				if ir.Success && ir.ID != 0 {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(ir)
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

func (h *ServiceHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/item/update"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		Ok := h.checkContent(r)
		if !Ok {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var i manager.Item
			pi, err := h.processBody(r, &i)
			log.Println("bs: ", pi)
			log.Println("err: ", err)
			if !pi || err != nil {
				http.Error(w, "Trouble parsing body", http.StatusBadRequest)
			} else {
				ir := h.Manager.UpdateItem(&i)
				log.Println("oo: ", ir)
				if ir.Success {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(ir)
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

func (h *ServiceHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/items"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var oidStr = vars["oid"]
			oid, oiderr := strconv.ParseInt(oidStr, 10, 64)
			if oiderr == nil {
				is := h.Manager.GetItems(oid)
				if is != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(is)
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

func (h *ServiceHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var c jv.Claim
	c.Role = "user"
	c.URL = "/rs/item/delete"
	auth := h.Security.ValidateToken(&c, r)
	if auth {
		h.setContentType(w)
		vars := mux.Vars(r)
		log.Println("vars: ", len(vars))
		if len(vars) == 1 {
			var iidStr = vars["id"]
			iid, iiderr := strconv.ParseInt(iidStr, 10, 64)
			if iiderr == nil {
				ir := h.Manager.DeleteItem(iid)
				if ir != nil {
					w.WriteHeader(http.StatusOK)
					resJSON, _ := json.Marshal(ir)
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
