package ns

type Nspbr6 struct {
	Action            string `json:"action,omitempty"`
	Curstate          int    `json:"curstate,omitempty"`
	Data              bool   `json:"data,omitempty"`
	Destipop          string `json:"destipop,omitempty"`
	Destipv6          bool   `json:"destipv6,omitempty"`
	Destipv6val       string `json:"destipv6val,omitempty"`
	Destport          bool   `json:"destport,omitempty"`
	Destportop        string `json:"destportop,omitempty"`
	Destportval       string `json:"destportval,omitempty"`
	Detail            bool   `json:"detail,omitempty"`
	Failedprobes      int    `json:"failedprobes,omitempty"`
	Hits              int    `json:"hits,omitempty"`
	Interface         string `json:"Interface,omitempty"`
	Iptunnel          string `json:"iptunnel,omitempty"`
	Kernelstate       string `json:"kernelstate,omitempty"`
	Monitor           string `json:"monitor,omitempty"`
	Monstatcode       int    `json:"monstatcode,omitempty"`
	Monstatparam1     int    `json:"monstatparam1,omitempty"`
	Monstatparam2     int    `json:"monstatparam2,omitempty"`
	Monstatparam3     int    `json:"monstatparam3,omitempty"`
	Msr               string `json:"msr,omitempty"`
	Name              string `json:"name,omitempty"`
	Nexthop           bool   `json:"nexthop,omitempty"`
	Nexthopval        string `json:"nexthopval,omitempty"`
	Nexthopvlan       int    `json:"nexthopvlan,omitempty"`
	Ownergroup        string `json:"ownergroup,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Protocolnumber    int    `json:"protocolnumber,omitempty"`
	Srcipop           string `json:"srcipop,omitempty"`
	Srcipv6           bool   `json:"srcipv6,omitempty"`
	Srcipv6val        string `json:"srcipv6val,omitempty"`
	Srcmac            string `json:"srcmac,omitempty"`
	Srcmacmask        string `json:"srcmacmask,omitempty"`
	Srcport           bool   `json:"srcport,omitempty"`
	Srcportop         string `json:"srcportop,omitempty"`
	Srcportval        string `json:"srcportval,omitempty"`
	State             string `json:"state,omitempty"`
	Td                int    `json:"td,omitempty"`
	Totalfailedprobes int    `json:"totalfailedprobes,omitempty"`
	Totalprobes       int    `json:"totalprobes,omitempty"`
	Vlan              int    `json:"vlan,omitempty"`
	Vxlan             int    `json:"vxlan,omitempty"`
	Vxlanvlanmap      string `json:"vxlanvlanmap,omitempty"`
}
