package ssl

type Sslservicesslcertkeybinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Servicename   string `json:"servicename,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
}
