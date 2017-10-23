package gslb

type Gslbdomainlbmonitorbinding struct {
	Customheaders              string `json:"customheaders,omitempty"`
	Httprequest                string `json:"httprequest,omitempty"`
	Iptunnel                   string `json:"iptunnel,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitorname                string `json:"monitorname,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Name                       string `json:"name,omitempty"`
	Respcode                   string `json:"respcode,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Vservername                string `json:"vservername,omitempty"`
}
