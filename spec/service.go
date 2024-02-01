package spec

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Service :saints service struct
type Service struct {
	Address  string `json:"address"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
}

// ServiceIPHolder :holder for address
const ServiceIPHolder = "ip"

// ServiceRegisterInterval :register interval
const ServiceRegisterInterval = 10 * time.Second

// Service type enum
const (
	ServiceTypeIdentity = iota
	ServiceTypeFileHub
	ServiceTypeDevice
	ServiceTypeCloudProxy
	ServiceTypeMonitor
	ServiceTypeMessageHub
	ServiceTypeEvent
	ServiceTypeLog
	ServiceTypeTimeLine
)

// ServiceTypeDict :service type name dict
var ServiceTypeDict = map[int]string{
	ServiceTypeIdentity:   "身份认证服务",
	ServiceTypeFileHub:    "文件管理服务",
	ServiceTypeDevice:     "设备信息服务",
	ServiceTypeCloudProxy: "云代理服务",
	ServiceTypeMonitor:    "实时监控服务",
	ServiceTypeMessageHub: "消息服务",
	ServiceTypeEvent:      "事件服务",
	ServiceTypeLog:        "日志服务",
	ServiceTypeTimeLine:   "时间流服务",
}

// ServiceRegister :service register
type ServiceRegister struct {
	service    Service
	hub        string
	timer      *SimpleTicker
	httpClient *http.Client
}

// NewServiceRegister :new service register
func NewServiceRegister(ip string, service Service) (register *ServiceRegister, err error) {
	r := &ServiceRegister{
		service:    service,
		hub:        fmt.Sprintf("http://%s:%d/rest/register", ip, ServiceHubPort),
		httpClient: http.DefaultClient,
	}

	r.timer = NewSimpleTicker(ServiceRegisterInterval, r.register)
	return r, err
}

func (r *ServiceRegister) register(ticker *SimpleTicker) {
	var req *http.Request
	var resp *http.Response
	var err error

	var serviceJSON []byte
	if serviceJSON, err = json.Marshal(r.service); err != nil {
		fmt.Printf("service serialize failed %s\n", err)
		return
	}

	jsonReader := strings.NewReader(string(serviceJSON))
	if req, err = http.NewRequest("POST", r.hub, jsonReader); err != nil {
		fmt.Printf("create request failed %s\n", err)
		return
	}

	if resp, err = r.httpClient.Do(req); err != nil {
		fmt.Printf("registe service failed %s\n", err)
		return
	}

	var respBytes []byte
	if respBytes, err = io.ReadAll(resp.Body); err != nil {
		fmt.Printf("registe service failed %s\n", err)
		return
	}

	var restResp RestResponse
	if err = json.Unmarshal(respBytes, &restResp); err != nil {
		fmt.Printf("registe service failed %s\n", err)
		return
	}

	if !restResp.Result {
		fmt.Printf("registe service failed %s\n", restResp.ErrMsg)
	}
}

// Start :start register
func (r *ServiceRegister) Start() {
	r.timer.Start()
}

// Stop :stop register
func (r *ServiceRegister) Stop() {
	r.timer.Stop()
}
