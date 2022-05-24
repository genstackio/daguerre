package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createPersonaNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "persona",
		Type:     "persona",
		ListName: "personae",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("personae", m)
			utils.PopulateNodes[commons.PersonaConfig](ctx, t, l.Name, l.Hidden, lt.Personae, lt, l, m, 150)
			utils.PopulateNodes[commons.PersonaConfig](ctx, t, l.Name, l.Hidden, l.Personae, lt, l, m, 150)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.PersonaConfig](ctx, "personae", lt.Personae, l.Personae, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.General.User(), n, label)
		},
	}
}
