package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/dnodes"
)

func createDiagramNode(n *commons.Node, withLabel bool) *diagram.Node {
	creator, found := dnodes.Dnodes[n.Type]

	if !found {
		return nil
	}

	if withLabel {
		if len(n.Name) > 0 {
			return creator(n, n.Type+" - "+n.Name)
		}
		return creator(n, n.Type)
	}

	return creator(n, n.Name)
}