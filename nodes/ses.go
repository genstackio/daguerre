package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Ses(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_ses_cluster", m)
	utils.PopulateNodes[commons.SesConfig](t, l.Name, l.Hidden, lt.Ses, lt, l, m)
	utils.PopulateNodes[commons.SesConfig](t, l.Name, l.Hidden, l.Ses, lt, l, m)
}
