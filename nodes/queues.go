package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Queues(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("queues", m)
	utils.PopulateNodes[commons.QueueConfig](t, l.Name, l.Hidden, lt.Queues, lt, l, m, 60)
	utils.PopulateNodes[commons.QueueConfig](t, l.Name, l.Hidden, l.Queues, lt, l, m, 60)
}
