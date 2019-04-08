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

func (this *ServiceDiscovery) GetService(servtype uint) (targets []saints_specification_go.ProtocolService, err error) {
	var services []saints_specification_go.ProtocolService
	targets = make([]saints_specification_go.ProtocolService, 0)
	if services, err = this.getServices(); err == nil {
		for _, service := range services {
			if service.ServiceType == servtype {
				targets = append(targets, service)
			}
		}

		if len(targets) == 0 {
			err = errors.New("service not found")
		}
	}
	return targets, err
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
