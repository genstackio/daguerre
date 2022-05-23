package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Topics(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_sns_topic", m)
	utils.PopulateNodes[commons.TopicConfig](t, l.Name, l.Hidden, lt.Topics, lt, l, m)
	utils.PopulateNodes[commons.TopicConfig](t, l.Name, l.Hidden, l.Topics, lt, l, m)
}
