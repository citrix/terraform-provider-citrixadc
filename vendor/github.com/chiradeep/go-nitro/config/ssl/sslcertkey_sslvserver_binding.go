package ssl

type Sslcertkeysslvserverbinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkey     string `json:"certkey,omitempty"`
	Data        int    `json:"data,omitempty"`
	Servername  string `json:"servername,omitempty"`
	Version     int    `json:"version,omitempty"`
	Vserver     bool   `json:"vserver,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
