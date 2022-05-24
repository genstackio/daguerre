package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createLambdaNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "lambda",
		Type:     "lambda",
		ListName: "lambdas",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("lambdas", m)
			utils.PopulateNodes[commons.LambdaConfig](ctx, t, l.Name, l.Hidden, lt.Lambdas, lt, l, m, 100)
			utils.PopulateNodes[commons.LambdaConfig](ctx, t, l.Name, l.Hidden, l.Lambdas, lt, l, m, 100)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.LambdaConfig](ctx, "lambdas", lt.Lambdas, l.Lambdas, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Compute.Lambda(), n, label)
		},
	}
}
