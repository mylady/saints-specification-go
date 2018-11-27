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
const ServiceRegisterInterval = 10

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
	srv        Service
	hub        string
	timer      *util.SimpleTimer
	httpClient *http.Client
	req        *http.Request
}

//NewServiceRegister :new service register
func NewServiceRegister(ip string, srv Service) *ServiceRegister {
	s := &ServiceRegister{
		srv:        srv,
		hub:        fmt.Sprintf("http://%s:%d/rest/register", ip, ServiceHubPort),
		httpClient: http.DefaultClient,
	}

	srvJSON, _ := json.Marshal(srv)
	jsonReader := strings.NewReader(string(srvJSON))
	s.req, _ = http.NewRequest("POST", s.hub, jsonReader)
	s.timer = util.NewSimpleTimer(ServiceRegisterInterval*time.Second, s.register)

	return s
}

func (s *ServiceRegister) register() {
	var resp *http.Response
	var err error

	if resp, err = s.httpClient.Do(s.req); err != nil {
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
func (s *ServiceRegister) Start(srv Service) {
	s.timer.Start()
}

//Stop :stop register
func (s *ServiceRegister) Stop() {
	s.timer.Stop()
}
