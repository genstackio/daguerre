package base

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func AwsLambdaNode(ctx *commons.Ctx) *diagram.Node {
	return aws.Compute.Lambda()
}
func AwsLambdaLayerNode(ctx *commons.Ctx) *diagram.Node {
	return aws.Compute.Lambda()
}
func AwsFrameworkCreateReactAppNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "framework-create-react-app")
}
func AwsFrameworkRazzleNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "framework-razzle")
}
func AwsFrameworkWebNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "framework-web")
}
func AwsCloudfrontFunctionNode(ctx *commons.Ctx) *diagram.Node {
	return utils.CustomDiagramNode(ctx, "aws-cloudfront-function")
}
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
