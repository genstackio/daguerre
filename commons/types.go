package commons

import "github.com/blushft/go-diagrams/diagram"

type Order struct {
	Input    string   `json:"input"`
	Show     []string `json:"show"`
	Expand   []string `json:"expand"`
	Collapse []string `json:"collapse"`
	Clusters []string `json:"clusters"`
	Personae []string `json:"personae"`
	Partners []string `json:"partners"`
	Output   string   `json:"output"`
}

type Config struct {
	Name       string                   `json:"name"`
	Direction  string                   `json:"direction"`
	Styles     map[string]string        `json:"styles"`
	Layers     map[string]LayerConfig   `json:"layers"`
	Personae   map[string]PersonaConfig `json:"personae"`
	Partners   map[string]PartnerConfig `json:"partners"`
	Buses      map[string]BusConfig     `json:"buses"`
	LayerTypes map[string]LayerConfig   `json:"layerTypes"`
}

type Model struct {
	Name      string                          `json:"name"`
	Direction string                          `json:"direction"`
	Styles    map[string]string               `json:"styles"`
	Clusters  map[string]ClusterModel         `json:"clusters"`
	Links     *[]LinkModel                    `json:"links"`
	Lists     map[string]map[string]ItemModel `json:"lists"`
}

type CtxEntry struct {
	Dnode  *diagram.Node  `json:"dnode"`
	Dgroup *diagram.Group `json:"dgroup"`
}

type Ctx struct {
	Clusters map[string]CtxEntry            `json:"clusters"`
	Items    map[string]map[string]CtxEntry `json:"items"`
}

type Node struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Points int    `json:"points"`
}

type MapperFunc func(n string, v interface{}, lt *LayerConfig, l *LayerConfig) interface{}
