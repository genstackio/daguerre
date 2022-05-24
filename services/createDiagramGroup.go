package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func createDiagramGroup(ctx *commons.Ctx, g *diagram.Group, name string, m *commons.Model) *diagram.Group {
	gg := g.NewGroup(name).Label(name)

	//g.Group(gg)

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
			return gg.Add(createDiagramNode(ctx, &commons.Node{
				Type: vn,
				Name: name,
			}, false))
		}
		return gg.Add(createDiagramNode(ctx, &commons.Node{
			Type: "cluster_item",
			Name: name,
		}, false))
	}
	return gg.Add(createDiagramNode(ctx, &commons.Node{
		Type: "cluster_item",
		Name: name,
	}, false))
}
