package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
	"strings"
)

var partners = map[string]func(ctx *commons.Ctx) *diagram.Node{
	"stripe":    PartnerStripeNode,
	"paypal":    PartnerPaypalNode,
	"callr":     PartnerCallrNode,
	"smsfactor": PartnerSmsfactorNode,
	"lydia":     PartnerLydiaNode,
	"slack":     PartnerSlackNode,
}

func createPartnerNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "partner",
		Type:     "partner",
		ListName: "partners",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("partners", m)
			utils.PopulateNodes[commons.PartnerConfig](ctx, t, l.Name, l.Hidden, lt.Partners, lt, l, m, 150, func(config commons.PartnerConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.PartnerConfig](ctx, t, l.Name, l.Hidden, l.Partners, lt, l, m, 150, func(config commons.PartnerConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.PartnerConfig](ctx, "partners", lt.Partners, l.Partners, m, map[string]string{"name": l.Name})
		},
		DiagramNodeLabeller: func(ctx *commons.Ctx, n *commons.Node, mode string) string {
			if mode == "none" {
				return ""
			}
			return strings.ToUpper(n.Name)
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			var node *diagram.Node
			if p, found := partners[n.Name]; found {
				node = p(ctx)
			} else {
				node = aws.Database.Database()
			}
			node.Options.Style = "rounded"

			return utils.EnrichDiagramNode(node, n, label)
		},
	}
}
