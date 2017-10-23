package utility

type Callhome struct {
	Callhomestatus       interface{} `json:"callhomestatus,omitempty"`
	Emailaddress         string      `json:"emailaddress,omitempty"`
	Flashfirstfail       string      `json:"flashfirstfail,omitempty"`
	Flashlatestfailure   string      `json:"flashlatestfailure,omitempty"`
	Hddfirstfail         string      `json:"hddfirstfail,omitempty"`
	Hddlatestfailure     string      `json:"hddlatestfailure,omitempty"`
	Ipaddress            string      `json:"ipaddress,omitempty"`
	Port                 int         `json:"port,omitempty"`
	Powfirstfail         string      `json:"powfirstfail,omitempty"`
	Powlatestfailure     string      `json:"powlatestfailure,omitempty"`
	Proxymode            string      `json:"proxymode,omitempty"`
	Restartlatestfail    string      `json:"restartlatestfail,omitempty"`
	Sslcardfirstfailure  string      `json:"sslcardfirstfailure,omitempty"`
	Sslcardlatestfailure string      `json:"sslcardlatestfailure,omitempty"`
}
