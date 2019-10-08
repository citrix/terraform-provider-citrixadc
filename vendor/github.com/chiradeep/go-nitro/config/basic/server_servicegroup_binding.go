package basic

type Serverservicegroupbinding struct {
	Appflowlog           string `json:"appflowlog,omitempty"`
	Boundtd              int    `json:"boundtd,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	Cip                  string `json:"cip,omitempty"`
	Cipheader            string `json:"cipheader,omitempty"`
	Cka                  string `json:"cka,omitempty"`
	Clttimeout           int    `json:"clttimeout,omitempty"`
	Cmp                  string `json:"cmp,omitempty"`
	Customserverid       string `json:"customserverid,omitempty"`
	Downstateflush       string `json:"downstateflush,omitempty"`
	Dupport              int    `json:"dup_port,omitempty"`
	Dupsvctype           string `json:"dup_svctype,omitempty"`
	Maxbandwidth         int    `json:"maxbandwidth,omitempty"`
	Maxclient            int    `json:"maxclient,omitempty"`
	Maxreq               int    `json:"maxreq,omitempty"`
	Monthreshold         int    `json:"monthreshold,omitempty"`
	Name                 string `json:"name,omitempty"`
	Port                 int    `json:"port,omitempty"`
	Sc                   string `json:"sc,omitempty"`
	Servicegroupentname2 string `json:"servicegroupentname2,omitempty"`
	Servicegroupname     string `json:"servicegroupname,omitempty"`
	Serviceipaddress     string `json:"serviceipaddress,omitempty"`
	Serviceipstr         string `json:"serviceipstr,omitempty"`
	Sp                   string `json:"sp,omitempty"`
	Svcitmactsvcs        int    `json:"svcitmactsvcs,omitempty"`
	Svcitmboundsvcs      int    `json:"svcitmboundsvcs,omitempty"`
	Svcitmpriority       int    `json:"svcitmpriority,omitempty"`
	Svctype              string `json:"svctype,omitempty"`
	Svrcfgflags          int    `json:"svrcfgflags,omitempty"`
	Svrstate             string `json:"svrstate,omitempty"`
	Svrtimeout           int    `json:"svrtimeout,omitempty"`
	Tcpb                 string `json:"tcpb,omitempty"`
	Usip                 string `json:"usip,omitempty"`
	Weight               int    `json:"weight,omitempty"`
}
