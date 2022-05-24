package commons

type LayerConfig struct {
	Apigws       map[string]ApigwConfig      `json:"apigws"`
	Buckets      map[string]BucketConfig     `json:"buckets"`
	Buses        map[string]BusConfig        `json:"buses"`
	Cdns         map[string]CdnConfig        `json:"cdns"`
	Cloudfuncs   map[string]CloudfuncConfig  `json:"cloudfuncs"`
	Edges        map[string]EdgeConfig       `json:"edges"`
	Hidden       bool                        `json:"hidden"`
	Lambdas      map[string]LambdaConfig     `json:"lambdas"`
	Links        *[]string                   `json:"links"`
	Name         string                      `json:"name"`
	Opensearches map[string]OpensearchConfig `json:"opensearches"`
	Partners     map[string]PartnerConfig    `json:"partners"`
	Personae     map[string]PersonaConfig    `json:"personae"`
	Queues       map[string]QueueConfig      `json:"queues"`
	Records      map[string]RecordConfig     `json:"records"`
	Ses          map[string]SesConfig        `json:"ses"`
	Tables       map[string]TableConfig      `json:"tables"`
	Topics       map[string]TopicConfig      `json:"topics"`
	Type         string                      `json:"type"`
	Zones        map[string]ZoneConfig       `json:"zones"`
}

type LinksAware interface {
	GetLinks() *[]string
}

type ApigwConfig struct {
	Links *[]string `json:"links"`
}

func (c ApigwConfig) GetLinks() *[]string {
	return c.Links
}

type BucketConfig struct {
	Links *[]string `json:"links"`
}

func (c BucketConfig) GetLinks() *[]string {
	return c.Links
}

type BusConfig struct {
	Links *[]string `json:"links"`
}

func (c BusConfig) GetLinks() *[]string {
	return c.Links
}

type CdnConfig struct {
	Links *[]string `json:"links"`
}

func (c CdnConfig) GetLinks() *[]string {
	return c.Links
}

type CloudfuncConfig struct {
	Links *[]string `json:"links"`
}

func (c CloudfuncConfig) GetLinks() *[]string {
	return c.Links
}

type EdgeConfig struct {
	Links *[]string `json:"links"`
}

func (c EdgeConfig) GetLinks() *[]string {
	return c.Links
}

type LambdaConfig struct {
	Links *[]string `json:"links"`
}

func (c LambdaConfig) GetLinks() *[]string {
	return c.Links
}

type OpensearchConfig struct {
	Links *[]string `json:"links"`
}

func (c OpensearchConfig) GetLinks() *[]string {
	return c.Links
}

type PartnerConfig struct {
	Links *[]string `json:"links"`
}

func (c PartnerConfig) GetLinks() *[]string {
	return c.Links
}

type PersonaConfig struct {
	Links *[]string `json:"links"`
}

func (c PersonaConfig) GetLinks() *[]string {
	return c.Links
}

type QueueConfig struct {
	Links *[]string `json:"links"`
}

func (c QueueConfig) GetLinks() *[]string {
	return c.Links
}

type RecordConfig struct {
	Links *[]string `json:"links"`
}

func (c RecordConfig) GetLinks() *[]string {
	return c.Links
}

type SesConfig struct {
	Links *[]string `json:"links"`
}

func (c SesConfig) GetLinks() *[]string {
	return c.Links
}

type TableConfig struct {
	Links *[]string `json:"links"`
}

func (c TableConfig) GetLinks() *[]string {
	return c.Links
}

type TopicConfig struct {
	Links *[]string `json:"links"`
}

func (c TopicConfig) GetLinks() *[]string {
	return c.Links
}

type ZoneConfig struct {
	Links *[]string `json:"links"`
}

func (c ZoneConfig) GetLinks() *[]string {
	return c.Links
}
