package sd

import (
    "github.com/ncraft-io/ncraft-go/pkg/config"
    "github.com/ncraft-io/ncraft-gokit/pkg/retry"
    "github.com/ncraft-io/ncraft-gokit/pkg/sd/direct"
    "github.com/ncraft-io/ncraft-gokit/pkg/sd/etcdv3"
    "github.com/ncraft-io/ncraft-gokit/pkg/sd/nacos"
)

type Config struct {
    Mode   string                    `json:"mode" yaml:"mode" db:"mode"`
    Url    string                    `json:"url" yaml:"url"`
    Retry  *retry.Config             `json:"retry" yaml:"retry" db:"retry"`
    EtcdV3 *etcdv3.Config            `json:"etcd" yaml:"etcd"`
    Nacos  *nacos.Config             `json:"nacos" yaml:"nacos"`
    Direct map[string]*direct.Config `json:"direct" yaml:"direct" db:"direct"`
}

func NewConfig(path string) *Config {
    cfg := &Config{}
    if err := config.ScanKey(path, cfg); err != nil {
        return nil
    }
    return cfg
}
