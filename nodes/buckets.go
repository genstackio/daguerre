package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Buckets(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_s3_bucket", m)
	utils.PopulateNodes[commons.BucketConfig](t, l.Name, l.Hidden, lt.Buckets, lt, l, m)
	utils.PopulateNodes[commons.BucketConfig](t, l.Name, l.Hidden, l.Buckets, lt, l, m)
}
