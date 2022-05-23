package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Cloudfuncs(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("cloudfuncs", m)
	utils.PopulateNodes[commons.CloudfuncConfig](t, l.Name, l.Hidden, lt.Cloudfuncs, lt, l, m, 70)
	utils.PopulateNodes[commons.CloudfuncConfig](t, l.Name, l.Hidden, l.Cloudfuncs, lt, l, m, 70)
}
