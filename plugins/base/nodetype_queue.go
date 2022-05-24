package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createQueueNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "sqs-queue",
		Type:     "queue",
		ListName: "queues",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("queues", m)
			utils.PopulateNodes[commons.QueueConfig](ctx, t, l.Name, l.Hidden, lt.Queues, lt, l, m, 60)
			utils.PopulateNodes[commons.QueueConfig](ctx, t, l.Name, l.Hidden, l.Queues, lt, l, m, 60)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.QueueConfig](ctx, "queues", lt.Queues, l.Queues, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Integration.SimpleQueueServiceSqs(), n, label)
		},
	}
}
