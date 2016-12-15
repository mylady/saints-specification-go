package discovery

import (
	"encoding/json"
	"time"

	"github.com/mylady/saints-specification-go"
)

type ProductRegister struct {
	product    saints_specification_go.ProductInfo
	regchannel chan<- bool
}

func NewProductRegister(prod saints_specification_go.ProductInfo) *ProductRegister {
	return &ProductRegister{
		product: prod,
	}
}

func (this *ProductRegister) StartRegister() {
	interval, _ := time.ParseDuration(RegisterInterval)
	this.regchannel = SetInterval(this.doregister, interval)
}

func (this *ProductRegister) StopRegister() {
	this.regchannel <- true
}

func (this *ProductRegister) GetProduct() saints_specification_go.ProductInfo {
	return this.product
}

func (this *ProductRegister) doregister() {
	//register product
	data, _ := json.Marshal(this.product)
	Post(DiscoveryProductAddress, data)
}
