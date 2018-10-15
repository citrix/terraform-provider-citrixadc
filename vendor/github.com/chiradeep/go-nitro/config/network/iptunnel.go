package network

type Iptunnel struct {
	Channel           int         `json:"channel,omitempty"`
	Encapip           string      `json:"encapip,omitempty"`
	Grepayload        string      `json:"grepayload,omitempty"`
	Ipsecprofilename  string      `json:"ipsecprofilename,omitempty"`
	Ipsectunnelstatus string      `json:"ipsectunnelstatus,omitempty"`
	Local             string      `json:"local,omitempty"`
	Name              string      `json:"name,omitempty"`
	Ownergroup        string      `json:"ownergroup,omitempty"`
	Protocol          string      `json:"protocol,omitempty"`
	Refcnt            int         `json:"refcnt,omitempty"`
	Remote            string      `json:"remote,omitempty"`
	Remotesubnetmask  string      `json:"remotesubnetmask,omitempty"`
	Sysname           string      `json:"sysname,omitempty"`
	Tunneltype        interface{} `json:"tunneltype,omitempty"`
	Type              int         `json:"type,omitempty"`
	Vlan              int         `json:"vlan,omitempty"`
}
