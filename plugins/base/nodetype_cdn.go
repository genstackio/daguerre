package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createCdnNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "cloudfront-distribution",
		Type:     "distrib",
		ListName: "cdns",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("cdns", m)
			utils.PopulateNodes[commons.CdnConfig](ctx, t, l.Name, l.Hidden, lt.Cdns, lt, l, m, 50, func(config commons.CdnConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
			utils.PopulateNodes[commons.CdnConfig](ctx, t, l.Name, l.Hidden, l.Cdns, lt, l, m, 50, func(config commons.CdnConfig) map[string]commons.ParamModel {
				return map[string]commons.ParamModel{}
			})
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.CdnConfig](ctx, "cdns", lt.Cdns, l.Cdns, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Network.Cloudfront(), n, label)
		},
		LinkEndpointAligner: func(ctx *commons.Ctx, le commons.LinkEndpointModel, m *commons.Model) commons.LinkEndpointModel {
			tt := le.Type
			x := m.Lists[tt][le.Name]

			if len(x.LayerType.Cloudfuncs) > 0 {
				var clf string
				for k, _ := range x.LayerType.Cloudfuncs {
					if len(clf) > 0 {
						break
					}
					clf = k
				}
				if len(clf) > 0 {
					return commons.LinkEndpointModel{Name: clf, Type: "cloudfuncs"}
				}
				return le
			}
			return le
		},
	}
}
