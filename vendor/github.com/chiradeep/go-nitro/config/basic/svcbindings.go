package basic

type Svcbindings struct {
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Svrstate    string `json:"svrstate,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
