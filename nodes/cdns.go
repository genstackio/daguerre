package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Cdns(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_cloudfront_distribution", m)
	utils.PopulateNodes[commons.CdnConfig](t, l.Name, l.Hidden, lt.Cdns, lt, l, m)
	utils.PopulateNodes[commons.CdnConfig](t, l.Name, l.Hidden, l.Cdns, lt, l, m)
}
