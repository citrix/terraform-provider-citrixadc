package ns

type Nsip struct {
	Arp              string      `json:"arp,omitempty"`
	Arpresponse      string      `json:"arpresponse,omitempty"`
	Bgp              string      `json:"bgp,omitempty"`
	Dynamicrouting   string      `json:"dynamicrouting,omitempty"`
	Flags            int         `json:"flags,omitempty"`
	Freeports        int         `json:"freeports,omitempty"`
	Ftp              string      `json:"ftp,omitempty"`
	Gui              string      `json:"gui,omitempty"`
	Hostroute        string      `json:"hostroute,omitempty"`
	Hostrtgw         string      `json:"hostrtgw,omitempty"`
	Hostrtgwact      string      `json:"hostrtgwact,omitempty"`
	Icmp             string      `json:"icmp,omitempty"`
	Icmpresponse     string      `json:"icmpresponse,omitempty"`
	Ipaddress        string      `json:"ipaddress,omitempty"`
	Iptype           interface{} `json:"iptype,omitempty"`
	Metric           int         `json:"metric,omitempty"`
	Mgmtaccess       string      `json:"mgmtaccess,omitempty"`
	Netmask          string      `json:"netmask,omitempty"`
	Ospf             string      `json:"ospf,omitempty"`
	Ospfarea         int         `json:"ospfarea,omitempty"`
	Ospfareaval      int         `json:"ospfareaval,omitempty"`
	Ospflsatype      string      `json:"ospflsatype,omitempty"`
	Ownernode        int         `json:"ownernode,omitempty"`
	Restrictaccess   string      `json:"restrictaccess,omitempty"`
	Rip              string      `json:"rip,omitempty"`
	Snmp             string      `json:"snmp,omitempty"`
	Ssh              string      `json:"ssh,omitempty"`
	State            string      `json:"state,omitempty"`
	Td               int         `json:"td,omitempty"`
	Telnet           string      `json:"telnet,omitempty"`
	Type             string      `json:"type,omitempty"`
	Viprtadv2bsd     bool        `json:"viprtadv2bsd,omitempty"`
	Vipvsercount     int         `json:"vipvsercount,omitempty"`
	Vipvserdowncount int         `json:"vipvserdowncount,omitempty"`
	Vrid             int         `json:"vrid,omitempty"`
	Vserver          string      `json:"vserver,omitempty"`
	Vserverrhilevel  string      `json:"vserverrhilevel,omitempty"`
}
