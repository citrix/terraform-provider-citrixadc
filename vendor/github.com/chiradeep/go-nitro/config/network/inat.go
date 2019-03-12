package network

type Inat struct {
	Flags        int    `json:"flags,omitempty"`
	Ftp          string `json:"ftp,omitempty"`
	Mode         string `json:"mode,omitempty"`
	Name         string `json:"name,omitempty"`
	Privateip    string `json:"privateip,omitempty"`
	Proxyip      string `json:"proxyip,omitempty"`
	Publicip     string `json:"publicip,omitempty"`
	Tcpproxy     string `json:"tcpproxy,omitempty"`
	Td           int    `json:"td,omitempty"`
	Tftp         string `json:"tftp,omitempty"`
	Useproxyport string `json:"useproxyport,omitempty"`
	Usip         string `json:"usip,omitempty"`
	Usnip        string `json:"usnip,omitempty"`
}
