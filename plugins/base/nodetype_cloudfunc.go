package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createCloudfuncNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "cloudfront-distribution",
		Type:     "cloudfunc",
		ListName: "cloudfuncs",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("cloudfuncs", m)
			utils.PopulateNodes[commons.CloudfuncConfig](ctx, t, l.Name, l.Hidden, lt.Cloudfuncs, lt, l, m, 70, func(config commons.CloudfuncConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.CloudfuncConfig](ctx, t, l.Name, l.Hidden, l.Cloudfuncs, lt, l, m, 70, func(config commons.CloudfuncConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.CloudfuncConfig](ctx, "cloudfuncs", lt.Cloudfuncs, l.Cloudfuncs, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(AwsCloudfrontFunctionNode(ctx), n, label)
		},
	}
}
