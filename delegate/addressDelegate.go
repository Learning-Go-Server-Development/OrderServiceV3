package delegate

type ProxyAddress struct {
	ID      int64  `json:"id"`
	CID     int64  `json:"cid"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type Address struct {
	ID      int64  `json:"id"`
	CID     int64  `json:"cid"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

func (d *ServiceDelegate) ProcessCustomerAddresses(padds *[]ProxyAddress) *[]Address {
	var rtn = []Address{}
	if len(*padds) > 0 {
		for _, pa := range *padds {
			var a Address
			a.ID = pa.ID
			a.CID = pa.CID
			a.City = pa.City
			a.Street = pa.Street
			a.State = pa.State
			a.ZipCode = pa.ZipCode
			rtn = append(rtn, a)
		}
	}
	return &rtn
}
