package tracing

import (
	"strings"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"github.com/ncraft-io/ncraft-gokit/pkg/utils"
)

type Config struct {
	Enable bool    `json:"enable" yaml:"Enable" default:"false"`
	Url    string  `json:"url" yaml:"url" default:"localhost:6831"`
	Param  float64 `json:"param" json:"param" default:"100000"`
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	if err := utils.GetNcraftConfigValue("tracing").Scan(cfg); err != nil {
		logs.Errorw("failed to get the ncraft.tracing config from "+strings.Join(path, "."), "error", err)
		return nil
	}
	return cfg
}
