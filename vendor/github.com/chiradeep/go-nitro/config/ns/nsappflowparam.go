package ns

type Nsappflowparam struct {
	Clienttrafficonly string `json:"clienttrafficonly,omitempty"`
	Httpcookie        string `json:"httpcookie,omitempty"`
	Httphost          string `json:"httphost,omitempty"`
	Httpmethod        string `json:"httpmethod,omitempty"`
	Httpreferer       string `json:"httpreferer,omitempty"`
	Httpurl           string `json:"httpurl,omitempty"`
	Httpuseragent     string `json:"httpuseragent,omitempty"`
	Templaterefresh   int    `json:"templaterefresh,omitempty"`
	Udppmtu           int    `json:"udppmtu,omitempty"`
}
