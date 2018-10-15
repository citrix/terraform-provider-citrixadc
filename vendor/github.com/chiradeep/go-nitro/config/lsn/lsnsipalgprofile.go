package lsn

type Lsnsipalgprofile struct {
	Datasessionidletimeout int    `json:"datasessionidletimeout,omitempty"`
	Opencontactpinhole     string `json:"opencontactpinhole,omitempty"`
	Openrecordroutepinhole string `json:"openrecordroutepinhole,omitempty"`
	Openregisterpinhole    string `json:"openregisterpinhole,omitempty"`
	Openroutepinhole       string `json:"openroutepinhole,omitempty"`
	Openviapinhole         string `json:"openviapinhole,omitempty"`
	Registrationtimeout    int    `json:"registrationtimeout,omitempty"`
	Rport                  string `json:"rport,omitempty"`
	Sipalgprofilename      string `json:"sipalgprofilename,omitempty"`
	Sipdstportrange        string `json:"sipdstportrange,omitempty"`
	Sipsessiontimeout      int    `json:"sipsessiontimeout,omitempty"`
	Sipsrcportrange        string `json:"sipsrcportrange,omitempty"`
	Siptransportprotocol   string `json:"siptransportprotocol,omitempty"`
}
