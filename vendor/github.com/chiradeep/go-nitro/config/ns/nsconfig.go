package ns

type Nsconfig struct {
	Cip                     string      `json:"cip,omitempty"`
	Cipheader               string      `json:"cipheader,omitempty"`
	Config                  string      `json:"config,omitempty"`
	Config1                 string      `json:"config1,omitempty"`
	Config2                 string      `json:"config2,omitempty"`
	Configchanged           bool        `json:"configchanged,omitempty"`
	Cookieversion           string      `json:"cookieversion,omitempty"`
	Crportrange             string      `json:"crportrange,omitempty"`
	Currentsytemtime        string      `json:"currentsytemtime,omitempty"`
	Exclusivequotamaxclient int         `json:"exclusivequotamaxclient,omitempty"`
	Exclusivequotaspillover int         `json:"exclusivequotaspillover,omitempty"`
	Flags                   int         `json:"flags,omitempty"`
	Force                   bool        `json:"force,omitempty"`
	Ftpportrange            string      `json:"ftpportrange,omitempty"`
	Grantquotamaxclient     int         `json:"grantquotamaxclient,omitempty"`
	Grantquotaspillover     int         `json:"grantquotaspillover,omitempty"`
	Httpport                interface{} `json:"httpport,omitempty"`
	Ifnum                   interface{} `json:"ifnum,omitempty"`
	Ignoredevicespecific    bool        `json:"ignoredevicespecific,omitempty"`
	Ipaddress               string      `json:"ipaddress,omitempty"`
	Lastconfigchangedtime   string      `json:"lastconfigchangedtime,omitempty"`
	Lastconfigsavetime      string      `json:"lastconfigsavetime,omitempty"`
	Level                   string      `json:"level,omitempty"`
	Mappedip                string      `json:"mappedip,omitempty"`
	Maxconn                 int         `json:"maxconn,omitempty"`
	Maxreq                  int         `json:"maxreq,omitempty"`
	Message                 string      `json:"message,omitempty"`
	Netmask                 string      `json:"netmask,omitempty"`
	Nsvlan                  int         `json:"nsvlan,omitempty"`
	Outtype                 string      `json:"outtype,omitempty"`
	Pmtumin                 int         `json:"pmtumin,omitempty"`
	Pmtutimeout             int         `json:"pmtutimeout,omitempty"`
	Primaryip               string      `json:"primaryip,omitempty"`
	Primaryip6              string      `json:"primaryip6,omitempty"`
	Range                   int         `json:"range,omitempty"`
	Rbaconfig               string      `json:"rbaconfig,omitempty"`
	Response                string      `json:"response,omitempty"`
	Securecookie            string      `json:"securecookie,omitempty"`
	Systemtime              int         `json:"systemtime,omitempty"`
	Systemtype              string      `json:"systemtype,omitempty"`
	Tagged                  string      `json:"tagged,omitempty"`
	Template                bool        `json:"template,omitempty"`
	Timezone                string      `json:"timezone,omitempty"`
	Weakpassword            bool        `json:"weakpassword,omitempty"`
}
