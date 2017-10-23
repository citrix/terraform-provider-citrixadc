package aaa

type Aaasession struct {
	All         bool   `json:"all,omitempty"`
	Destip      string `json:"destip,omitempty"`
	Destport    int    `json:"destport,omitempty"`
	Groupname   string `json:"groupname,omitempty"`
	Iip         string `json:"iip,omitempty"`
	Intranetip  string `json:"intranetip,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Peid        int    `json:"peid,omitempty"`
	Port        int    `json:"port,omitempty"`
	Privateip   string `json:"privateip,omitempty"`
	Privateport int    `json:"privateport,omitempty"`
	Publicip    string `json:"publicip,omitempty"`
	Publicport  int    `json:"publicport,omitempty"`
	Username    string `json:"username,omitempty"`
}
