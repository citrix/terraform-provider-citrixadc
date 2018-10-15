package gslb

type Gslbservicelbmonitorbinding struct {
	Failedprobes               int    `json:"failedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitorname                string `json:"monitor_name,omitempty"`
	Monitorstate               string `json:"monitor_state,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monstatparam1              int    `json:"monstatparam1,omitempty"`
	Monstatparam2              int    `json:"monstatparam2,omitempty"`
	Monstatparam3              int    `json:"monstatparam3,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Totalfailedprobes          int    `json:"totalfailedprobes,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}
