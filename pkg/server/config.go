package server

import (
	"github.com/ncraft-io/ncraft-go/pkg/config"
	"github.com/ncraft-io/ncraft-go/pkg/logs"
	"strings"
)

type Config struct {
	HttpAddr  string `json:"httpAddr" yaml:"httpAddr" default:":10000"`
	GrpcAddr  string `json:"grpcAddr" yaml:"grpcAddr" default:":10001"`
	DebugAddr string `json:"debugAddr" yaml:"debugAddr" default:":10002"`
	UdpAddr   string `json:"udpAddr" yaml:"udpAddr" default:":10003"`
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	err := config.Get(path...).Scan(cfg)
	if err != nil {
		logs.Errorw("failed to get the server config from "+strings.Join(path, "."), "error", err.Error())
		return nil
	}
	return cfg
}
