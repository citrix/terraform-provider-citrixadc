package system

type Systemcounters struct {
	Countergroup string `json:"countergroup,omitempty"`
	Datasource   string `json:"datasource,omitempty"`
	Response     string `json:"response,omitempty"`
}
