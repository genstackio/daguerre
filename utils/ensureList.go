package utils

import "github.com/genstackio/daguerre/commons"

func EnsureList(t string, m *commons.Model) string {
	if nil == m.Lists {
		m.Lists = map[string]map[string]commons.ItemModel{}
	}
	if _, found := m.Lists[t]; !found {
		m.Lists[t] = map[string]commons.ItemModel{}
	}
	return t
}
