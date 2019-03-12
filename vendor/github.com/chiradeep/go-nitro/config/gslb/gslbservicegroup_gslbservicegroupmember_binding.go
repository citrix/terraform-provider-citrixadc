package gslb

type Gslbservicegroupgslbservicegroupmemberbinding struct {
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Gslbthreshold             int    `json:"gslbthreshold,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Preferredlocation         string `json:"preferredlocation,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int    `json:"publicport,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Siteprefix                string `json:"siteprefix,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}
