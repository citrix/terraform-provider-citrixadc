package utility

type Callhome struct {
	Callhomestatus       interface{} `json:"callhomestatus,omitempty"`
	Emailaddress         string      `json:"emailaddress,omitempty"`
	Flashfirstfail       string      `json:"flashfirstfail,omitempty"`
	Flashlatestfailure   string      `json:"flashlatestfailure,omitempty"`
	Hbcustominterval     int         `json:"hbcustominterval,omitempty"`
	Hddfirstfail         string      `json:"hddfirstfail,omitempty"`
	Hddlatestfailure     string      `json:"hddlatestfailure,omitempty"`
	Ipaddress            string      `json:"ipaddress,omitempty"`
	Memthrefirstanomaly  string      `json:"memthrefirstanomaly,omitempty"`
	Memthrelatestanomaly string      `json:"memthrelatestanomaly,omitempty"`
	Mode                 string      `json:"mode,omitempty"`
	Nodeid               int         `json:"nodeid,omitempty"`
	Port                 int         `json:"port,omitempty"`
	Powfirstfail         string      `json:"powfirstfail,omitempty"`
	Powlatestfailure     string      `json:"powlatestfailure,omitempty"`
	Proxyauthservice     string      `json:"proxyauthservice,omitempty"`
	Proxymode            string      `json:"proxymode,omitempty"`
	Restartlatestfail    string      `json:"restartlatestfail,omitempty"`
	Rlfirsthighdrop      string      `json:"rlfirsthighdrop,omitempty"`
	Rllatesthighdrop     string      `json:"rllatesthighdrop,omitempty"`
	Sslcardfirstfailure  string      `json:"sslcardfirstfailure,omitempty"`
	Sslcardlatestfailure string      `json:"sslcardlatestfailure,omitempty"`
}
