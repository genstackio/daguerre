package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func defaultLabeller(ctx *commons.Ctx, n *commons.Node, mode string) string {
	ntn, found := ctx.GetNodeTypeType(n.Type)
	if !found {
		ntn = n.Type
	}
	var label string
	if mode == "full" {
		label = ntn + " - " + n.Name
	} else if mode == "light" {
		label = ntn
	} else if mode == "none" {
		label = ""
	} else {
		label = n.Name
	}
	return label
}
func createDiagramNode(ctx *commons.Ctx, n *commons.Node, withLabel bool) *diagram.Node {
	creator, found := ctx.GetDiagramNodeCreator(n.Type)
	if !found {
		return nil
	}
	labeller, found2 := ctx.GetDiagramNodeLabeller(n.Type)

	mode := "default"

	if withLabel {
		if len(n.Name) > 0 {
			mode = "full"
		} else {
			mode = "light"
		}
	} else {
		mode = "none"
	}

	var label string
	if found2 {
		label = labeller(ctx, n, mode)
	} else {
		label = defaultLabeller(ctx, n, mode)
	}

	return creator(ctx, n, label)
}
