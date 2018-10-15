package rdp

type Rdpconnections struct {
	All          bool   `json:"all,omitempty"`
	Endpointip   string `json:"endpointip,omitempty"`
	Endpointport int    `json:"endpointport,omitempty"`
	Peid         int    `json:"peid,omitempty"`
	Targetip     string `json:"targetip,omitempty"`
	Targetport   int    `json:"targetport,omitempty"`
	Username     string `json:"username,omitempty"`
}
