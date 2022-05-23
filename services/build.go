package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func Build(order *commons.Order) (*diagram.Diagram, error) {

	config, err := loadConfig(order.Input)

	if nil != err {
		return nil, err
	}

	model, err := buildModelFromConfig(config)

	if nil != err {
		return nil, err
	}

	return buildDiagramFromModelAndOrder(model, order)
}
