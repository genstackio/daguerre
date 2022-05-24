package utils

import "github.com/genstackio/daguerre/commons"

func PopulateAllLinks[T commons.LinksAware](ctx *commons.Ctx, t string, ltItems map[string]T, lItems map[string]T, m *commons.Model, vars map[string]string) {
	if nil != ltItems {
		for k, x := range ltItems {
			PopulateLinks(ctx, x.GetLinks(), m, vars, t+"/"+k+" ")
		}
	}
	if nil != lItems {
		for k, x := range lItems {
			PopulateLinks(ctx, x.GetLinks(), m, vars, t+"/"+k+" ")
		}
	}
}
