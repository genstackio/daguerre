package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func loadModel(ctx *commons.Ctx, d *diagram.Diagram, m *commons.Model, o *commons.Order) error {

	l := prepareModelListings(ctx, m, o)

	rootGroup := diagram.NewGroup("platform").Label("platform")

	loadModelClusters(ctx, m, l, rootGroup)

	d.Group(rootGroup)

	if nil != m.Lists {
		if ps, found := m.Lists["personae"]; found {
			for i, p := range ps {
				_, found1 := l.RequiredSpecials["personae"][i]
				_, found2 := l.RequiredSpecials["personae"]["all"]
				if found1 || found2 {
					if nil == ctx.Items {
						ctx.Items = map[string]map[string]commons.CtxEntry{}
					}
					if _, found3 := ctx.Items["personae"]; !found3 {
						ctx.Items["personae"] = map[string]commons.CtxEntry{}
					}
					if cc, found4 := p.Params["multiple"]; found4 && cc.Type == "bool" && cc.BoolValue {
						dnode := createDiagramNode(ctx, &commons.Node{
							Type: "personae",
							Name: i,
						}, true)
						ctx.Items["personae"][i] = commons.CtxEntry{
							Dnode: dnode,
						}
					}
				}
			}
		}
		if ps, found := m.Lists["partners"]; found {
			for i, _ := range ps {
				_, found1 := l.RequiredSpecials["partners"][i]
				_, found2 := l.RequiredSpecials["partners"]["all"]
				if found1 || found2 {
					if nil == ctx.Items {
						ctx.Items = map[string]map[string]commons.CtxEntry{}
					}
					if _, found3 := ctx.Items["partners"]; !found3 {
						ctx.Items["partners"] = map[string]commons.CtxEntry{}
					}
					dnode := createDiagramNode(ctx, &commons.Node{
						Type: "partners",
						Name: i,
					}, true)
					ctx.Items["partners"][i] = commons.CtxEntry{
						Dnode: dnode,
					}
				} else {
					dnode := createDiagramNode(ctx, &commons.Node{
						Type: "unknown",
						Name: i,
					}, true)
					ctx.Items["partners"][i] = commons.CtxEntry{
						Dnode: dnode,
					}
				}
			}
		}
	}

	keptLinks := map[string]commons.KeptLinkModel{}

	for _, ll := range *m.Links {
		a, b, md, c, d, e := explodeLink(ctx, ll, m)
		kk := c + md + d
		if v, found := keptLinks[kk]; found {
			v.Count = v.Count + 1
			if len(e) == 0 {
				v.Labels[e] = true
			} else {
				v.Label = e
			}
		} else {
			v := commons.KeptLinkModel{From: a, To: b, Count: 1, Labels: map[string]bool{}, Mode: md, FullName: ll.ToString()}
			keptLinks[kk] = v
			if len(e) > 0 {
				v.Labels[e] = true
			}
		}
	}

	for _, l := range keptLinks {
		a := l.From
		b := l.To
		cc := l.Count
		md := l.Mode
		if nil != a && nil != b {
			if "-" == md {
				if cc > 1 {
					connectCtxEntries(d, a, b, l.Mode, "("+string(cc)+")")
				} else {
					connectCtxEntries(d, a, b, l.Mode, l.Label)
				}
			} else if ">" == md {
				if cc > 1 {
					connectCtxEntries(d, a, b, l.Mode, "("+string(cc)+")")
				} else {
					connectCtxEntries(d, a, b, l.Mode, l.Label)
				}
			} else if "<" == md {
				if cc > 1 {
					connectCtxEntries(d, a, b, l.Mode, "("+string(cc)+")")
				} else {
					connectCtxEntries(d, a, b, l.Mode, l.Label)
				}
			} else if "=" == md {
				if cc > 1 {
					connectCtxEntries(d, a, b, l.Mode, "("+string(cc)+")")
				} else {
					connectCtxEntries(d, a, b, l.Mode, l.Label)
				}
			} else if cc > 1 {
				connectCtxEntries(d, a, b, l.Mode, "("+string(cc)+")")
			} else {
				connectCtxEntries(d, a, b, l.Mode, l.Label)
			}
		}
	}

	return nil
}
