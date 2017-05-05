package discovery

import (
	"encoding/json"
	"errors"

	"strings"

	"github.com/mylady/saints-specification-go"
)

type ServiceDiscovery struct {
	proxy string
}

func NewServiceDiscovery(p string) *ServiceDiscovery {
	return &ServiceDiscovery{
		proxy: p,
	}
}

func (this *ServiceDiscovery) GetService(servtype uint) (target saints_specification_go.ProtocolService, err error) {
	var services []saints_specification_go.ProtocolService
	found := false
	if services, err = this.getServices(); err == nil {
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

func (this *ServiceDiscovery) getServices() (services []saints_specification_go.ProtocolService, err error) {
	var data []byte
	address := DiscoveryServiceAddress
	if this.proxy != "" {
		address = strings.Replace(ProxyServiceAddress, "ip", this.proxy, 1)
	}
	if data, err = Get(address); err == nil {
		err = json.Unmarshal(data, &services)
	}
	return services, err
}
