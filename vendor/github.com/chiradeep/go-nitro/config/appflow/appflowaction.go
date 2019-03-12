package appflow

type Appflowaction struct {
	Clientsidemeasurements string      `json:"clientsidemeasurements,omitempty"`
	Collectors             interface{} `json:"collectors,omitempty"`
	Comment                string      `json:"comment,omitempty"`
	Description            string      `json:"description,omitempty"`
	Distributionalgorithm  string      `json:"distributionalgorithm,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Metricslog             bool        `json:"metricslog,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Newname                string      `json:"newname,omitempty"`
	Pagetracking           string      `json:"pagetracking,omitempty"`
	Referencecount         int         `json:"referencecount,omitempty"`
	Securityinsight        string      `json:"securityinsight,omitempty"`
	Transactionlog         string      `json:"transactionlog,omitempty"`
	Videoanalytics         string      `json:"videoanalytics,omitempty"`
	Webinsight             string      `json:"webinsight,omitempty"`
}
