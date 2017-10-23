package ssl

type Sslvserversslcertkeybinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Vservername   string `json:"vservername,omitempty"`
}
