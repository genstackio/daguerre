package services

import (
	"errors"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func Build(order *commons.Order) (string, *diagram.Diagram, *commons.Ctx, error) {

	ctx := utils.CreateCtx()

	config, err := loadConfig(ctx, order.Input)

	if nil != err {
		return "", nil, ctx, err
	}

	model, err := buildModelFromConfig(ctx, config)

	if nil != err {
		return "", nil, ctx, err
	}

	filename := order.Output
	if len(filename) == 0 {
		if len(order.Schema) > 0 {
			filename = order.Schema
		} else {
			filename = "schema"
		}
	}

	diag, err := diagram.New(diagram.Filename(filename), diagram.Label(model.Name), diagram.Direction(model.Direction))

	if err != nil {
		return "", diag, ctx, err
	}

	schemaName := "default"
	if len(order.Schema) > 0 {
		schemaName = order.Schema
	}
	schema, found := config.Schemas[schemaName]
	if !found {
		return "", nil, ctx, errors.New("unknown schema '" + schemaName + "'")
	}

	err = loadModel(ctx, diag, model, utils.MergeOrderAndSchemaConfig(order, &schema))

	return filename, diag, ctx, err
}
