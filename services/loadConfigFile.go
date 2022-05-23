package services

import "github.com/genstackio/daguerre/commons"

func loadConfigFile(path string) (commons.Config, error) {
	return commons.Config{
		Layers: map[string]commons.LayerConfig{},
	}, nil
}
