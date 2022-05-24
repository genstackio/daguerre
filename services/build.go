package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Build(order *commons.Order) (*diagram.Diagram, *commons.Ctx, error) {

	ctx := utils.CreateCtx()

	config, err := loadConfig(ctx, order.Input)

	if nil != err {
		return nil, ctx, err
	}

	model, err := buildModelFromConfig(ctx, config)

	if nil != err {
		return nil, ctx, err
	}

	diag, err := diagram.New(diagram.Filename(order.Output), diagram.Label(model.Name), diagram.Direction(model.Direction))

	if err != nil {
		return diag, ctx, err
	}

	err = loadModel(ctx, diag, model, order)

	return diag, ctx, err
}
