package lb

type Lbparameter struct {
	Allowboundsvcremoval          string      `json:"allowboundsvcremoval,omitempty"`
	Builtin                       interface{} `json:"builtin,omitempty"`
	Consolidatedlconn             string      `json:"consolidatedlconn,omitempty"`
	Cookiepassphrase              string      `json:"cookiepassphrase,omitempty"`
	Httponlycookieflag            string      `json:"httponlycookieflag,omitempty"`
	Monitorconnectionclose        string      `json:"monitorconnectionclose,omitempty"`
	Monitorskipmaxclient          string      `json:"monitorskipmaxclient,omitempty"`
	Preferdirectroute             string      `json:"preferdirectroute,omitempty"`
	Retainservicestate            string      `json:"retainservicestate,omitempty"`
	Sessionsthreshold             int         `json:"sessionsthreshold,omitempty"`
	Startuprrfactor               int         `json:"startuprrfactor,omitempty"`
	Useencryptedpersistencecookie string      `json:"useencryptedpersistencecookie,omitempty"`
	Useportforhashlb              string      `json:"useportforhashlb,omitempty"`
	Usesecuredpersistencecookie   string      `json:"usesecuredpersistencecookie,omitempty"`
	Vserverspecificmac            string      `json:"vserverspecificmac,omitempty"`
}
