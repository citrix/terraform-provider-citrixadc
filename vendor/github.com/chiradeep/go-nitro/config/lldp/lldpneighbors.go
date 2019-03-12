package lldp

type Lldpneighbors struct {
	Autonegadvertised  string `json:"autonegadvertised,omitempty"`
	Autonegenabled     string `json:"autonegenabled,omitempty"`
	Autonegmautype     string `json:"autonegmautype,omitempty"`
	Autonegsupport     string `json:"autonegsupport,omitempty"`
	Chassisid          string `json:"chassisid,omitempty"`
	Chassisidsubtype   string `json:"chassisidsubtype,omitempty"`
	Flag               int    `json:"flag,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Ifnumber           int    `json:"ifnumber,omitempty"`
	Iftype             string `json:"iftype,omitempty"`
	Linkaggrcapable    string `json:"linkaggrcapable,omitempty"`
	Linkaggrenabled    string `json:"linkaggrenabled,omitempty"`
	Linkaggrid         int    `json:"linkaggrid,omitempty"`
	Mgmtaddress        string `json:"mgmtaddress,omitempty"`
	Mgmtaddresssubtype string `json:"mgmtaddresssubtype,omitempty"`
	Mtu                int    `json:"mtu,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Portdescription    string `json:"portdescription,omitempty"`
	Portid             string `json:"portid,omitempty"`
	Portidsubtype      string `json:"portidsubtype,omitempty"`
	Portprotoenabled   int    `json:"portprotoenabled,omitempty"`
	Portprotoid        int    `json:"portprotoid,omitempty"`
	Portprotosupported int    `json:"portprotosupported,omitempty"`
	Portvlanid         int    `json:"portvlanid,omitempty"`
	Protocolid         string `json:"protocolid,omitempty"`
	Sys                string `json:"sys,omitempty"`
	Syscapabilities    string `json:"syscapabilities,omitempty"`
	Syscapenabled      string `json:"syscapenabled,omitempty"`
	Sysdesc            string `json:"sysdesc,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Vlanid             int    `json:"vlanid,omitempty"`
}
