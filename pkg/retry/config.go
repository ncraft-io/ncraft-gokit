package retry

import (
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
    "strings"
)

type Config struct {
    Enable  bool `json:"enable" yaml:"enable" default:"false"`
    Timeout int  `json:"timeout" yaml:"timeout" default:"1000"`
    Max     int  `json:"max" yaml:"max" default:"3"`
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
