package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createZoneNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "route53-zone",
		Type:     "zone",
		ListName: "zones",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("zones", m)
			utils.PopulateNodes[commons.ZoneConfig](ctx, t, l.Name, l.Hidden, lt.Zones, lt, l, m, 10)
			utils.PopulateNodes[commons.ZoneConfig](ctx, t, l.Name, l.Hidden, l.Zones, lt, l, m, 10)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.ZoneConfig](ctx, "zones", lt.Zones, l.Zones, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Network.Route53(), n, label)
		},
	}
}
