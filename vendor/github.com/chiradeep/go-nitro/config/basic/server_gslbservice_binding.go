package basic

type Servergslbservicebinding struct {
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
}
