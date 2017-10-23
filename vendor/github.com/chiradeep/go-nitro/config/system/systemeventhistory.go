package system

type Systemeventhistory struct {
	Datasource string `json:"datasource,omitempty"`
	Endtime    string `json:"endtime,omitempty"`
	Last       int    `json:"last,omitempty"`
	Response   string `json:"response,omitempty"`
	Starttime  string `json:"starttime,omitempty"`
	Unit       string `json:"unit,omitempty"`
}
