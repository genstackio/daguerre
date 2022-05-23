package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func createDiagramGroup(g *diagram.Group, name string, m *commons.Model) *diagram.Group {
	gg := diagram.NewGroup(name).Label(name)

	g.Group(gg)

	if _, found := m.Clusters[name]; found {
		points := map[string]int{}
		for _, nn := range *m.Clusters[name].Nodes {
			points[nn.Type] = nn.Points
		}
		v := 0
		vn := ""
		for kk, nn := range points {
			if nn > v {
				v = nn
				vn = kk
			}
		}
		if len(vn) > 0 {
			return gg.Add(createDiagramNode(&commons.Node{
				Type: vn,
				Name: name,
			}, true))
		}
		return gg.Add(createDiagramNode(&commons.Node{
			Type: "cluster_item",
			Name: name,
		}, true))
	}
	return gg.Add(createDiagramNode(&commons.Node{
		Type: "cluster_item",
		Name: name,
	}, true))
}
