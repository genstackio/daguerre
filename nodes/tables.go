package nodes

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Tables(lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	t := utils.EnsureList("aws_dynamodb_table", m)
	utils.PopulateNodes[commons.TableConfig](t, l.Name, l.Hidden, lt.Tables, lt, l, m)
	utils.PopulateNodes[commons.TableConfig](t, l.Name, l.Hidden, l.Tables, lt, l, m)
}
