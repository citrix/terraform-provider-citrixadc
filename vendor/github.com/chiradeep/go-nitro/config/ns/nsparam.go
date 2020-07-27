package ns

type Nsparam struct {
	Advancedanalyticsstats    string      `json:"advancedanalyticsstats,omitempty"`
	Aftpallowrandomsourceport string      `json:"aftpallowrandomsourceport,omitempty"`
	Cip                       string      `json:"cip,omitempty"`
	Cipheader                 string      `json:"cipheader,omitempty"`
	Cookieversion             string      `json:"cookieversion,omitempty"`
	Crportrange               string      `json:"crportrange,omitempty"`
	Exclusivequotamaxclient   int         `json:"exclusivequotamaxclient"` // 0 is a valid value
	Exclusivequotaspillover   int         `json:"exclusivequotaspillover"` // 0 is a valid value
	Ftpportrange              string      `json:"ftpportrange,omitempty"`
	Grantquotamaxclient       int         `json:"grantquotamaxclient"` // 0 is a valid value
	Grantquotaspillover       int         `json:"grantquotaspillover"` // 0 is a valid value
	Httpport                  interface{} `json:"httpport,omitempty"`
	Icaports                  interface{} `json:"icaports,omitempty"`
	Internaluserlogin         string      `json:"internaluserlogin,omitempty"`
	Maxconn                   int         `json:"maxconn"` // 0 is a valid value
	Maxreq                    int         `json:"maxreq"`  // 0 is a valid value
	Mgmthttpport              int         `json:"mgmthttpport,omitempty"`
	Mgmthttpsport             int         `json:"mgmthttpsport,omitempty"`
	Pmtumin                   int         `json:"pmtumin,omitempty"`
	Pmtutimeout               int         `json:"pmtutimeout,omitempty"`
	Proxyprotocol             string      `json:"proxyprotocol,omitempty"`
	Securecookie              string      `json:"securecookie,omitempty"`
	Secureicaports            interface{} `json:"secureicaports,omitempty"`
	Servicepathingressvlan    int         `json:"servicepathingressvlan,omitempty"`
	Tcpcip                    string      `json:"tcpcip,omitempty"`
	Timezone                  string      `json:"timezone,omitempty"`
	Useproxyport              string      `json:"useproxyport,omitempty"`
}
