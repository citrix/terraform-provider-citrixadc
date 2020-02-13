package basic

type Servicegroupservicegroupentitymonbindingsbinding struct {
	Customserverid             string `json:"customserverid,omitempty"`
	Dbsttl                     int    `json:"dbsttl,omitempty"`
	Hashid                     int    `json:"hashid,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitorname                string `json:"monitor_name,omitempty"`
	Monitorstate               string `json:"monitor_state,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Nameserver                 string `json:"nameserver,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Serverid                   int    `json:"serverid,omitempty"`
	Servicegroupentname2       string `json:"servicegroupentname2,omitempty"`
	Servicegroupname           string `json:"servicegroupname,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}
