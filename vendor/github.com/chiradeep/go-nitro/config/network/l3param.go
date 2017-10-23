package network

type L3param struct {
	Acllogtime           int    `json:"acllogtime,omitempty"`
	Dropdfflag           string `json:"dropdfflag,omitempty"`
	Dropipfragments      string `json:"dropipfragments,omitempty"`
	Externalloopback     string `json:"externalloopback,omitempty"`
	Forwardicmpfragments string `json:"forwardicmpfragments,omitempty"`
	Icmpgenratethreshold int    `json:"icmpgenratethreshold,omitempty"`
	Miproundrobin        string `json:"miproundrobin,omitempty"`
	Overridernat         string `json:"overridernat,omitempty"`
	Srcnat               string `json:"srcnat,omitempty"`
	Tnlpmtuwoconn        string `json:"tnlpmtuwoconn,omitempty"`
	Usipserverstraypkt   string `json:"usipserverstraypkt,omitempty"`
}
