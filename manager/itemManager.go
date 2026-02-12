package manager

import (
	"net/http"

	"github.com/Learning-Go-Server-Development/OrderServiceV3/database"
)

func (s *ServiceManager) AddItem(i *Item) *ResponseID {
	var rtn ResponseID
	if i != nil {
		var ni database.Item
		ni.OrderID = i.OrderID
		ni.ProductID = i.ProductID
		suc, id := s.DB.AddItem(&ni)
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

func (s *ServiceManager) UpdateItem(i *Item) *Response {
	var rtn Response
	if i != nil {
		var ui database.Item
		ui.ID = i.ID
		ui.OrderID = i.OrderID
		ui.ProductID = i.ProductID
		suc := s.DB.UpdateItem(&ui)
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

func (s *ServiceManager) GetItems(oid int64) *[]Item {
	var rtn = []Item{}
	if oid != 0 {
		ids := s.DB.GetItems(oid)
		if len(*ids) > 0 {
			for _, i := range *(ids) {
				var ii Item
				ii.ID = i.ID
				ii.OrderID = i.OrderID
				ii.ProductID = i.ProductID
				rtn = append(rtn, ii)
			}
		}
	}
	return &rtn
}

func (s *ServiceManager) DeleteItem(iid int64) *Response {
	var rtn Response
	if iid != 0 {
		suc := s.DB.DeleteItem(iid)
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
