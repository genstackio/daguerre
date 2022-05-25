package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createRecordNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "route35-record",
		Type:     "dnsrecord",
		ListName: "records",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("records", m)
			utils.PopulateNodes[commons.RecordConfig](ctx, t, l.Name, l.Hidden, lt.Records, lt, l, m, 5, func(config commons.RecordConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.RecordConfig](ctx, t, l.Name, l.Hidden, l.Records, lt, l, m, 5, func(config commons.RecordConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.RecordConfig](ctx, "records", lt.Records, l.Records, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Network.Route53(), n, label)
		},
	}
}
