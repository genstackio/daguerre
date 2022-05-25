package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createTopicNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "sns-topic",
		Type:     "topic",
		ListName: "topics",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("topics", m)
			utils.PopulateNodes[commons.TopicConfig](ctx, t, l.Name, l.Hidden, lt.Topics, lt, l, m, 150, func(config commons.TopicConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.TopicConfig](ctx, t, l.Name, l.Hidden, l.Topics, lt, l, m, 150, func(config commons.TopicConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.TopicConfig](ctx, "topics", lt.Topics, l.Topics, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Integration.SimpleNotificationServiceSns(), n, label)
		},
	}
}
