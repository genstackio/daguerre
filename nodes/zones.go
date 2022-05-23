package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Zones(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("zones", m)
	utils.PopulateNodes[commons.ZoneConfig](t, l.Name, l.Hidden, lt.Zones, lt, l, m, 10)
	utils.PopulateNodes[commons.ZoneConfig](t, l.Name, l.Hidden, l.Zones, lt, l, m, 10)
}
