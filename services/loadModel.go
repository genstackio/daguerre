package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
	"log"
)

func loadModel(d *diagram.Diagram, m *commons.Model, o *commons.Order) error {
	ctx := commons.Ctx{}

	initCtx(&ctx)

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
				} else {
					if len(a) > 0 {
						requireds[a] = true
					}
				}
				if "personae" == b || "partners" == b {
					requiredSpecials[b][d] = true
				} else {
					if len(b) > 0 {
						requireds[b] = true
					}
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
						log.Println(kk, v.Type, v.Name)
						dnode := createDiagramNode(&v, true)
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
				dgroup := createDiagramGroup(rootGroup, kk, m)
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
					dgroup := createDiagramGroup(rootGroup, kk, m)
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
						dnode := createDiagramNode(&commons.Node{
							Type: "users",
							Name: i,
						}, true)
						ctx.Items["personae"][i] = commons.CtxEntry{
							Dnode: dnode,
						}
					} else {
						dnode := createDiagramNode(&commons.Node{
							Type: "user",
							Name: i,
						}, true)
						ctx.Items["personae"][i] = commons.CtxEntry{
							Dnode: dnode,
						}
					}
				} else {
					dnode := createDiagramNode(&commons.Node{
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
					dnode := createDiagramNode(&commons.Node{
						Type: i,
						Name: i,
					}, true)
					ctx.Items["partners"][i] = commons.CtxEntry{
						Dnode: dnode,
					}
				} else {
					dnode := createDiagramNode(&commons.Node{
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

	kept_links := map[string]commons.LinkModel{}

	for _, ll := range *m.Links {
		a, b, md, c, d, e := explodeLink(ll, &ctx, expandeds, collapseds, m)
		kk := c + md + d
		if v, found := kept_links[kk]; found {
			v.Count = kept_links[kk].Count + 1
			if len(e) == 0 {
				kept_links[kk].Labels[e] = true
			} else {
				kept_links[kk] = commons.LinkModel{From: a, To: b, Count: 1, Labels: map[string]bool{}, Mode: md}
				if len(e) > 0 {
					kept_links[kk].Labels[e] = true
				}
			}
		}
	}

	for l := range kept_links {
		a := l.From
		b := l.To
		cc := l.Count
		md := l.Mode
		if len(a) > 0 && len(b) > 0 {
			if "-" == md {
				if cc > 1 {
					d.Connect(a, b) // "(" + string(cc) + ")"
				} else {
					d.Connect(a, b)
				}
			} else {
				if ">" == md {
					if cc > 1 {
						d.Connect(a, b) // "(" + string(cc) + ")"
					} else {
						d.Connect(a, b)
					}
				} else {
					if "<" == md {
						if cc > 1 {
							d.Connect(a, b) // "(" + string(cc) + ")"
						} else {
							d.Connect(a, b)
						}
					} else {
						if cc > 1 {
							d.Connect(a, b) // "(" + string(cc) + ")"
						} else {
							d.Connect(a, b)
						}
					}
				}
			}
		}
	}

	return nil
}
