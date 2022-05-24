package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createEdgeNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "lambda-edge",
		Type:     "edge",
		ListName: "edges",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("edges", m)
			utils.PopulateNodes[commons.EdgeConfig](ctx, t, l.Name, l.Hidden, lt.Edges, lt, l, m, 80)
			utils.PopulateNodes[commons.EdgeConfig](ctx, t, l.Name, l.Hidden, l.Edges, lt, l, m, 80)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.EdgeConfig](ctx, "edges", lt.Edges, l.Edges, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Compute.Lambda(), n, label)
		},
	}
}
