package network

type L3param struct {
	Acllogtime           int    `json:"acllogtime,omitempty"`
	Allowclasseipv4      string `json:"allowclasseipv4,omitempty"`
	Dropdfflag           string `json:"dropdfflag,omitempty"`
	Dropipfragments      string `json:"dropipfragments,omitempty"`
	Dynamicrouting       string `json:"dynamicrouting,omitempty"`
	Externalloopback     string `json:"externalloopback,omitempty"`
	Forwardicmpfragments string `json:"forwardicmpfragments,omitempty"`
	Icmpgenratethreshold int    `json:"icmpgenratethreshold,omitempty"`
	Implicitaclallow     string `json:"implicitaclallow,omitempty"`
	Ipv6dynamicrouting   string `json:"ipv6dynamicrouting,omitempty"`
	Miproundrobin        string `json:"miproundrobin,omitempty"`
	Overridernat         string `json:"overridernat,omitempty"`
	Srcnat               string `json:"srcnat,omitempty"`
	Tnlpmtuwoconn        string `json:"tnlpmtuwoconn,omitempty"`
	Usipserverstraypkt   string `json:"usipserverstraypkt,omitempty"`
}
