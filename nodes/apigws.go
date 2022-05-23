package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Apigws(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_apigateway_api", m)
	utils.PopulateNodes[commons.ApigwConfig](t, l.Name, l.Hidden, lt.Apigws, lt, l, m)
	utils.PopulateNodes[commons.ApigwConfig](t, l.Name, l.Hidden, l.Apigws, lt, l, m)
}