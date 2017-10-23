package network

type Nd6ravariables struct {
	Ceaserouteradv           string `json:"ceaserouteradv,omitempty"`
	Currhoplimit             int    `json:"currhoplimit,omitempty"`
	Defaultlifetime          int    `json:"defaultlifetime,omitempty"`
	Lastrtadvtime            int    `json:"lastrtadvtime,omitempty"`
	Linkmtu                  int    `json:"linkmtu,omitempty"`
	Managedaddrconfig        string `json:"managedaddrconfig,omitempty"`
	Maxrtadvinterval         int    `json:"maxrtadvinterval,omitempty"`
	Minrtadvinterval         int    `json:"minrtadvinterval,omitempty"`
	Nextrtadvdelay           int    `json:"nextrtadvdelay,omitempty"`
	Onlyunicastrtadvresponse string `json:"onlyunicastrtadvresponse,omitempty"`
	Otheraddrconfig          string `json:"otheraddrconfig,omitempty"`
	Reachabletime            int    `json:"reachabletime,omitempty"`
	Retranstime              int    `json:"retranstime,omitempty"`
	Sendrouteradv            string `json:"sendrouteradv,omitempty"`
	Srclinklayeraddroption   string `json:"srclinklayeraddroption,omitempty"`
	Vlan                     int    `json:"vlan,omitempty"`
}
