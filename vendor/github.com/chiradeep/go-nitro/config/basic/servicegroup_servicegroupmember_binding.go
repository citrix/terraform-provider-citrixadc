package basic

type Servicegroupservicegroupmemberbinding struct {
	Customserverid            string `json:"customserverid,omitempty"`
	Dbsttl                    int    `json:"dbsttl,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Nameserver                string `json:"nameserver,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Serverid                  int    `json:"serverid,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Svcitmpriority            int    `json:"svcitmpriority,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Trofsreason               string `json:"trofsreason,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}
