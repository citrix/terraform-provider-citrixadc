package basic

type Servergslbservicegroupbinding struct {
	Appflowlog       string `json:"appflowlog,omitempty"`
	Boundtd          int    `json:"boundtd,omitempty"`
	Cip              string `json:"cip,omitempty"`
	Cipheader        string `json:"cipheader,omitempty"`
	Clttimeout       int    `json:"clttimeout,omitempty"`
	Customserverid   string `json:"customserverid,omitempty"`
	Downstateflush   string `json:"downstateflush,omitempty"`
	Dupport          int    `json:"dup_port,omitempty"`
	Dupsvctype       string `json:"dup_svctype,omitempty"`
	Maxbandwidth     int    `json:"maxbandwidth,omitempty"`
	Maxclient        int    `json:"maxclient,omitempty"`
	Maxreq           int    `json:"maxreq,omitempty"`
	Monthreshold     int    `json:"monthreshold,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrcfgflags      int    `json:"svrcfgflags,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
	Svrtimeout       int    `json:"svrtimeout,omitempty"`
}
