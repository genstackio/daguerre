package services

import (
	"embed"
	"github.com/genstackio/daguerre/commons"
	"strings"
)

//go:embed configs/layer-types/*.json
var embeddedLayerTypesConfigs embed.FS

func loadConfig(path string) (*commons.Config, error) {
	defaultName := "platform"
	defaultGraphAttr := map[string]string{
		"fontsize": "12",
	}

	config, err := loadConfigFile(path)

	if nil != err {
		return nil, err
	}

	if len(config.Name) == 0 {
		config.Name = defaultName
	}
	if len(config.Styles) == 0 {
		config.Styles = defaultGraphAttr
	}
	if len(config.Direction) == 0 {
		config.Direction = "LR"
	}

	entries, err := embeddedLayerTypesConfigs.ReadDir("configs/layer-types")

	if nil != err {
		return nil, err
	}

	if nil == config.LayerTypes {
		config.LayerTypes = map[string]commons.LayerConfig{}
	}
	for _, v := range entries {
		c, err := embeddedLayerTypesConfigs.ReadFile("configs/layer-types/" + v.Name())
		lt, err := loadLayerTypeConfigFromString(c)
		if len(lt.Name) == 0 {
			lt.Name = strings.ReplaceAll(v.Name(), "__", "/")
			lt.Name = strings.ReplaceAll(lt.Name, ".json", "")
		}
		if nil != err {
			return nil, err
		}
		config.LayerTypes[lt.Name] = lt
	}

	return &config, nil
}
