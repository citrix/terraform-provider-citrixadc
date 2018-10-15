package ns

type Nsacl struct {
	Aclaction      string      `json:"aclaction,omitempty"`
	Aclassociate   interface{} `json:"aclassociate,omitempty"`
	Aclname        string      `json:"aclname,omitempty"`
	Destip         bool        `json:"destip,omitempty"`
	Destipop       string      `json:"destipop,omitempty"`
	Destipval      string      `json:"destipval,omitempty"`
	Destport       bool        `json:"destport,omitempty"`
	Destportop     string      `json:"destportop,omitempty"`
	Destportval    string      `json:"destportval,omitempty"`
	Dfdhash        string      `json:"dfdhash,omitempty"`
	Established    bool        `json:"established,omitempty"`
	Hits           int         `json:"hits,omitempty"`
	Icmpcode       int         `json:"icmpcode,omitempty"`
	Icmptype       int         `json:"icmptype,omitempty"`
	Interface      string      `json:"Interface,omitempty"`
	Kernelstate    string      `json:"kernelstate,omitempty"`
	Logstate       string      `json:"logstate,omitempty"`
	Newname        string      `json:"newname,omitempty"`
	Priority       int         `json:"priority,omitempty"`
	Protocol       string      `json:"protocol,omitempty"`
	Protocolnumber int         `json:"protocolnumber,omitempty"`
	Ratelimit      int         `json:"ratelimit,omitempty"`
	Srcip          bool        `json:"srcip,omitempty"`
	Srcipop        string      `json:"srcipop,omitempty"`
	Srcipval       string      `json:"srcipval,omitempty"`
	Srcmac         string      `json:"srcmac,omitempty"`
	Srcmacmask     string      `json:"srcmacmask,omitempty"`
	Srcport        bool        `json:"srcport,omitempty"`
	Srcportop      string      `json:"srcportop,omitempty"`
	Srcportval     string      `json:"srcportval,omitempty"`
	State          string      `json:"state,omitempty"`
	Stateful       string      `json:"stateful,omitempty"`
	Td             int         `json:"td,omitempty"`
	Ttl            int         `json:"ttl,omitempty"`
	Type           string      `json:"type,omitempty"`
	Vlan           int         `json:"vlan,omitempty"`
	Vxlan          int         `json:"vxlan,omitempty"`
}
