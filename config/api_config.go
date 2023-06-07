package config

import (
	"encoding/json"
	"fmt"
)

const (
	DefaultConfigFile                          = ""
	DefaultListenHost                          = "17.0.0.1"
	DefaultListenPort                          = 9104
	DefaultEnv                                 = "prod"
	DefaultTokenExpire                         = 24 * 3600
	DefaultQueryMySQLUser                      = "root"
	DefaultQueryMySQLPassword                  = "root"
	DefaultAdminMySQLUser                      = "root"
	DefaultAdminMySQLPassword                  = "root"
	DefaultForwardRequestDialTimeout           = 10
	DefaultForwardRequestResponseHeaderTimeout = 5
)

type ApiConfig struct {
	ConfigFile                          string `json:"config_file" toml:"config_file"`                                                         // 配置文件
	ListenHost                          string `json:"listen_host" toml:"listen_host"`                                                         // api监听host
	ListenPort                          int64  `json:"listen_port" toml:"listen_port"`                                                         // api监听port
	Env                                 string `json:"env" toml:"env"`                                                                         // 环境 prod, dev
	TokenExpire                         int64  `json:"token_expire" toml:"token_expire"`                                                       // token过期时间单位s
	QueryMysqlUser                      string `json:"query_mysql_user" toml:"query_mysql_user"`                                               // 查询数据库用户名
	QueryMysqlPassword                  string `json:"query_mysql_password" toml:"query_mysql_password"`                                       // 查询数据库管理员密码
	AdminMysqlUser                      string `json:"admin_mysql_user" toml:"admin_mysql_user"`                                               // 查询数据库用户名
	AdminMysqlPassword                  string `json:"admin_mysql_password" toml:"admin_mysql_password"`                                       // 查询数据库管理员密码
	ForwardRequestDialTimeout           int64  `json:"forward_request_dial_timeout" toml:"forward_request_dial_timeout"`                       // 请求转发链接超时时间
	ForwardRequestResponseHeaderTimeout int64  `json:"forward_request_response_header_timeout" toml:"forward_request_response_header_timeout"` // 请求转发Response超时时间
}

func (this *ApiConfig) DeepClone() (*ApiConfig, error) {
	raw, err := json.Marshal(this)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, ApiConfig -> Json: %v", err.Error())
	}

	var apiConfig ApiConfig
	err = json.Unmarshal(raw, &apiConfig)
	if err != nil {
		return nil, fmt.Errorf("启动配置DeepClone出错, Json -> ApiConfig: %v", err.Error())
	}

	return &apiConfig, nil
}

func (this *ApiConfig) Address() string {
	return fmt.Sprintf("%v:%v", this.ListenHost, this.ListenPort)
}
