package ns

type Nsacl6 struct {
	Acl6action     string `json:"acl6action,omitempty"`
	Acl6name       string `json:"acl6name,omitempty"`
	Aclaction      string `json:"aclaction,omitempty"`
	Destipop       string `json:"destipop,omitempty"`
	Destipv6       bool   `json:"destipv6,omitempty"`
	Destipv6val    string `json:"destipv6val,omitempty"`
	Destport       bool   `json:"destport,omitempty"`
	Destportop     string `json:"destportop,omitempty"`
	Destportval    string `json:"destportval,omitempty"`
	Established    bool   `json:"established,omitempty"`
	Hits           int    `json:"hits,omitempty"`
	Icmpcode       int    `json:"icmpcode,omitempty"`
	Icmptype       int    `json:"icmptype,omitempty"`
	Interface      string `json:"Interface,omitempty"`
	Kernelstate    string `json:"kernelstate,omitempty"`
	Newname        string `json:"newname,omitempty"`
	Priority       int    `json:"priority,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Protocolnumber int    `json:"protocolnumber,omitempty"`
	Srcipop        string `json:"srcipop,omitempty"`
	Srcipv6        bool   `json:"srcipv6,omitempty"`
	Srcipv6val     string `json:"srcipv6val,omitempty"`
	Srcmac         string `json:"srcmac,omitempty"`
	Srcport        bool   `json:"srcport,omitempty"`
	Srcportop      string `json:"srcportop,omitempty"`
	Srcportval     string `json:"srcportval,omitempty"`
	State          string `json:"state,omitempty"`
	Td             int    `json:"td,omitempty"`
	Ttl            int    `json:"ttl,omitempty"`
	Vlan           int    `json:"vlan,omitempty"`
}
