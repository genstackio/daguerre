package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func AwsLambdaEdgeNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "aws-lambda-edge")
}
func PartnerStripeNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-stripe")
}
