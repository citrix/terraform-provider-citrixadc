package lb

type Lbmetrictable struct {
	Metric      string `json:"metric,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
	Metrictype  string `json:"metrictype,omitempty"`
	Snmpoid     string `json:"Snmpoid,omitempty"`
}
