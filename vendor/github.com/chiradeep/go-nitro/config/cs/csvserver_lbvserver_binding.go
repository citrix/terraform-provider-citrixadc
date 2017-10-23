package cs

type Csvserverlbvserverbinding struct {
	Hits          int    `json:"hits,omitempty"`
	Lbvserver     string `json:"lbvserver,omitempty"`
	Name          string `json:"name,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
}
