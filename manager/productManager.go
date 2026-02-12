package manager

import (
	"log"
	"net/http"
)

func (s *ServiceManager) GetProduct(sku string) *Product {
	var rtn Product
	if sku != "" {
		var pcus ProxyProduct
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/product/get/"+sku, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pcus)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				rtn.ID = pcus.ID
				rtn.SKU = pcus.SKU
				rtn.Description = pcus.Description
				rtn.Price = pcus.Price
			}
		}
	}
	return &rtn
}
