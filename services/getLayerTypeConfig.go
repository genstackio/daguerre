package services

import (
	"github.com/genstackio/daguerre/commons"
	"log"
)

func getLayerTypeConfig(l *commons.LayerConfig, c *commons.Config) (*commons.LayerConfig, error) {
	if len(l.Type) == 0 {
		l.Type = "custom"
	}
	if "custom" == l.Type {
		return l, nil
	}

	lt, found := c.LayerTypes[l.Type]

	if !found {
		log.Println("layer type > unknown '" + l.Type + "'")
		return &commons.LayerConfig{}, nil
	}

	return &lt, nil
}
