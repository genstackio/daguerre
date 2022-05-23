package dnodes

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
)

func Zone(n *commons.Node, label string) *diagram.Node {
	return aws.Network.Route53().Label(label)
}
