package manager

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
)

func (s *ServiceManager) AddOrder(o *Order) *ResponseID {
	var rtn ResponseID
	if o != nil {
		var no database.Order
		no.CID = o.CustomerID
		no.OrderNumber = o.OrderNumber
		suc, id := s.DB.AddOrder(&no)
		if suc {
			rtn.ID = id
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	} else {
		rtn.Code = http.StatusBadRequest
		rtn.Success = false
	}
	return &rtn
}

func (s *ServiceManager) UpdateOrder(o *Order) *Response {
	var rtn Response
	if o != nil {
		var uo database.Order
		uo.ID = o.ID
		uo.CID = o.CustomerID
		uo.OrderNumber = o.OrderNumber
		suc := s.DB.UpdateOrder(&uo)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	} else {
		rtn.Code = http.StatusBadRequest
		rtn.Success = false
	}
	return &rtn
}

func (s *ServiceManager) GetOrder(id int64) *Order {
	var rtn Order
	if id != 0 {
		od := s.DB.GetOrder(id)
		if od != nil && od.ID != 0 {
			rtn.ID = od.ID
			rtn.CustomerID = od.CID
			rtn.Entered = od.DateEntered
			rtn.Updated = od.DateUpdated
			rtn.OrderNumber = od.OrderNumber
		}
	}
	return &rtn
}

func (s *ServiceManager) GetCurrentOrders(cid int64) *[]Order {
	var rtn = []Order{}
	if cid != 0 {
		ods := s.DB.GetAllOrders(cid)
		if len(*ods) > 0 {
			for _, o := range *(ods) {
				var oo Order
				oo.ID = o.ID
				oo.CustomerID = o.CID
				oo.OrderNumber = o.OrderNumber
				oo.Entered = o.DateEntered
				oo.Updated = o.DateUpdated
				rtn = append(rtn, oo)
			}
		}
	}
	return &rtn
}

func (s *ServiceManager) GetPastOrders(cid int64) *[]Order {
	var rtn = []Order{}
	if cid != 0 {
		var pods []ProxyOrder
		scid := strconv.FormatInt(cid, 10)
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/orders/get/"+scid, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pods)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				for _, po := range pods {
					var o Order
					o.ID = po.ID
					o.CustomerID = po.CID
					o.OrderNumber = po.OID
					rtn = append(rtn, o)
				}
			}
		}
	}
	return &rtn
}

func (s *ServiceManager) DeleteCurrentOrder(id int64) *Response {
	var rtn Response
	if id != 0 {
		suc := s.DB.DeleteOrder(id)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusInternalServerError
		}
	} else {
		rtn.Code = http.StatusBadRequest
		rtn.Success = false
	}
	return &rtn
}
