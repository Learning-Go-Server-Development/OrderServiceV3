package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/manager"
	"github.com/Learning-Go-Server-Development/OrderServiceV3/security"
)

type ServiceHandler struct {
	Manager  manager.Manager
	Security security.Security
}

func (h *ServiceHandler) New() Handler {
	return h
}

// CheckContent CheckContent
func (h *ServiceHandler) checkContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

// SetContentType SetContentType
func (h *ServiceHandler) setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// ProcessBody ProcessBody
func (h *ServiceHandler) processBody(r *http.Request, obj any) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	log.Println("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			log.Println("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
