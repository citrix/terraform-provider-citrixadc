package ns

type Nsparam struct {
	Advancedanalyticsstats    string      `json:"advancedanalyticsstats,omitempty"`
	Aftpallowrandomsourceport string      `json:"aftpallowrandomsourceport,omitempty"`
	Cip                       string      `json:"cip,omitempty"`
	Cipheader                 string      `json:"cipheader,omitempty"`
	Cookieversion             string      `json:"cookieversion,omitempty"`
	Crportrange               string      `json:"crportrange,omitempty"`
	Exclusivequotamaxclient   int         `json:"exclusivequotamaxclient,omitempty"`
	Exclusivequotaspillover   int         `json:"exclusivequotaspillover,omitempty"`
	Ftpportrange              string      `json:"ftpportrange,omitempty"`
	Grantquotamaxclient       int         `json:"grantquotamaxclient,omitempty"`
	Grantquotaspillover       int         `json:"grantquotaspillover,omitempty"`
	Httpport                  interface{} `json:"httpport,omitempty"`
	Icaports                  interface{} `json:"icaports,omitempty"`
	Internaluserlogin         string      `json:"internaluserlogin,omitempty"`
	Maxconn                   int         `json:"maxconn,omitempty"`
	Maxreq                    int         `json:"maxreq,omitempty"`
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
