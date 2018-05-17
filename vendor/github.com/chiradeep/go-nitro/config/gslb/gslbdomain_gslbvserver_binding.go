package gslb

type Gslbdomaingslbvserverbinding struct {
	Backuplbmethod     string `json:"backuplbmethod,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Customheaders      string `json:"customheaders,omitempty"`
	Dnsrecordtype      string `json:"dnsrecordtype,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Edr                string `json:"edr,omitempty"`
	Lbmethod           string `json:"lbmethod,omitempty"`
	Mir                string `json:"mir,omitempty"`
	Name               string `json:"name,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Persistenceid      int    `json:"persistenceid,omitempty"`
	Persistencetype    string `json:"persistencetype,omitempty"`
	Persistmask        string `json:"persistmask,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sitename           string `json:"sitename,omitempty"`
	Sitepersistence    string `json:"sitepersistence,omitempty"`
	Siteprefix         string `json:"siteprefix,omitempty"`
	State              string `json:"state,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	V6netmasklen       int    `json:"v6netmasklen,omitempty"`
	V6persistmasklen   int    `json:"v6persistmasklen,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
}
