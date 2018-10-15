package gslb

type Gslbsite struct {
	Backupparentlist       interface{} `json:"backupparentlist,omitempty"`
	Clip                   string      `json:"clip,omitempty"`
	Curbackupparentip      string      `json:"curbackupparentip,omitempty"`
	Metricexchange         string      `json:"metricexchange,omitempty"`
	Naptrreplacementsuffix string      `json:"naptrreplacementsuffix,omitempty"`
	Nwmetricexchange       string      `json:"nwmetricexchange,omitempty"`
	Parentsite             string      `json:"parentsite,omitempty"`
	Persistencemepstatus   string      `json:"persistencemepstatus,omitempty"`
	Publicclip             string      `json:"publicclip,omitempty"`
	Publicip               string      `json:"publicip,omitempty"`
	Sessionexchange        string      `json:"sessionexchange,omitempty"`
	Siteipaddress          string      `json:"siteipaddress,omitempty"`
	Sitename               string      `json:"sitename,omitempty"`
	Sitetype               string      `json:"sitetype,omitempty"`
	Status                 string      `json:"status,omitempty"`
	Triggermonitor         string      `json:"triggermonitor,omitempty"`
	Version                int         `json:"version,omitempty"`
}
