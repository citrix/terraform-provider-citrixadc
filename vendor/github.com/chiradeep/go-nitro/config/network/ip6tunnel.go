package network

type Ip6tunnel struct {
	Encapip  string `json:"encapip,omitempty"`
	Local    string `json:"local,omitempty"`
	Name     string `json:"name,omitempty"`
	Remote   string `json:"remote,omitempty"`
	Remoteip string `json:"remoteip,omitempty"`
	Type     int    `json:"type,omitempty"`
}
