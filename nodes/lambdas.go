package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Lambdas(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("lambdas", m)
	utils.PopulateNodes[commons.LambdaConfig](t, l.Name, l.Hidden, lt.Lambdas, lt, l, m, 100)
	utils.PopulateNodes[commons.LambdaConfig](t, l.Name, l.Hidden, l.Lambdas, lt, l, m, 100)
}
