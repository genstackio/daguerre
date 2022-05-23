package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Opensearches(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_opensearch_cluster", m)
	utils.PopulateNodes[commons.OpensearchConfig](t, l.Name, l.Hidden, lt.Opensearches, lt, l, m)
	utils.PopulateNodes[commons.OpensearchConfig](t, l.Name, l.Hidden, l.Opensearches, lt, l, m)
}
