package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createBucketNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "s3-bucket",
		Type:     "bucket",
		ListName: "buckets",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("buckets", m)
			utils.PopulateNodes[commons.BucketConfig](ctx, t, l.Name, l.Hidden, lt.Buckets, lt, l, m, 30)
			utils.PopulateNodes[commons.BucketConfig](ctx, t, l.Name, l.Hidden, l.Buckets, lt, l, m, 30)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.BucketConfig](ctx, "buckets", lt.Buckets, l.Buckets, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Storage.SimpleStorageServiceS3(), n, label)
		},
	}
}
