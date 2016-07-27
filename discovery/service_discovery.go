package discovery

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/mylady/saints-specification-go"
)

var proxy string

func SetProxy(proxy_ip string) {
	proxy = proxy_ip
	ProxyBaseAddress = strings.Replace(ProxyBaseAddress, "ip", proxy_ip, 1)
}

func GetService(servtype uint) (target saints_specification_go.ProtocolService, err error) {
	var services []saints_specification_go.ProtocolService
	found := false
	if services, err = getServices(); err == nil {
		for _, service := range services {
			if service.ServiceType == servtype {
				target = service
				found = true
				break
			}
		}

		if !found {
			err = errors.New("service not found")
		}
	}
	return target, err
}

func getServices() (services []saints_specification_go.ProtocolService, err error) {
	var data []byte
	address := DiscoveryServiceAddress
	if proxy != "" {
		address = ProxyServiceAddress
	}
	if data, err = Get(address); err == nil {
		err = json.Unmarshal(data, &services)
	}
	return services, err
}
