package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func loadModel(ctx *commons.Ctx, d *diagram.Diagram, m *commons.Model, o *commons.Order) error {

	expandeds := map[string]bool{}
	collapseds := map[string]bool{}
	requireds := map[string]bool{}
	requiredSpecials := map[string]map[string]bool{
		"personae": {},
		"partners": {},
	}

	for _, v := range o.Expand {
		expandeds[v] = true
	}
	for _, v := range o.Clusters {
		requireds[v] = true
	}
	for _, v := range o.Personae {
		requiredSpecials["personae"][v] = true
	}
	for _, v := range o.Partners {
		requiredSpecials["partners"][v] = true
	}
	for _, v := range o.Show {
		_, found1 := expandeds[v]
		_, found2 := collapseds["all"]
		if !found1 && !found2 {
			collapseds[v] = true
		}
	}

	if nil != m.Links {
		for _, l := range *m.Links {
			a, b, c, d := extractClusterFromLink(l, m)
			_, found := expandeds[a]
			_, found2 := expandeds[b]
			_, found3 := expandeds["all"]
			if found || found2 || found3 {
				if "personae" == a || "partners" == a {
					requiredSpecials[a][c] = true
				} else if len(a) > 0 {
					requireds[a] = true
				}
				if "personae" == b || "partners" == b {
					requiredSpecials[b][d] = true
				} else if len(b) > 0 {
					requireds[b] = true
				}
			}
		}
	}

	rootGroup := diagram.NewGroup("platform").Label("platform")

	for kk := range m.Clusters {
		if "personae" == kk || "partners" == kk {
			continue
		}
		_, found1 := expandeds[kk]
		_, found2 := expandeds["all"]
		if found1 || found2 {
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
			_, found1 = collapseds[kk]
			_, found2 = collapseds["all"]
			if found1 || found2 {
				dgroup := createDiagramGroup(ctx, rootGroup, kk, m)
				if nil == ctx.Clusters {
					ctx.Clusters = map[string]commons.CtxEntry{}
				}
				ctx.Clusters[kk] = commons.CtxEntry{
					Dgroup: dgroup,
				}
			} else {
				_, found1 = requireds[kk]
				_, found2 = requireds["all"]
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

	d.Group(rootGroup)

	if nil != m.Lists {
		if ps, found := m.Lists["personae"]; found {
			for i, p := range ps {
				_, found1 := requiredSpecials["personae"][i]
				_, found2 := requiredSpecials["personae"]["all"]
				if found1 || found2 {
					if nil == ctx.Items {
						ctx.Items = map[string]map[string]commons.CtxEntry{}
					}
					if _, found3 := ctx.Items["personae"]; !found3 {
						ctx.Items["personae"] = map[string]commons.CtxEntry{}
					}
					if cc, found4 := p.Params["multiple"]; found4 && cc.Type == "bool" && cc.BoolValue {
						dnode := createDiagramNode(ctx, &commons.Node{
							Type: "users",
							Name: i,
						}, true)
						ctx.Items["personae"][i] = commons.CtxEntry{
							Dnode: dnode,
						}
					} else {
						dnode := createDiagramNode(ctx, &commons.Node{
							Type: "user",
							Name: i,
						}, true)
						ctx.Items["personae"][i] = commons.CtxEntry{
							Dnode: dnode,
						}
					}
				} else {
					dnode := createDiagramNode(ctx, &commons.Node{
						Type: "user",
						Name: i,
					}, true)
					ctx.Items["personae"][i] = commons.CtxEntry{
						Dnode: dnode,
					}
				}
			}
		}
		if ps, found := m.Lists["partners"]; found {
			for i, _ := range ps {
				_, found1 := requiredSpecials["partners"][i]
				_, found2 := requiredSpecials["partners"]["all"]
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
			v := commons.KeptLinkModel{From: a, To: b, Count: 1, Labels: map[string]bool{}, Mode: md}
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
