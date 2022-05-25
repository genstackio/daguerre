package services

import (
	"github.com/genstackio/daguerre/commons"
	"log"
)

func explodeLink(ctx *commons.Ctx, l commons.LinkModel, m *commons.Model) (*commons.CtxEntry, *commons.CtxEntry, string, string, string, string) {
	if aligner, found := ctx.GetLinkEndpointAligner(l.To.Type); found {
		l.To = aligner(ctx, l.To, m)
	}

	f := l.From.Type
	t := l.To.Type
	var a *commons.CtxEntry = nil
	var b *commons.CtxEntry = nil
	c := "?"
	d := "?"
	e := ""
	clusterName := ""

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
						clusterName = item2.Cluster
						if item3, found5 := ctx.Clusters[clusterName]; found5 {
							a = &item3
							c = "clusters/" + clusterName
						}
					}
				}
			} else {
				log.Println("link > unknown 'from' name : " + l.ToString())
			}
		}
	} else {
		log.Println("link > unknown 'from' type : " + l.ToString())
	}
	if items, found := ctx.Items[t]; found {
		if item, found2 := items[l.To.Name]; found2 {
			b = &item
			d = t + "/" + l.To.Name
		} else {
			if items2, found3 := m.Lists[t]; found3 {
				if item2, found4 := items2[l.To.Name]; found4 {
					if len(item2.Cluster) > 0 {
						clusterName = item2.Cluster
						if item3, found5 := ctx.Clusters[clusterName]; found5 {
							b = &item3
							d = "clusters/" + clusterName
						}
					}
				}
			} else {
				log.Println("link > unknown 'to' name : " + l.ToString())
			}
		}
	} else {
		log.Println("link > unknown 'to' type : " + l.ToString())
	}

	if nil != a && nil != b {
		if a == b {
			return nil, nil, l.Mode, c, d, e
		}
		if a.Dnode != nil && a.Dnode == b.Dnode {
			return nil, nil, l.Mode, c, d, e
		}
		if a.Dgroup != nil && a.Dgroup == b.Dgroup {
			return nil, nil, l.Mode, c, d, e
		}
	}

	return a, b, l.Mode, c, d, e
}
