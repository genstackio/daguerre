package utils

import "github.com/genstackio/daguerre/commons"

func PopulateNodes[T interface{}](t string, name string, hidden bool, x map[string]T, lt *commons.LayerConfig, l *commons.LayerConfig, m *commons.Model) {
	if nil != x {
		for k := range x {
			if "{{name}}" == k {
				k = name
			}
			nod := commons.Node{Type: t, Name: k, Hidden: hidden}
			if !hidden {
				if _, found := m.Lists[t][k]; !found {
					m.Lists[t][k] = commons.ItemModel{
						Name:      k,
						LayerType: lt,
						Layer:     l,
						Type:      t,
						Features:  map[string]bool{},
						Params:    map[string]commons.ParamModel{},
						Cluster:   l.Name,
					}
					*m.Clusters[l.Name].Nodes = append(*m.Clusters[l.Name].Nodes, nod)
				}

			}
		}
	}
}
