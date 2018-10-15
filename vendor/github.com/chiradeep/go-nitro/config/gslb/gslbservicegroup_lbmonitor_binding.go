package gslb

type Gslbservicegrouplbmonitorbinding struct {
	Hashid           int    `json:"hashid,omitempty"`
	Monitorname      string `json:"monitor_name,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Monweight        int    `json:"monweight,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	Publicip         string `json:"publicip,omitempty"`
	Publicport       int    `json:"publicport,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Siteprefix       string `json:"siteprefix,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
