package retry

import (
	"strings"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"github.com/ncraft-io/ncraft-gokit/pkg/utils"
)

type Config struct {
	Enable  bool `json:"enable" yaml:"enable" default:"false"`
	Timeout int  `json:"timeout" yaml:"timeout" default:"1000"`
	Max     int  `json:"max" yaml:"max" default:"3"`
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	if err := utils.GetNcraftConfigValue("retry").Scan(cfg); err != nil {
		logs.Errorw("failed to get the ncraft.retry config from "+strings.Join(path, "."), "error", err)
		return nil
	}
	return cfg
}
