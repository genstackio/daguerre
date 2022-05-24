package utils

import "github.com/genstackio/daguerre/commons"

func CreateCtx() *commons.Ctx {
	return &commons.Ctx{
		NodeTypes:     map[string]commons.PluginNodeType{},
		NodeListNames: map[string]struct{ NodeTypeIndex string }{},
		Clusters:      map[string]commons.CtxEntry{},
		Items:         map[string]map[string]commons.CtxEntry{},
		Aliases:       map[string]string{},
		CustomAssets:  map[string]bool{},
	}
}
