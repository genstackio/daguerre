package utils

import (
	"github.com/genstackio/daguerre/commons"
)

func PopulateLinks(ctx *commons.Ctx, links *[]string, m *commons.Model, vars map[string]string, prefix string) {
	if nil == links {
		return
	}
	if nil == m.Links {
		m.Links = &[]commons.LinkModel{}
	}
	for _, v := range *links {
		if len(prefix) > 0 {
			v = prefix + v
		}
		for _, vv := range ParseLink(ctx, v, m, vars) {
			*m.Links = append(*m.Links, vv)
		}
	}
}
