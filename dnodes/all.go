package dnodes

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

var Dnodes map[string]func(n *commons.Node, label string) *diagram.Node = map[string]func(n *commons.Node, label string) *diagram.Node{
	"apigws":       Apigw,
	"buckets":      Bucket,
	"buses":        Bus,
	"cdns":         Cdn,
	"cloudfuncs":   Cloudfunc,
	"edges":        Edge,
	"lambdas":      Lambda,
	"opensearches": Opensearch,
	"partners":     Partner,
	"personae":     Persona,
	"queues":       Queue,
	"records":      Record,
	"ses":          Ses,
	"tables":       Table,
	"zones":        Zone,
}
