package vpn

type Vpnicaconnection struct {
	All        bool   `json:"all,omitempty"`
	Destip     string `json:"destip,omitempty"`
	Destport   int    `json:"destport,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Nodeid     int    `json:"nodeid,omitempty"`
	Peid       int    `json:"peid,omitempty"`
	Srcip      string `json:"srcip,omitempty"`
	Srcport    int    `json:"srcport,omitempty"`
	Transproto string `json:"transproto,omitempty"`
	Username   string `json:"username,omitempty"`
}
