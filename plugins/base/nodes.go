package base

import (
	"github.com/genstackio/daguerre/commons"
)

var nodeTypes = []commons.PluginNodeType{
	createApigwNodeType(),
	createBucketNodeType(),
	createBusNodeType(),
	createCdnNodeType(),
	createCloudfuncNodeType(),
	createEdgeNodeType(),
	createLambdaNodeType(),
	createOpensearchNodeType(),
	createPartnerNodeType(),
	createPersonaNodeType(),
	createQueueNodeType(),
	createRecordNodeType(),
	createSesNodeType(),
	createTableNodeType(),
	createZoneNodeType(),
	createTopicNodeType(),
}
