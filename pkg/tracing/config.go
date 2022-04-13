package tracing

import (
    "github.com/ncraft-io/ncraft-go/pkg/config"
    "github.com/ncraft-io/ncraft-go/pkg/logs"
    "strings"
)

type Config struct {
    Enable bool    `json:"enable" yaml:"Enable" default:"false"`
    Url    string  `json:"url" yaml:"url" default:"localhost:6831"`
    Param  float64 `json:"param" json:"param" default:"100000"`
}

func NewConfig(path ...string) *Config {
    cfg := &Config{}
    err := config.Get(path...).Scan(cfg)
    if err != nil {
        logs.Errorw("failed to get the tracing config from "+strings.Join(path, "."), "error", err.Error())
        return nil
    }
    return cfg
}
