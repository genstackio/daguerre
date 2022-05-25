package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createBusNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "bus",
		Type:     "bus",
		ListName: "buses",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("buses", m)
			utils.PopulateNodes[commons.BusConfig](ctx, t, l.Name, l.Hidden, lt.Buses, lt, l, m, 150, func(config commons.BusConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.BusConfig](ctx, t, l.Name, l.Hidden, l.Buses, lt, l, m, 150, func(config commons.BusConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.BusConfig](ctx, "buses", lt.Buses, l.Buses, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Integration.Eventbridge(), n, label)
		},
		Aliases: map[string]string{
			"bus": "buses/default",
		},
	}
}
