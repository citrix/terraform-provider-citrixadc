package cmp

type Cmpparameter struct {
	Addvaryheader    string `json:"addvaryheader,omitempty"`
	Cmpbypasspct     int    `json:"cmpbypasspct,omitempty"`
	Cmplevel         string `json:"cmplevel,omitempty"`
	Cmponpush        string `json:"cmponpush,omitempty"`
	Externalcache    string `json:"externalcache,omitempty"`
	Heurexpiry       string `json:"heurexpiry,omitempty"`
	Heurexpiryhistwt int    `json:"heurexpiryhistwt,omitempty"`
	Heurexpirythres  int    `json:"heurexpirythres,omitempty"`
	Minressize       int    `json:"minressize,omitempty"`
	Policytype       string `json:"policytype,omitempty"`
	Quantumsize      int    `json:"quantumsize,omitempty"`
	Servercmp        string `json:"servercmp,omitempty"`
}
