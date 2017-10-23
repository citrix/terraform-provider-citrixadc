package network

type Iptunnelparam struct {
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
}
