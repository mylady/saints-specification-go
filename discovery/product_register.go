package discovery

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mylady/saints-specification-go"
)

type ProductRegisger struct {
	product    saints_specification_go.ProductInfo
	regchannel chan<- bool
}

func NewProductRegisger(prod saints_specification_go.ProductInfo) *ProductRegisger {
	return &ProductRegisger{
		product: prod,
	}
}

func (this *ProductRegisger) StartRegister() {
	interval, _ := time.ParseDuration(RegisterInterval)
	this.regchannel = SetInterval(this.doregister, interval)
}

func (this *ProductRegisger) StopRegister() {
	this.regchannel <- true
}

func (this *ProductRegisger) GetProduct() saints_specification_go.ProductInfo {
	return this.product
}

func (this *ProductRegisger) doregister() {
	//register product
	data, _ := json.Marshal(this.product)
	if err := Post(DiscoveryProductAddress, data); err != nil {
		fmt.Println(err)
	}
}
