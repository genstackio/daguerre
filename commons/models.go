package commons

type ClusterModel struct {
	Name     string                `json:"name"`
	Nodes    *[]Node               `json:"nodes"`
	Features map[string]bool       `json:"features"`
	Params   map[string]ParamModel `json:"params"`
}
type ItemModel struct {
	Type      string                `json:"type"`
	Name      string                `json:"name"`
	Layer     *LayerConfig          `json:"layer"`
	LayerType *LayerConfig          `json:"layerType"`
	Cluster   string                `json:"cluster"`
	Features  map[string]bool       `json:"features"`
	Params    map[string]ParamModel `json:"params"`
}
type LinkModel struct {
	From   LinkEndpointModel `json:"from"`
	To     LinkEndpointModel `json:"to"`
	Mode   string            `json:"mode"`
	Labels map[string]bool   `json:"labels"`
	Label  string            `json:"label"`
}

func (l LinkModel) ToString() string {
	return l.From.Type + "/" + l.From.Name + " " + l.Mode + " " + l.To.Type + "/" + l.To.Name
}

type KeptLinkModel struct {
	From   *CtxEntry       `json:"from"`
	To     *CtxEntry       `json:"to"`
	Mode   string          `json:"mode"`
	Count  int             `json:"count"`
	Labels map[string]bool `json:"labels"`
	Label  string          `json:"label"`
}

type LinkEndpointModel struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type ParamModel struct {
	Type       string  `json:"type"`
	Value      string  `json:"value"`
	IntValue   int64   `json:"intValue"`
	FloatValue float64 `json:"floatValue"`
	BoolValue  bool    `json:"boolValue"`
}
