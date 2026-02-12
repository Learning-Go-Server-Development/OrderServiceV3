package delegate

type Delegate interface {
	ProcessCustomerAddresses(padds *[]ProxyAddress) *[]Address
}

type ServiceDelegate struct {
}

func (d *ServiceDelegate) New() Delegate {
	return d
}
