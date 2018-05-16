package system

type Systemglobaldata struct {
	Core         int    `json:"core,omitempty"`
	Countergroup string `json:"countergroup,omitempty"`
	Counters     string `json:"counters,omitempty"`
	Datasource   string `json:"datasource,omitempty"`
	Endtime      string `json:"endtime,omitempty"`
	Last         int    `json:"last,omitempty"`
	Lastupdate   int    `json:"lastupdate,omitempty"`
	Response     string `json:"response,omitempty"`
	Starttime    string `json:"starttime,omitempty"`
	Startupdate  int    `json:"startupdate,omitempty"`
	Unit         string `json:"unit,omitempty"`
}
