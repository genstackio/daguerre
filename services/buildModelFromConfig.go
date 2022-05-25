package services

import (
	"github.com/genstackio/daguerre/commons"
	"github.com/genstackio/daguerre/utils"
)

func buildModelFromConfig(ctx *commons.Ctx, c *commons.Config) (*commons.Model, error) {
	m := commons.Model{
		Name:      c.Name,
		Direction: c.Direction,
		Styles:    c.Styles,
	}

	personaeLayer := commons.LayerConfig{Name: "personae", Personae: map[string]commons.PersonaConfig{}, Type: "custom"}
	partnersLayer := commons.LayerConfig{Name: "partners", Partners: map[string]commons.PartnerConfig{}, Type: "custom"}
	busesLayer := commons.LayerConfig{Name: "buses", Buses: map[string]commons.BusConfig{"default": {}}, Type: "custom"}

	for k, v := range c.Personae {
		personaeLayer.Personae[k] = v
	}
	for k, v := range c.Partners {
		partnersLayer.Partners[k] = v
	}
	for k, v := range c.Buses {
		busesLayer.Buses[k] = v
	}

	c.Layers["personae"] = personaeLayer
	c.Layers["partners"] = partnersLayer
	c.Layers["buses"] = busesLayer

	// initialize
	for k, v := range c.Layers {
		if len(v.Name) == 0 {
			v.Name = k
		}
	}

	// analyze
	for k, v := range c.Layers {
		if len(v.Name) == 0 {
			v.Name = k
		}

		lt, err := getLayerTypeConfig(&v, c)

		if nil != err {
			return nil, err
		}

		if nil == m.Clusters {
			m.Clusters = map[string]commons.ClusterModel{}
		}
		cm := commons.ClusterModel{
			Nodes:    &[]commons.Node{},
			Requires: []string{},
		}

		if len(v.Requires) > 0 {
			reqs := map[string]bool{}
			for _, rq := range v.Requires {
				reqs[rq] = true
			}
			reqList := []string{}
			for k := range reqs {
				reqList = append(reqList, k)
			}
			cm.Requires = reqList
		}

		m.Clusters[v.Name] = cm

		for _, nt := range ctx.NodeTypes {
			nt.NodeCreator(ctx, lt, &v, &m)
		}

	}

	// link
	for k, v := range c.Layers {
		if len(v.Name) == 0 {
			v.Name = k
		}

		lt, err := getLayerTypeConfig(&v, c)

		if nil != err {
			return nil, err
		}

		vars := map[string]string{}

		utils.PopulateLinks(ctx, lt.Links, &m, vars, "")
		utils.PopulateLinks(ctx, v.Links, &m, vars, "")

		for _, nt := range ctx.NodeTypes {
			nt.LinkPopulator(ctx, lt, &v, &m)
		}

		if nil != err {
			return nil, err
		}
	}

	return &m, nil
}
