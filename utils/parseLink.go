package utils

import (
	"github.com/genstackio/daguerre/commons"
	"strings"
)

func ParseLink(ctx *commons.Ctx, s string, m *commons.Model, vars map[string]string) []commons.LinkModel {
	s = ReplaceVars(s, vars)

	x1 := strings.Split(s, " - ")
	x2 := strings.Split(s, " > ")
	x3 := strings.Split(s, " < ")
	x4 := strings.Split(s, " = ")

	y1 := []string{}
	y2 := []string{}

	mod := "-"

	if len(x1) > 1 {
		x1[0] = ConvertAliases(ctx, x1[0])
		x1[1] = ConvertAliases(ctx, x1[1])
		y1 = strings.Split(x1[0], "/")
		y2 = strings.Split(x1[1], "/")
		mod = "-"
	} else if len(x2) > 1 {
		x2[0] = ConvertAliases(ctx, x2[0])
		x2[1] = ConvertAliases(ctx, x2[1])
		y1 = strings.Split(x2[0], "/")
		y2 = strings.Split(x2[1], "/")
		mod = ">"
	} else if len(x3) > 1 {
		x3[0] = ConvertAliases(ctx, x3[0])
		x3[1] = ConvertAliases(ctx, x3[1])
		y1 = strings.Split(x3[0], "/")
		y2 = strings.Split(x3[1], "/")
		mod = "<"
	} else if len(x4) > 1 {
		x4[0] = ConvertAliases(ctx, x4[0])
		x4[1] = ConvertAliases(ctx, x4[1])
		y1 = strings.Split(x4[0], "/")
		y2 = strings.Split(x4[1], "/")
		mod = "="
	}

	llinks := []commons.LinkModel{}

	if len(y1) > 0 {
		if "all" == y1[1] {
			items, found := m.Lists[y1[0]]
			if found {
				for _, item := range items {
					if "all" == y2[1] {
						subitems, subfound := m.Lists[y2[0]]
						if subfound {
							for _, subitem := range subitems {
								llinks = append(llinks, commons.LinkModel{
									From: commons.LinkEndpointModel{Type: y1[0], Name: item.Name},
									To:   commons.LinkEndpointModel{Type: y2[0], Name: subitem.Name},
									Mode: mod,
								})
							}
						}
					} else {
						llinks = append(llinks, commons.LinkModel{
							From: commons.LinkEndpointModel{Type: y1[0], Name: item.Name},
							To:   commons.LinkEndpointModel{Type: y2[0], Name: y2[1]},
							Mode: mod,
						})
					}
				}
			}

			return llinks
		}
		if "all" == y2[1] {
			items, found := m.Lists[y2[0]]
			if found {
				for _, item := range items {
					if "all" == y1[1] {
						subitems, subfound := m.Lists[y1[0]]
						if subfound {
							for _, subitem := range subitems {
								llinks = append(llinks, commons.LinkModel{
									From: commons.LinkEndpointModel{Type: y1[0], Name: subitem.Name},
									To:   commons.LinkEndpointModel{Type: y2[0], Name: item.Name},
									Mode: mod,
								})
							}
						}
					} else {
						llinks = append(llinks, commons.LinkModel{
							From: commons.LinkEndpointModel{Type: y1[0], Name: y1[1]},
							To:   commons.LinkEndpointModel{Type: y2[0], Name: item.Name},
							Mode: mod,
						})
					}
				}
			}
		} else {
			llinks = append(llinks, commons.LinkModel{
				From: commons.LinkEndpointModel{Type: y1[0], Name: y1[1]},
				To:   commons.LinkEndpointModel{Type: y2[0], Name: y2[1]},
				Mode: mod,
			})
		}
	}

	return llinks
}
