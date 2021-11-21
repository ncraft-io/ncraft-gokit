package server

import (
	"github.com/ncraft-io/ncraft-go/pkg/config"
	"github.com/ncraft-io/ncraft-go/pkg/logs"
	"strings"
)

type Config struct {
	DebugAddr string `json:"debugAddr" yaml:"debugAddr" default:":20170"`
	HttpAddr  string `json:"httpAddr" yaml:"httpAddr" default:":20171"`
	GrpcAddr  string `json:"grpcAddr" yaml:"grpcAddr" default:":20172"`
	UdpAddr   string `json:"udpAddr" yaml:"udpAddr" default:":20173"`
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
