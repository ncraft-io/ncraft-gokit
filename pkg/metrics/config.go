package metrics

import (
	"strings"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"github.com/ncraft-io/ncraft-gokit/pkg/utils"
)

type Config struct {
	Enable     bool   `json:"enable" default:"true"`
	Department string `json:"department"`
	Project    string `json:"project"`
}

func (c *Config) Enabled() bool {
	if c != nil {
		return c.Enable
	}
	return false
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}

	if err := utils.GetNcraftConfigValue("metrics").Scan(cfg); err != nil {
		logs.Warnw("failed to get the ncraft.metrics config from ", "path", strings.Join(path, "."), "error", err)
		return nil
	}

	return cfg
}
