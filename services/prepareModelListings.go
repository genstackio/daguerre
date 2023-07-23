package services

import "github.com/genstackio/daguerre/commons"

func prepareModelListings(ctx *commons.Ctx, m *commons.Model, o *commons.Order) *commons.ModelListings {
	l := commons.ModelListings{
		Expandeds:  map[string]bool{},
		Collapseds: map[string]bool{},
		Clusters:   map[string]bool{},
		Requireds:  map[string]bool{},
		NeedLinks:  map[string]bool{},
		RequiredSpecials: map[string]map[string]bool{
			"personae": {},
			"partners": {},
		},
	}

	for _, v := range o.Collapse {
		l.Collapseds[v] = true
		l.NeedLinks[v] = true
	}
	for _, v := range o.Clusters {
		l.Requireds[v] = true
		l.NeedLinks[v] = true
	}
	for _, v := range o.Personae {
		l.RequiredSpecials["personae"][v] = true
	}
	for _, v := range o.Partners {
		l.RequiredSpecials["partners"][v] = true
	}
	for _, v := range o.Show {
		_, found1 := l.Expandeds[v]
		_, found2 := l.Collapseds["all"]
		if !found1 && !found2 {
			l.Collapseds[v] = true
		}
	}
	for _, v := range o.Expand {
		if _, found := l.Collapseds[v]; !found {
			if _, found2 := l.Collapseds["all"]; !found2 {
				l.Expandeds[v] = true
				l.NeedLinks[v] = true
			}
		}
	}

	if nil != m.Links {
		for _, ll := range *m.Links {
			a, b, c, d := extractClusterFromLink(ll, m)
			_, found := l.NeedLinks[a]
			_, found2 := l.NeedLinks[b]
			_, found3 := l.NeedLinks["all"]
			if found || found2 || found3 {
				if "personae" == a || "partners" == a {
					l.RequiredSpecials[a][c] = true
				} else if len(a) > 0 {
					l.Requireds[a] = true
				}
				if "personae" == b || "partners" == b {
					l.RequiredSpecials[b][d] = true
				} else if len(b) > 0 {
					l.Requireds[b] = true
				}
			}
		}
	}

	for kk, vv := range m.Clusters {
		if "personae" == kk || "partners" == kk {
			continue
		}
		_, found1 := l.Expandeds[kk]
		_, found2 := l.Expandeds["all"]
		_, found11 := l.Collapseds[kk]
		_, found21 := l.Collapseds["all"]
		if (found1 || found2) && !(found11 || found21) {
			for _, vvv := range vv.Requires {
				l.Requireds[vvv] = true
				l.NeedLinks[vvv] = true
			}
		}
		_, found3 := l.Collapseds[kk]
		_, found4 := l.Collapseds["all"]
		if found3 || found4 {
			for _, vvv := range vv.Requires {
				l.Requireds[vvv] = true
				l.NeedLinks[vvv] = true
			}
		}
	}

	return &l
}
