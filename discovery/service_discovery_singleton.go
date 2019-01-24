package discovery

import (
	"errors"

	"github.com/mylady/saints-specification-go"
)

var servdiscov *ServiceDiscovery

func InitServiceDiscovery(proxy string) {
	servdiscov = NewServiceDiscovery(proxy)
}

func GetService(servtype uint) (targets []saints_specification_go.ProtocolService, err error) {
	if servdiscov == nil {
		err = errors.New("not initialized")
	} else {
		targets, err = servdiscov.GetService(servtype)
	}
	return
}
