package utils

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
)

func CustomDiagramNode(ctx *commons.Ctx, key string) *diagram.Node {
	ctx.RegisterCustomAsset(key)
	nn := aws.Compute.Lambda()
	nn.Options.Attributes["image"] = "assets/custom/" + key + ".png"

	return nn
}
