package services

import (
	"github.com/genstackio/daguerre/commons"
)

func buildModelFromConfig(c *commons.Config) (*commons.Model, error) {
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
		err := analyzeLayer(&v, &m, c)
		if nil != err {
			return nil, err
		}
	}

	// link
	for k, v := range c.Layers {
		if len(v.Name) == 0 {
			v.Name = k
		}
		err := linkLayer(&v, &m, c)
		if nil != err {
			return nil, err
		}
	}

	return &m, nil
}
