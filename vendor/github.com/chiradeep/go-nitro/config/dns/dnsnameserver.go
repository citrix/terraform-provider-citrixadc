package dns

type Dnsnameserver struct {
	Clmonowner      int    `json:"clmonowner,omitempty"`
	Clmonview       int    `json:"clmonview,omitempty"`
	Dnsprofilename  string `json:"dnsprofilename,omitempty"`
	Dnsvservername  string `json:"dnsvservername,omitempty"`
	Ip              string `json:"ip,omitempty"`
	Local           bool   `json:"local,omitempty"`
	Nameserverstate string `json:"nameserverstate,omitempty"`
	Port            int    `json:"port,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
	State           string `json:"state,omitempty"`
	Type            string `json:"type,omitempty"`
}
