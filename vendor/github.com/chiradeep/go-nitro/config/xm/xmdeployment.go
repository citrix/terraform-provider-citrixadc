package xm

type Xmdeployment struct {
	Config      string `json:"config,omitempty"`
	Frompackage string `json:"frompackage,omitempty"`
	Meta        string `json:"meta,omitempty"`
	Name        string `json:"name,omitempty"`
}
