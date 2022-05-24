package services

import (
	"github.com/genstackio/daguerre/assets"
	"github.com/genstackio/daguerre/commons"
	"os"
)

func Generate(order *commons.Order) error {

	diag, ctx, err := Build(order)

	if nil != err {
		return err
	}

	if err := os.MkdirAll("./go-diagrams/assets/custom", os.ModePerm); err != nil {
		return err
	}

	if nil != ctx.CustomAssets {
		for k := range ctx.CustomAssets {
			content, err := assets.Images.ReadFile("images/" + k + ".png")
			if nil != err {
				return err
			}
			err = os.WriteFile("./go-diagrams/assets/custom/"+k+".png", content, os.ModePerm)
			if nil != err {
				return err
			}

		}
	}
	return diag.Render()
}
