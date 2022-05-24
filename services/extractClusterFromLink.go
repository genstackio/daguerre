package services

import "github.com/genstackio/daguerre/commons"

func extractClusterFromLink(l commons.LinkModel, m *commons.Model) (string, string, string, string) {
	a := ""
	b := ""
	c := ""
	d := ""

	if items, found := m.Lists[l.From.Type]; found {
		if item, found2 := items[l.From.Name]; found2 {
			a = item.Cluster
			c = l.From.Name
		}
	}
	if items, found := m.Lists[l.To.Type]; found {
		if item, found2 := items[l.To.Name]; found2 {
			b = item.Cluster
			d = l.To.Name
		}
	}

	return a, b, c, d
}
