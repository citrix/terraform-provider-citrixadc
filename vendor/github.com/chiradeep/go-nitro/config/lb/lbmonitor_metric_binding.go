package lb

type Lbmonitormetricbinding struct {
	Metric          string `json:"metric,omitempty"`
	Metrictable     string `json:"metrictable,omitempty"`
	Metricthreshold int    `json:"metricthreshold,omitempty"`
	Metricunit      string `json:"metric_unit,omitempty"`
	Metricweight    int    `json:"metricweight,omitempty"`
	Monitorname     string `json:"monitorname,omitempty"`
}
