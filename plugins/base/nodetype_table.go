package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func createTableNodeType() commons.PluginNodeType {
	return commons.PluginNodeType{
		Name:     "dynamodb-table",
		Type:     "table",
		ListName: "tables",
		NodeCreator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			t := utils.EnsureList("tables", m)
			utils.PopulateNodes[commons.TableConfig](ctx, t, l.Name, l.Hidden, lt.Tables, lt, l, m, 150)
			utils.PopulateNodes[commons.TableConfig](ctx, t, l.Name, l.Hidden, l.Tables, lt, l, m, 150)
		},
		LinkPopulator: func(ctx *commons.Ctx, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
			utils.PopulateAllLinks[commons.TableConfig](ctx, "tables", lt.Tables, l.Tables, m, map[string]string{"name": l.Name})
		},
		DiagramNodeCreator: func(ctx *commons.Ctx, n *commons.Node, label string) *diagram.Node {
			return utils.EnrichDiagramNode(aws.Database.DynamodbTable(), n, label)
		},
		Aliases: map[string]string{
			"tables": "tables/all",
		},
	}
}
