package manager

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/delegate"
)

func (s *ServiceManager) GetCustomer(phone string) *Customer {
	var rtn Customer
	if phone != "" {
		var pcus ProxyCustomer
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/customer/get/"+phone, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pcus)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				rtn.ID = pcus.ID
				rtn.FirstName = pcus.FirstName
				rtn.LastName = pcus.LastName
				rtn.PhoneNumber = pcus.PhoneNumber
			}
		}
	}

	return &rtn
}

func (s *ServiceManager) GetCustomerAdresses(cid int64) *[]delegate.Address {
	var rtn = []delegate.Address{}
	if cid != 0 {
		var pads []delegate.ProxyAddress
		scid := strconv.FormatInt(cid, 10)
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/addresses/get/"+scid, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pads)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				rtn = *s.Delegate.ProcessCustomerAddresses(&pads)
			}
		}
	}
	return &rtn
}
