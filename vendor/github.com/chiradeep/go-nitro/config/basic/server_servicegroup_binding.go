package basic

type Serverservicegroupbinding struct {
	Dupport          int    `json:"dup_port,omitempty"`
	Dupsvctype       string `json:"dup_svctype,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrcfgflags      int    `json:"svrcfgflags,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
}
