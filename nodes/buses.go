package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Buses(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_eventbridge_bus", m)
	utils.PopulateNodes[commons.BusConfig](t, l.Name, l.Hidden, lt.Buses, lt, l, m)
	utils.PopulateNodes[commons.BusConfig](t, l.Name, l.Hidden, l.Buses, lt, l, m)
}
