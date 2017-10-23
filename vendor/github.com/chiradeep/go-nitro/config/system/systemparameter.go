package system

type Systemparameter struct {
	Natpcbforceflushlimit int    `json:"natpcbforceflushlimit,omitempty"`
	Natpcbrstontimeout    string `json:"natpcbrstontimeout,omitempty"`
	Promptstring          string `json:"promptstring,omitempty"`
	Rbaonresponse         string `json:"rbaonresponse,omitempty"`
	Timeout               int    `json:"timeout,omitempty"`
}
