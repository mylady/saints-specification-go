package discovery

var (
	ProxyBaseAddress    = "http://ip:23003/api/v1"
	ProxyServiceAddress = ProxyBaseAddress + "/services"
)

const (
	discoveryBaseAddress    = "http://localhost:23001/api/v1"
	DiscoveryServiceAddress = discoveryBaseAddress + "/services"
	DiscoveryProductAddress = discoveryBaseAddress + "/products"
	DiscoveryNetworkAddress = discoveryBaseAddress + "/network"
	SendMessageApi          = "/send"
	DefaultPort             = 8092
	RegisterInterval        = "5s"
)
