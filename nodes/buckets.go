package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Buckets(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("buckets", m)
	utils.PopulateNodes[commons.BucketConfig](t, l.Name, l.Hidden, lt.Buckets, lt, l, m, 30)
	utils.PopulateNodes[commons.BucketConfig](t, l.Name, l.Hidden, l.Buckets, lt, l, m, 30)
}
