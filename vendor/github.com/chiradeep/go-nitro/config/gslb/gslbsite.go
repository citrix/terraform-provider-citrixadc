package gslb

type Gslbsite struct {
	Metricexchange       string `json:"metricexchange,omitempty"`
	Nwmetricexchange     string `json:"nwmetricexchange,omitempty"`
	Parentsite           string `json:"parentsite,omitempty"`
	Persistencemepstatus string `json:"persistencemepstatus,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Sessionexchange      string `json:"sessionexchange,omitempty"`
	Siteipaddress        string `json:"siteipaddress,omitempty"`
	Sitename             string `json:"sitename,omitempty"`
	Sitetype             string `json:"sitetype,omitempty"`
	Status               string `json:"status,omitempty"`
	Triggermonitor       string `json:"triggermonitor,omitempty"`
	Version              int    `json:"version,omitempty"`
}
