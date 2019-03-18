package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mylady/saints-specification-go/pkg/util"
)

//Service :saints service struct
type Service struct {
	Address  string `json:"address"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
}

//ServiceIPHolder :holder for address
const ServiceIPHolder = "ip"

//ServiceRegisterInterval :register interval
const ServiceRegisterInterval = 10 * time.Second

//Service type enum
const (
	ServiceTypeIdentity = iota
	ServiceTypeDevice
	ServiceTypeFileHub
	ServiceTypeGIS
)

//ServiceTypeDict :service type name dict
var ServiceTypeDict = map[int]string{
	ServiceTypeIdentity: "身份认证服务",
	ServiceTypeDevice:   "设备信息服务",
	ServiceTypeFileHub:  "文件管理服务",
	ServiceTypeGIS:      "地理信息服务",
}

//ServiceRegister :service register
type ServiceRegister struct {
	service    Service
	hub        string
	timer      *util.SimpleTicker
	httpClient *http.Client
	req        *http.Request
}

//NewServiceRegister :new service register
func NewServiceRegister(ip string, service Service) (register *ServiceRegister, err error) {
	r := &ServiceRegister{
		service:    service,
		hub:        fmt.Sprintf("http://%s:%d/rest/register", ip, ServiceHubPort),
		httpClient: http.DefaultClient,
	}

	var serviceJSON []byte
	if serviceJSON, err = json.Marshal(r.service); err != nil {
		return nil, err
	}

	jsonReader := strings.NewReader(string(serviceJSON))
	if r.req, err = http.NewRequest("POST", r.hub, jsonReader); err != nil {
		return nil, err
	}

	r.timer = util.NewSimpleTicker(ServiceRegisterInterval, r.register)
	return r, err
}

func (r *ServiceRegister) register() {
	var resp *http.Response
	var err error

	if resp, err = r.httpClient.Do(r.req); err != nil {
		fmt.Printf("registe service failed %s", err)
		return
	}

	var respBytes []byte
	if respBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("registe service failed %s", err)
		return
	}

	var restResp RestResponse
	if err = json.Unmarshal(respBytes, &restResp); err != nil {
		fmt.Printf("registe service failed %s", err)
		return
	}

	if !restResp.Result {
		fmt.Printf("registe service failed %s", restResp.ErrMsg)
	}
}

//Start :start register
func (r *ServiceRegister) Start() {
	r.timer.Start()
}

//Stop :stop register
func (r *ServiceRegister) Stop() {
	r.timer.Stop()
}
