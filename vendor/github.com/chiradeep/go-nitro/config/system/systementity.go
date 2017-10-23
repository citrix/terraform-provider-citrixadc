package system

type Systementity struct {
	Core       int    `json:"core,omitempty"`
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
	Type       string `json:"type,omitempty"`
}
