package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Topics(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("topics", m)
	utils.PopulateNodes[commons.TopicConfig](t, l.Name, l.Hidden, lt.Topics, lt, l, m, 150)
	utils.PopulateNodes[commons.TopicConfig](t, l.Name, l.Hidden, l.Topics, lt, l, m, 150)
}
