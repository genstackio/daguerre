package services

import "github.com/genstackio/daguerre/commons"

var items = []string{
	"aws_lambdas", "aws_opensearch_clusters", "aws_cloudfront_distributions", "aws_apigateways",
	"aws_s3_buckets", "aws_route53_zones", "aws_route53_records", "aws_ses", "aws_dynamodb_tables",
	"aws_sns_topics", "aws_lambda_edges", "aws_cloudfront_functions", "aws_eventbridges",
	"personae", "partners",
}

func initCtx(ctx *commons.Ctx) {
	ctx.Items = map[string]map[string]commons.CtxEntry{}

	for _, v := range items {
		ctx.Items[v] = map[string]commons.CtxEntry{}
	}
}
