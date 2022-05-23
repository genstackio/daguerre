package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Records(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("records", m)
	utils.PopulateNodes[commons.RecordConfig](t, l.Name, l.Hidden, lt.Records, lt, l, m, 5)
	utils.PopulateNodes[commons.RecordConfig](t, l.Name, l.Hidden, l.Records, lt, l, m, 5)
}
