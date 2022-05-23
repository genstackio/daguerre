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

type ApigwConfig struct {
}
type BucketConfig struct {
}
type BusConfig struct {
}
type CdnConfig struct {
}
type CloudfuncConfig struct {
}
type EdgeConfig struct {
}
type LambdaConfig struct {
}
type OpensearchConfig struct {
}
type PartnerConfig struct {
}
type PersonaConfig struct {
}
type QueueConfig struct {
}
type RecordConfig struct {
}
type SesConfig struct {
}
type TableConfig struct {
}
type TopicConfig struct {
}
type ZoneConfig struct {
}
