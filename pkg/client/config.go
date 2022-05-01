package client

import (
    "github.com/ncraft-io/ncraft-gokit/pkg/sd"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
    "strings"
)

type Config struct {
    sd.Config
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
