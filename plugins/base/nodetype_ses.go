package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createSesNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "ses",
		Type:     "ses",
		ListName: "ses",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("ses", m)
			utils.PopulateNodes[commons.SesConfig](ctx, t, l.Name, l.Hidden, lt.Ses, lt, l, m, 150)
			utils.PopulateNodes[commons.SesConfig](ctx, t, l.Name, l.Hidden, l.Ses, lt, l, m, 150)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.SesConfig](ctx, "ses", lt.Ses, l.Ses, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Engagement.SimpleEmailServiceSes(), n, label)
		},
	}
}
