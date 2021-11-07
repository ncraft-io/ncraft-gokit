package metrics

import (
	"github.com/ncraft-io/ncraft-go/pkg/config"
)

type Config struct {
	Enable     bool   `json:"enable" default:"true"`
	Department string `json:"department"`
	Project    string `json:"project"`
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	err := config.GetValue(path...).Scan(cfg)
	if err != nil {
		//logs.Errorw("failed to get the server config from "+strings.Join(path, "."), "error", err.Error())
		return nil
	}

	return cfg
}
