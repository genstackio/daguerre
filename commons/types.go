package commons

import "github.com/blushft/go-diagrams/diagram"

type Order struct {
	Schema   string   `json:"schema"`
	Input    string   `json:"input"`
	Format   string   `json:"format"`
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
	Schemas    map[string]SchemaConfig  `json:"schemas"`
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
	Clusters      map[string]CtxEntry                       `json:"clusters"`
	Items         map[string]map[string]CtxEntry            `json:"items"`
	NodeTypes     map[string]PluginNodeType                 `json:"nodeTypes"`
	NodeListNames map[string]struct{ NodeTypeIndex string } `json:"nodeListNames"`
	Aliases       map[string]string                         `json:"aliases"`
	CustomAssets  map[string]bool                           `json:"customAssets"`
}

func (c Ctx) RegisterCustomAsset(key string) {
	c.CustomAssets[key] = true
}
func (c Ctx) RegisterPluginNodeTypes(nodeTypes *[]PluginNodeType) {
	for _, nt := range *nodeTypes {
		c.NodeTypes[nt.Type] = nt
		c.NodeListNames[nt.ListName] = struct{ NodeTypeIndex string }{NodeTypeIndex: nt.Type}
	}
}

func (c Ctx) RegisterNodeTypesAliases(nodeTypes *[]PluginNodeType) {
	for _, nt := range *nodeTypes {
		c.RegisterAliases(nt.Aliases)
	}
}

func (c Ctx) RegisterAliases(aliases map[string]string) {
	if nil == aliases {
		return
	}
	for k, v := range aliases {
		c.Aliases[k] = v
	}
}
func (c Ctx) GetNodeTypeByNodeListName(name string) (*PluginNodeType, bool) {
	ntn, found1 := c.NodeListNames[name]
	if !found1 {
		return nil, false
	}
	x, found := c.NodeTypes[ntn.NodeTypeIndex]

	if !found {
		return nil, false
	}

	return &x, true
}
func (c Ctx) GetNodeTypeName(name string) (string, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return "", false
	}
	return x.Name, found
}

func (c Ctx) GetNodeTypeType(name string) (string, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return "", false
	}
	return x.Type, found
}

func (c Ctx) GetDiagramNodeCreator(name string) (PluginNodeTypeDiagramNodeCreator, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return nil, false
	}
	return x.DiagramNodeCreator, nil != x.DiagramNodeCreator
}

func (c Ctx) GetDiagramNodeLabeller(name string) (PluginNodeTypeDiagramNodeLabeller, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return nil, false
	}
	return x.DiagramNodeLabeller, nil != x.DiagramNodeLabeller
}

func (c Ctx) GetNodeCreator(name string) (PluginNodeTypeNodeCreator, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return nil, false
	}
	return x.NodeCreator, nil != x.NodeCreator
}

func (c Ctx) GetLinkEndpointAligner(name string) (PluginNodeTypeLinkEndpointAligner, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return nil, false
	}
	return x.LinkEndpointAligner, nil != x.LinkEndpointAligner
}

func (c Ctx) GetLinkPopulator(name string) (PluginNodeTypeLinkPopulator, bool) {
	x, found := c.GetNodeTypeByNodeListName(name)
	if !found {
		return nil, false
	}
	return x.LinkPopulator, nil != x.LinkPopulator
}

type Node struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Hidden  bool   `json:"hidden"`
	Points  int    `json:"points"`
	Variant string `json:"variant"`
}

type MapperFunc func(n string, v interface{}, lt *LayerConfig, l *LayerConfig) interface{}

type PluginNodeTypeNodeCreator func(ctx *Ctx, lt *LayerConfig, l *LayerConfig, m *Model)
type PluginNodeTypeDiagramNodeCreator func(ctx *Ctx, n *Node, label string) *diagram.Node
type PluginNodeTypeLinkPopulator func(ctx *Ctx, lt *LayerConfig, l *LayerConfig, m *Model)
type PluginNodeTypeLinkEndpointAligner func(ctx *Ctx, le LinkEndpointModel, m *Model) LinkEndpointModel
type PluginNodeTypeDiagramNodeLabeller func(ctx *Ctx, n *Node, mode string) string

type PluginNodeType struct {
	Name                string
	Type                string
	ListName            string
	NodeCreator         PluginNodeTypeNodeCreator
	DiagramNodeCreator  PluginNodeTypeDiagramNodeCreator
	DiagramNodeLabeller PluginNodeTypeDiagramNodeLabeller
	LinkPopulator       PluginNodeTypeLinkPopulator
	LinkEndpointAligner PluginNodeTypeLinkEndpointAligner
	Aliases             map[string]string
}

type Plugin interface {
	Register(ctx *Ctx)
}
