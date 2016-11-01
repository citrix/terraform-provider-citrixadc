package lb

type Lbparameter struct {
  Consolidatedlconn string `json:"consolidatedlconn,omitempty"`
  Httponlycookieflag string `json:"httponlycookieflag,omitempty"`
  Monitorconnectionclose string `json:"monitorconnectionclose,omitempty"`
  Monitorskipmaxclient string `json:"monitorskipmaxclient,omitempty"`
  Preferdirectroute string `json:"preferdirectroute,omitempty"`
  Sessionsthreshold int `json:"sessionsthreshold,omitempty"`
  Startuprrfactor int `json:"startuprrfactor,omitempty"`
  Useportforhashlb string `json:"useportforhashlb,omitempty"`
  Vserverspecificmac string `json:"vserverspecificmac,omitempty"`
}
