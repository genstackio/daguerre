package services

import (
	"github.com/genstackio/daguerre/commons"
)

func Generate(order *commons.Order) error {

	diag, err := Build(order)

	if nil != err {
		return err
	}

	return diag.Render()
}
