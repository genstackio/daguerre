package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Ses(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("ses", m)
	utils.PopulateNodes[commons.SesConfig](t, l.Name, l.Hidden, lt.Ses, lt, l, m, 150)
	utils.PopulateNodes[commons.SesConfig](t, l.Name, l.Hidden, l.Ses, lt, l, m, 150)
}
