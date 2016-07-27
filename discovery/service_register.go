package discovery

import (
	"encoding/json"
	"time"

	"github.com/mylady/saints-specification-go"
)

type ServiceRegisger struct {
	services   []saints_specification_go.ProtocolService
	proxy      string
	regchannel chan<- bool
}

func NewServiceRegisger(proxy_ip string) *ServiceRegisger {
	return &ServiceRegisger{
		services: make([]saints_specification_go.ProtocolService, 10),
		proxy:    proxy_ip,
	}
}

func (this *ServiceRegisger) AddService(service saints_specification_go.ProtocolService) {
	this.services = append(this.services, service)
}

func (this *ServiceRegisger) StartRegister() {
	interval, _ := time.ParseDuration(RegisterInterval)
	this.regchannel = SetInterval(this.doregister, interval)
}

func (this *ServiceRegisger) StopRegister() {
	this.regchannel <- true
}

func (this *ServiceRegisger) getServices() []saints_specification_go.ProtocolService {
	return this.services
}

func (this *ServiceRegisger) doregister() {
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
