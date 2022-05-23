package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func buildDiagramFromModelAndOrder(model *commons.Model, order *commons.Order) (*diagram.Diagram, error) {
	diag, err := diagram.New(diagram.Filename(order.Output), diagram.Label(model.Name), diagram.Direction(model.Direction))

	if err != nil {
		return diag, err
	}

	err = loadModel(diag, model, order)

	return diag, err
}
