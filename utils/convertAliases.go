package utils

import "github.com/genstackio/daguerre/commons"

func ConvertAliases(ctx *commons.Ctx, s string) string {
	if nil == ctx.Aliases {
		return s
	}
	v, found := ctx.Aliases[s]
	if found {
		return v
	}
	return s
}
