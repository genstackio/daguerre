package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

var variants = map[string]func(ctx *commons.Ctx) *diagram.Node{
	"create-react-app": AwsFrameworkCreateReactAppNode,
	"razzle":           AwsFrameworkRazzleNode,
}

func createWebclientNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "web-application",
		Type:     "webapp",
		ListName: "webclients",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("webclients", m)
			utils.PopulateNodes[commons.WebclientConfig](ctx, t, l.Name, l.Hidden, lt.Webclients, lt, l, m, 70, func(config commons.WebclientConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{
					"variant": {
						Type:  "string",
						Value: config.Framework,
					},
				}
			})
			utils.PopulateNodes[commons.WebclientConfig](ctx, t, l.Name, l.Hidden, l.Webclients, lt, l, m, 70, func(config commons.WebclientConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{
					"variant": {
						Type:  "string",
						Value: config.Framework,
					},
				}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.WebclientConfig](ctx, "webclients", lt.Webclients, l.Webclients, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			node := AwsFrameworkWebNode(ctx)
			if len(n.Variant) > 0 {
				if v, found := variants[n.Variant]; found {
					node = v(ctx)
				}
			}
			return utils.EnrichDiagramNode(node, n, label)
		},
	}
}
