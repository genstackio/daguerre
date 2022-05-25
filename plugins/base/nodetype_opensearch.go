package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createOpensearchNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "opensearch-cluster",
		Type:     "opensearch",
		ListName: "opensearches",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("opensearches", m)
			utils.PopulateNodes[commons.OpensearchConfig](ctx, t, l.Name, l.Hidden, lt.Opensearches, lt, l, m, 150, func(config commons.OpensearchConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.OpensearchConfig](ctx, t, l.Name, l.Hidden, l.Opensearches, lt, l, m, 150, func(config commons.OpensearchConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.OpensearchConfig](ctx, "opensearches", lt.Opensearches, l.Opensearches, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(AwsOpensearchCluster(ctx), n, label)
		},
	}
}
