package discovery

import (
	"encoding/json"
	"time"

	"github.com/mylady/saints-specification-go"
)

type ServiceRegister struct {
	services   []saints_specification_go.ProtocolService
	proxy      string
	regchannel chan<- bool
}

func NewServiceRegister(proxy_ip string) *ServiceRegister {
	return &ServiceRegister{
		services: make([]saints_specification_go.ProtocolService, 10),
		proxy:    proxy_ip,
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
		ProxyBaseAddress = proxy
		address = ProxyServiceAddress
	}

	data, _ := json.Marshal(this.services)
	if err := Post(address, data); err != nil {

	}
}
