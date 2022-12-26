package http

import (
	"strings"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"github.com/ncraft-io/ncraft-gokit/pkg/utils"
)

const (
	EnvelopStyle = "envelope"
	AIPStyle     = "aip"

	underScoreEnvelopStyle = "_envelope"
	underScoreAIPStyle     = "_aip"
)

type Config struct {
	Style    string         `json:"style"` // default, aip, envelope
	Envelope EnvelopeConfig `json:"envelope"`
}

func (c *Config) GetStyle() string {
	if c != nil {
		return c.Style
	}
	return ""
}

func (c *Config) GetEnvelop() *EnvelopeConfig {
	if c != nil {
		return &c.Envelope
	}
	return &EnvelopeConfig{}
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	if err := utils.GetNcraftConfigValue("transport.http").Scan(cfg); err != nil {
		logs.Errorw("failed to get the ncraft.http.transport config.", "path", strings.Join(path, "."), "error", err)
		return nil
	}
	return cfg
}
