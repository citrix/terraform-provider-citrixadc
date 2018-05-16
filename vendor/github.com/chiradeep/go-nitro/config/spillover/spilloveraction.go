package spillover

type Spilloveraction struct {
	Action  string      `json:"action,omitempty"`
	Builtin interface{} `json:"builtin,omitempty"`
	Name    string      `json:"name,omitempty"`
	Newname string      `json:"newname,omitempty"`
}
