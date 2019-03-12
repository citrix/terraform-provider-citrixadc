package ns

type Nsip6 struct {
	Curstate          string      `json:"curstate,omitempty"`
	Decrementhoplimit string      `json:"decrementhoplimit,omitempty"`
	Dynamicrouting    string      `json:"dynamicrouting,omitempty"`
	Ftp               string      `json:"ftp,omitempty"`
	Gui               string      `json:"gui,omitempty"`
	Hostroute         string      `json:"hostroute,omitempty"`
	Icmp              string      `json:"icmp,omitempty"`
	Ip6hostrtgw       string      `json:"ip6hostrtgw,omitempty"`
	Iptype            interface{} `json:"iptype,omitempty"`
	Ipv6address       string      `json:"ipv6address,omitempty"`
	Map               string      `json:"map,omitempty"`
	Metric            int         `json:"metric,omitempty"`
	Mgmtaccess        string      `json:"mgmtaccess,omitempty"`
	Nd                string      `json:"nd,omitempty"`
	Networkroute      string      `json:"networkroute,omitempty"`
	Ospf6lsatype      string      `json:"ospf6lsatype,omitempty"`
	Ospfarea          int         `json:"ospfarea,omitempty"`
	Ownerdownresponse string      `json:"ownerdownresponse,omitempty"`
	Ownernode         int         `json:"ownernode,omitempty"`
	Restrictaccess    string      `json:"restrictaccess,omitempty"`
	Scope             string      `json:"scope,omitempty"`
	Snmp              string      `json:"snmp,omitempty"`
	Ssh               string      `json:"ssh,omitempty"`
	State             string      `json:"state,omitempty"`
	Systemtype        string      `json:"systemtype,omitempty"`
	Tag               int         `json:"tag,omitempty"`
	Td                int         `json:"td,omitempty"`
	Telnet            string      `json:"telnet,omitempty"`
	Type              string      `json:"type,omitempty"`
	Viprtadv2bsd      bool        `json:"viprtadv2bsd,omitempty"`
	Vipvsercount      int         `json:"vipvsercount,omitempty"`
	Vipvserdowncount  int         `json:"vipvserdowncount,omitempty"`
	Vlan              int         `json:"vlan,omitempty"`
	Vrid6             int         `json:"vrid6,omitempty"`
	Vserver           string      `json:"vserver,omitempty"`
	Vserverrhilevel   string      `json:"vserverrhilevel,omitempty"`
}
