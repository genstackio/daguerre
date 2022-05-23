package dnodes

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
)

func Opensearch(n *commons.Node, label string) *diagram.Node {
	return aws.Database.Database().Label(label)
}
