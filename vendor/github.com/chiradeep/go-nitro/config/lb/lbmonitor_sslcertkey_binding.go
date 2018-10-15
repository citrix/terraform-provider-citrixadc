package lb

type Lbmonitorsslcertkeybinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkeyname string `json:"certkeyname,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
}
