package discovery

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mylady/saints-specification-go"
)

type ServiceRegister struct {
	services   []saints_specification_go.ProtocolService
	regchannel chan<- bool
}

func NewServiceRegister() *ServiceRegister {
	return &ServiceRegister{
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
	if proxy != "" {
		address = ProxyServiceAddress
	}

	for _, service := range this.services {
		data, _ := json.Marshal(service)
		if err := Post(address, data); err != nil {
			fmt.Println(err)
		}
	}
}
