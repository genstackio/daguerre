package services

import (
	"encoding/json"
	"github.com/genstackio/daguerre/assets"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/plugins/base"
	"strings"
)

func loadConfig(ctx *commons.Ctx, path string) (*commons.Config, error) {
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

	entries, err := assets.LayerTypes.ReadDir("layer-types")

	if nil != err {
		return nil, err
	}

	if nil == config.LayerTypes {
		config.LayerTypes = map[string]commons.LayerConfig{}
	}
	for _, v := range entries {
		c, err := assets.LayerTypes.ReadFile("layer-types/" + v.Name())
		if nil != err {
			return nil, err
		}
		var lt commons.LayerConfig
		err = json.NewDecoder(strings.NewReader(string(c))).Decode(&lt)
		if nil != err {
			return nil, err
		}
		if len(lt.Name) == 0 {
			lt.Name = strings.ReplaceAll(v.Name(), "__", "/")
			lt.Name = strings.ReplaceAll(lt.Name, ".json", "")
		}
		config.LayerTypes[lt.Name] = lt
	}

	var plugins = map[string]commons.Plugin{}

	plugins["base"] = base.New() // @todo make it more dynamic for loading multiple plugins

	for _, p := range plugins {
		p.Register(ctx)
	}

	for k := range ctx.NodeListNames {
		ctx.Items[k] = map[string]commons.CtxEntry{}
	}

	return config, nil
}
