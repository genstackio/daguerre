package dnodes

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
)

func Cloudfunc(n *commons.Node, label string) *diagram.Node {
	return aws.Compute.Lambda().Label(label)
}
