package user

type Uservserver struct {
	Comment                   string `json:"comment,omitempty"`
	Curstate                  string `json:"curstate,omitempty"`
	Defaultlb                 string `json:"defaultlb,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Name                      string `json:"name,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Params                    string `json:"Params,omitempty"`
	Port                      int    `json:"port,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimemsec       int    `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Userprotocol              string `json:"userprotocol,omitempty"`
	Value                     string `json:"value,omitempty"`
}
