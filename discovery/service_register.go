package discovery

import (
	"encoding/json"
	"time"

	"strings"
)

type ServiceRegister struct {
	proxy      string
	services   []saints_specification_go.ProtocolService
	regchannel chan<- bool
}

func NewServiceRegister(p string) *ServiceRegister {
	return &ServiceRegister{
		proxy:    p,
		services: make([]saints_specification_go.ProtocolService, 0),
	}
}

func (this *ServiceRegister) AddService(service saints_specification_go.ProtocolService) {
	this.services = append(this.services, service)
}

func (this *ServiceRegister) StartRegister() {
	interval, _ := time.ParseDuration(RegisterInterval)
	this.regchannel = SetInterval(this.doregister, interval)
}

func (this *ServiceRegister) StopRegister() {
	this.regchannel <- true
}

func (this *ServiceRegister) getServices() []saints_specification_go.ProtocolService {
	return this.services
}

func (this *ServiceRegister) doregister() {
	//register services

	address := DiscoveryServiceAddress
	if this.proxy != "" {
		address = strings.Replace(ProxyServiceAddress, "ip", this.proxy, 1)
	}

	for _, service := range this.services {
		data, _ := json.Marshal(service)
		Post(address, data)
	}
}
