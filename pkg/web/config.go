package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//DBConfig :database config
type DBConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	SSLMode     string `json:"ssl_mode"`
	PoolSetting *Pool  `json:"pool_setting"`
}

//Pool :database connection pool config
type Pool struct {
	MaxOpen     int `json:"max_open"`      //max open connection int connection pool
	MaxIdle     int `json:"max_idle"`      //max idle connection in connection pool
	MaxLifeTime int `json:"max_life_time"` //max life time of connection in thread pool
}

//PostgresqlConn :get connection string for postgresql
func (c *DBConfig) PostgresqlConn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}

//AppConfig :config for app
type AppConfig struct {
	Env       string                 `json:"env"`
	Port      int                    `json:"port"`
	Hub       string                 `json:"hub"`
	DB        map[string]*DBConfig   `json:"db"`
	GZip      bool                   `json:"gzip"`
	HTTPS     bool                   `json:"https"`
	HTTPSPort int                    `json:"https_port"`
	CertFile  string                 `json:"cert_file"`
	CertKey   string                 `json:"cert_key"`
	Extra     map[string]interface{} `json:"extra"`
}

//New :create new app with default path
func New() (c *AppConfig, err error) {
	return NewWithPath(DefaultConfigPath)
}

//NewWithPath :create new app config with given path
func NewWithPath(p string) (c *AppConfig, err error) {
	if _, err := os.Stat(p); err != nil {
		return nil, err
	}

	var j []byte
	if j, err = ioutil.ReadFile(p); err != nil {
		return nil, err
	}

	c = &AppConfig{}
	if err = json.Unmarshal(j, c); err != nil {
		return nil, err
	}

	return c, err
}

//EnableDebug :enable debug with env
func (c *AppConfig) EnableDebug() bool {
	if c.Env == EnvDevelopment {
		return true
	}
	return false
}
