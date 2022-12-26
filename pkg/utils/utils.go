package utils

import (
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config/reader"
)

func GetNcraftConfigValue(defaultPath string, path ...string) reader.Value {
	if len(path) == 0 {
		if value := config.Get("ncraft." + defaultPath); !value.Null() {
			return value
		}
		return config.Get(defaultPath)
	} else {
		return config.Get(path...)
	}
}
