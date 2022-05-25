package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createLambdaLayersNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "lambda-layer",
		Type:     "lambdaLayer",
		ListName: "lambdaLayers",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("lambdaLayers", m)
			utils.PopulateNodes[commons.LambdaLayerConfig](ctx, t, l.Name, l.Hidden, lt.LambdaLayers, lt, l, m, 100, func(config commons.LambdaLayerConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.LambdaLayerConfig](ctx, t, l.Name, l.Hidden, l.LambdaLayers, lt, l, m, 100, func(config commons.LambdaLayerConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.LambdaLayerConfig](ctx, "lambdaLayers", lt.LambdaLayers, l.LambdaLayers, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(AwsLambdaLayerNode(ctx), n, label)
		},
	}
}
