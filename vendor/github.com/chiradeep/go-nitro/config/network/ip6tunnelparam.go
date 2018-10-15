package network

type Ip6tunnelparam struct {
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Useclientsourceipv6  string `json:"useclientsourceipv6,omitempty"`
}
