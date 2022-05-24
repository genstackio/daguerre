package utils

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func EnrichDiagramNode(node *diagram.Node, n *commons.Node, label string) *diagram.Node {
	if len(label) > 0 {
		node.Label(label)
	} else {
		node.Label("_")
		node.Options.LabelLocation = "none"
		node.Options.Font = diagram.Font{
			Name:  "Sans-Serif",
			Size:  0,
			Color: "white",
		}
	}
	return node
}
