package ssl

type Sslservicegroupsslcertkeybinding struct {
	Ca               bool   `json:"ca,omitempty"`
	Certkeyname      string `json:"certkeyname,omitempty"`
	Crlcheck         string `json:"crlcheck,omitempty"`
	Ocspcheck        string `json:"ocspcheck,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Snicert          bool   `json:"snicert,omitempty"`
}
