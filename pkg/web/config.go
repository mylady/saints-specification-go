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
	Env       string         `json:"env"`
	Port      int            `json:"port"`
	Hub       string         `json:"hub"`
	GZip      bool           `json:"gzip"`
	Session   *SessionConfig `json:"session"`
	HTTPS     bool           `json:"https"`
	HTTPSPort int            `json:"https_port"`
	CertFile  string         `json:"cert_file"`
	CertKey   string         `json:"cert_key"`
}

//SessionConfig :http session config
type SessionConfig struct {
	SignKeys  []string `json:"sign_keys"`
	CookieKey string   `json:"cookie_key"`
	MaxAge    int      `json:"max_age"`
	Secure    bool     `json:"secure"`
	HTTPOnly  bool     `json:"http_only"`
	Renew     bool     `json:"renew"`
}

//Config :config for web app
type Config struct {
	App   *AppConfig           `json:"app"`
	DB    map[string]*DBConfig `json:"db"`
	Extra interface{}          `json:"extra"`
}

//NewConfig :create new with default path
func NewConfig() (c *Config, err error) {
	return NewConfigWithPath(DefaultConfigPath)
}

//NewConfigWithPath :create new config with given path
func NewConfigWithPath(p string) (c *Config, err error) {
	if _, err := os.Stat(p); err != nil {
		return nil, err
	}

	var j []byte
	if j, err = ioutil.ReadFile(p); err != nil {
		return nil, err
	}

	c = &Config{}
	if err = json.Unmarshal(j, c); err != nil {
		return nil, err
	}

	return c, err
}

//IsDebug :if debug is enabled
func (c *Config) IsDebug() bool {
	if c.App.Env == EnvDevelopment {
		return true
	}
	return false
}

//CanHTTPS :can run https server
func (c *Config) CanHTTPS() bool {
	if _, err := os.Stat(c.App.CertFile); err != nil {
		return false
	}

	if _, err := os.Stat(c.App.CertKey); err != nil {
		return false
	}

	if c.App.HTTPSPort <= 0 {
		return false
	}

	return c.App.HTTPS
}
