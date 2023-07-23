package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func loadModelClusters(ctx *commons.Ctx, m *commons.Model, l *commons.ModelListings, rootGroup *diagram.Group) {
	for kk := range m.Clusters {
		if "personae" == kk || "partners" == kk {
			continue
		}
		_, found1 := l.Expandeds[kk]
		_, found2 := l.Expandeds["all"]
		_, found11 := l.Collapseds[kk]
		_, found21 := l.Collapseds["all"]
		if (found1 || found2) && !(found11 || found21) {
			g := rootGroup.NewGroup(kk).Label(kk)
			for _, v := range *m.Clusters[kk].Nodes {
				if _, found := m.Lists[v.Type]; found {
					if !v.Hidden {
						dnode := createDiagramNode(ctx, &v, true)
						if nil != dnode {
							if nil == ctx.Items[v.Type] {
								ctx.Items[v.Type] = map[string]commons.CtxEntry{}
							}
							ctx.Items[v.Type][v.Name] = commons.CtxEntry{
								Dnode: dnode,
							}
							g.Add(dnode)
						}
					}
				}
			}
		} else {
			_, found1 = l.Collapseds[kk]
			_, found2 = l.Collapseds["all"]
			if found1 || found2 {
				dgroup := createDiagramGroup(ctx, rootGroup, kk, m)
				if nil == ctx.Clusters {
					ctx.Clusters = map[string]commons.CtxEntry{}
				}
				ctx.Clusters[kk] = commons.CtxEntry{
					Dgroup: dgroup,
				}
			} else {
				_, found1 = l.Requireds[kk]
				_, found2 = l.Requireds["all"]
				if found1 || found2 {
					dgroup := createDiagramGroup(ctx, rootGroup, kk, m)
					if nil == ctx.Clusters {
						ctx.Clusters = map[string]commons.CtxEntry{}
					}
					ctx.Clusters[kk] = commons.CtxEntry{
						Dgroup: dgroup,
					}
				}
			}
		}
	}

}
