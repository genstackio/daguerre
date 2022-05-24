package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

var partners = map[string]func(ctx *commons.Ctx) *diagram.Node{
	"stripe": PartnerStripeNode,
}

func createPartnerNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "partner",
		Type:     "partner",
		ListName: "partners",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("partners", m)
			utils.PopulateNodes[commons.PartnerConfig](ctx, t, l.Name, l.Hidden, lt.Partners, lt, l, m, 150)
			utils.PopulateNodes[commons.PartnerConfig](ctx, t, l.Name, l.Hidden, l.Partners, lt, l, m, 150)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.PartnerConfig](ctx, "partners", lt.Partners, l.Partners, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			if p, found := partners[n.Name]; found {
				return utils.EnrichDiagramNode(p(ctx), n, label)
			}

			return utils.EnrichDiagramNode(aws.Database.Database(), n, label)
		},
	}
}
