package system

type Systementitydata struct {
	Alldeleted  bool   `json:"alldeleted,omitempty"`
	Allinactive bool   `json:"allinactive,omitempty"`
	Core        int    `json:"core,omitempty"`
	Counters    string `json:"counters,omitempty"`
	Datasource  string `json:"datasource,omitempty"`
	Endtime     string `json:"endtime,omitempty"`
	Last        int    `json:"last,omitempty"`
	Lastupdate  int    `json:"lastupdate,omitempty"`
	Name        string `json:"name,omitempty"`
	Response    string `json:"response,omitempty"`
	Starttime   string `json:"starttime,omitempty"`
	Startupdate int    `json:"startupdate,omitempty"`
	Type        string `json:"type,omitempty"`
	Unit        string `json:"unit,omitempty"`
}
