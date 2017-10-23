package dns

type Dnsparameter struct {
	Cacherecords       string `json:"cacherecords,omitempty"`
	Dns64timeout       int    `json:"dns64timeout,omitempty"`
	Dnsrootreferral    string `json:"dnsrootreferral,omitempty"`
	Dnssec             string `json:"dnssec,omitempty"`
	Maxpipeline        int    `json:"maxpipeline,omitempty"`
	Maxttl             int    `json:"maxttl,omitempty"`
	Minttl             int    `json:"minttl,omitempty"`
	Namelookuppriority string `json:"namelookuppriority,omitempty"`
	Recursion          string `json:"recursion,omitempty"`
	Resolutionorder    string `json:"resolutionorder,omitempty"`
	Retries            int    `json:"retries,omitempty"`
}
