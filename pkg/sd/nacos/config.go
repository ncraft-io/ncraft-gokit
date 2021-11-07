package nacos

import "github.com/nacos-group/nacos-sdk-go/common/constant"

type Config struct {
	//ServerConfig constant.ServerConfig `json:"serverConfig"`
	ClientConfig constant.ClientConfig `json:"clientConfig"`
}
