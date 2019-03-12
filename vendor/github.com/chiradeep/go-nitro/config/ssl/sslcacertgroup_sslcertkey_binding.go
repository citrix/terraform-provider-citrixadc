package ssl

type Sslcacertgroupsslcertkeybinding struct {
	Cacertgroupname string `json:"cacertgroupname,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
	Crlcheck        string `json:"crlcheck,omitempty"`
	Ocspcheck       string `json:"ocspcheck,omitempty"`
}
