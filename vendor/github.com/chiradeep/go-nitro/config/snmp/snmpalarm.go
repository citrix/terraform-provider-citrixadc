package snmp

type Snmpalarm struct {
	Logging        string `json:"logging,omitempty"`
	Normalvalue    int    `json:"normalvalue,omitempty"`
	Severity       string `json:"severity,omitempty"`
	State          string `json:"state,omitempty"`
	Thresholdvalue int    `json:"thresholdvalue,omitempty"`
	Time           int    `json:"time,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	Trapname       string `json:"trapname,omitempty"`
}
