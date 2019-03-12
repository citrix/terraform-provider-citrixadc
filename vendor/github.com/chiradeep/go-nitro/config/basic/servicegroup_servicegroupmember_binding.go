package basic

type Servicegroupservicegroupmemberbinding struct {
	Customserverid            string `json:"customserverid,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Serverid                  int    `json:"serverid,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}
