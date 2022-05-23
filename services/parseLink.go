package services

import (
	"github.com/genstackio/daguerre/commons"
	"strings"
)

func parseLink(s string, m *commons.Model) []commons.LinkModel {
	x1 := strings.Split(s, " - ")
	x2 := strings.Split(s, " > ")
	x3 := strings.Split(s, " < ")

	y1 := []string{}
	y2 := []string{}

	mod := "-"

	if len(x1) > 1 {
		x1[0] = convertAliases(x1[0])
		x1[1] = convertAliases(x1[1])
		y1 = strings.Split(x1[0], "/")
		y2 = strings.Split(x1[1], "/")
		mod = "-"
	} else {
		if len(x2) > 1 {
			x2[0] = convertAliases(x2[0])
			x2[1] = convertAliases(x2[1])
			y1 = strings.Split(x2[0], "/")
			y2 = strings.Split(x2[1], "/")
			mod = ">"
		} else {
			if len(x3) > 1 {
				x3[0] = convertAliases(x3[0])
				x3[1] = convertAliases(x3[1])
				y1 = strings.Split(x3[0], "/")
				y2 = strings.Split(x3[1], "/")
				mod = "<"
			}
		}
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

			return llinks
		}

		llinks = append(llinks, commons.LinkModel{
			From: commons.LinkEndpointModel{Type: y1[0], Name: y1[1]},
			To:   commons.LinkEndpointModel{Type: y2[0], Name: y2[1]},
			Mode: mod,
		})
	}

	return llinks
}
