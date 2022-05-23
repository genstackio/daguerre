package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Edges(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_lambda_function_edge", m)
	utils.PopulateNodes[commons.EdgeConfig](t, l.Name, l.Hidden, lt.Edges, lt, l, m)
	utils.PopulateNodes[commons.EdgeConfig](t, l.Name, l.Hidden, l.Edges, lt, l, m)
}
