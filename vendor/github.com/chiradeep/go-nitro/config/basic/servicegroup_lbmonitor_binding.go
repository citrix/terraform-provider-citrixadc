package basic

type Servicegrouplbmonitorbinding struct {
	Customserverid   string `json:"customserverid,omitempty"`
	Dbsttl           int    `json:"dbsttl,omitempty"`
	Hashid           int    `json:"hashid,omitempty"`
	Monitorname      string `json:"monitor_name,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Monweight        int    `json:"monweight,omitempty"`
	Nameserver       string `json:"nameserver,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serverid         int    `json:"serverid,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}
