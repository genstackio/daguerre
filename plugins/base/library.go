package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func AwsLambdaEdgeNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "aws-lambda-edge")
}
func AwsOpensearchCluster(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "aws-opensearch-cluster")
}
func PartnerStripeNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-stripe")
}
func PartnerPaypalNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-paypal")
}
func PartnerCallrNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-callr")
}
func PartnerSmsfactorNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-smsfactor")
}
func PartnerLydiaNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-lydia")
}
func PartnerSlackNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "partner-slack")
}
