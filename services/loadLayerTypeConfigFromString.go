package services

import (
	"encoding/json"
	"github.com/genstackio/daguerre/commons"
	"strings"
)

func loadLayerTypeConfigFromString(s []byte) (commons.LayerConfig, error) {
	var data commons.LayerConfig
	err := json.NewDecoder(strings.NewReader(string(s))).Decode(&data)
	return data, err
}
