package services

import (
	"errors"
	"github.com/genstackio/daguerre/commons"
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
		return nil, errors.New("unknown layer type '" + l.Type + "'")
	}

	return &lt, nil
}
