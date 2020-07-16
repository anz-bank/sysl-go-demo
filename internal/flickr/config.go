package flickr

import (
	"github.com/anz-bank/sysl-go/config"
)

// ConfigContainer struct holds the config for the generated code and lib.
type ConfigContainer struct {
	config.DefaultConfig
	CustomConfig
}

// RootConfig struct holds the adminServer configuration loaded from config yml.
type RootConfig struct {
	CustomConfig CustomConfig `yaml:"custom"`
}

// CustomConfig struct holds the adminServer configuration loaded from config yml.
type CustomConfig struct {
	AdminServer config.CommonHTTPServerConfig `yaml:"adminServer"`
}
