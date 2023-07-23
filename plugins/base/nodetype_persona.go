package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
	"strings"
)

func createPersonaNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "persona",
		Type:     "persona",
		ListName: "personae",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("personae", m)
			utils.PopulateNodes[commons.PersonaConfig](ctx, t, l.Name, l.Hidden, lt.Personae, lt, l, m, 150, func(config commons.PersonaConfig) map[string]commons.ParamModel {
				v := "user"
				if config.Multiple {
					v = "users"
				}
				return map[string]commons.ParamModel{
					"variant": {
						Type:  "string",
						Value: v,
					},
				}
			})
			utils.PopulateNodes[commons.PersonaConfig](ctx, t, l.Name, l.Hidden, l.Personae, lt, l, m, 150, func(config commons.PersonaConfig) map[string]commons.ParamModel {
				v := "user"
				if config.Multiple {
					v = "users"
				}
				return map[string]commons.ParamModel{
					"variant": {
						Type:  "string",
						Value: v,
					},
				}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.PersonaConfig](ctx, "personae", lt.Personae, l.Personae, m, map[string]string{"name": l.Name})
		},
		DiagramNodeLabeller: func(ctx *commons.Ctx, n *commons.Node, mode string) string {
			if mode == "none" {
				return ""
			}
			s := ""
			for _, v := range strings.Split(strings.ReplaceAll(n.Name, "-", " "), " ") {
				s = s + " " + strings.ToUpper(v[0:1]) + v[1:]
			}
			return s
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			node := aws.General.User()
			if n.Variant == "users" {
				node = aws.General.Users()
			}
			node.Options.Style = "rounded"
			return utils.EnrichDiagramNode(node, n, label)
		},
	}
}
