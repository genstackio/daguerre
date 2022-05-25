package services

import (
	"encoding/json"
	"github.com/genstackio/daguerre/commons"
	"io/ioutil"
	"os"
)

func loadConfigFile(path string) (*commons.Config, error) {
	jsonFile, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if nil != err {
		return nil, err
	}

	config := &commons.Config{
		Layers:     map[string]commons.LayerConfig{},
		LayerTypes: map[string]commons.LayerConfig{},
		Schemas:    map[string]commons.SchemaConfig{},
	}

	err = json.Unmarshal(byteValue, config)

	if nil != err {
		return nil, err
	}

	return config, nil
}
