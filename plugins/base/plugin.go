package base

import (
	"github.com/genstackio/daguerre/commons"
)

type plugin struct {
}

func (p *plugin) Register(ctx *commons.Ctx) {
	ctx.RegisterPluginNodeTypes(&nodeTypes)
	ctx.RegisterNodeTypesAliases(&nodeTypes)
}

func New() commons.Plugin {
	return &plugin{}
}
