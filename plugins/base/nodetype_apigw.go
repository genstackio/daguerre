package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createApigwNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "api-gateway",
		Type:     "api-gw",
		ListName: "apigws",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("apigws", m)
			utils.PopulateNodes[commons.ApigwConfig](ctx, t, l.Name, l.Hidden, lt.Apigws, lt, l, m, 70)
			utils.PopulateNodes[commons.ApigwConfig](ctx, t, l.Name, l.Hidden, l.Apigws, lt, l, m, 70)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.ApigwConfig](ctx, "apigws", lt.Apigws, l.Apigws, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Network.ApiGateway(), n, label)
		},
	}
}
