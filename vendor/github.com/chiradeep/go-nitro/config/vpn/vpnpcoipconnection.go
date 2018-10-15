package vpn

type Vpnpcoipconnection struct {
	All      bool   `json:"all,omitempty"`
	Destip   string `json:"destip,omitempty"`
	Destport int    `json:"destport,omitempty"`
	Nodeid   int    `json:"nodeid,omitempty"`
	Peid     int    `json:"peid,omitempty"`
	Srcip    string `json:"srcip,omitempty"`
	Srcport  int    `json:"srcport,omitempty"`
	Username string `json:"username,omitempty"`
}
