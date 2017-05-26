// Package util includes objects and classes that are used by multiple services
package util

import (
	"encoding/json"
	"os"
)

// ConfigFileName is the path where the config file is stored
const ConfigFileName = "/var/lib/kontainerooo/config.json"

// ConfigFile represents the contents of the config file
type ConfigFile struct {
	RootfsPath           string
	CustomerPath         string
	NetNSPath            string
	StandardPathVariable string
}

var configLoaded = false
var conf = ConfigFile{}

// GetConfig loads and returns the config
func GetConfig() (ConfigFile, error) {
	if !configLoaded {
		f, err := os.Open(ConfigFileName)
		if err != nil {
			return ConfigFile{}, err
		}

		d := json.NewDecoder(f)
		err = d.Decode(&conf)
		if err != nil {
			return ConfigFile{}, err
		}

		return conf, nil
	}

	return conf, nil
}
