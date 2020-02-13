package lb

type Lbmetrictable struct {
	Builtin     interface{} `json:"builtin,omitempty"`
	Feature     string      `json:"feature,omitempty"`
	Metric      string      `json:"metric,omitempty"`
	Metrictable string      `json:"metrictable,omitempty"`
	Metrictype  string      `json:"metrictype,omitempty"`
	Snmpoid     string      `json:"Snmpoid,omitempty"`
}
