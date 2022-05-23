package services

import "github.com/genstackio/daguerre/commons"

func explodeLink(l commons.LinkModel, ctx *commons.Ctx, expandeds map[string]bool, collapseds map[string]bool, m *commons.Model) (*commons.CtxEntry, *commons.CtxEntry, string, string, string, string) {
	l.To = alignLinkEndpoint(l.To, m)
	f := l.From.Type
	t := l.To.Type
	var a *commons.CtxEntry = nil
	var b *commons.CtxEntry = nil
	c := "?"
	d := "?"
	e := ""
	cluster_name := ""

	if len(l.Label) > 0 {
		e = l.Label
	}
	if items, found := ctx.Items[f]; found {
		if item, found2 := items[l.From.Name]; found2 {
			a = &item
			c = f + "/" + l.From.Name
		} else {
			if items2, found3 := m.Lists[f]; found3 {
				if item2, found4 := items2[l.From.Name]; found4 {
					if len(item2.Cluster) > 0 {
						cluster_name = item2.Cluster
						if item3, found5 := ctx.Clusters[cluster_name]; found5 {
							a = &item3
							c = "clusters/" + cluster_name
						}
					}
				}
			}
		}
	}
	if items, found := ctx.Items[t]; found {
		if item, found2 := items[l.To.Name]; found2 {
			b = &item
			d = t + "/" + l.To.Name
		} else {
			if items2, found3 := m.Lists[t]; found3 {
				if item2, found4 := items2[l.To.Name]; found4 {
					if len(item2.Cluster) > 0 {
						cluster_name = item2.Cluster
						if item3, found5 := ctx.Clusters[cluster_name]; found5 {
							a = &item3
							c = "clusters/" + cluster_name
						}
					}
				}
			}
		}
	}
	if a == b {
		return nil, nil, l.Mode, c, d, e
	}

	return a, b, l.Mode, c, d, e
}
