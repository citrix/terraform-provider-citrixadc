package network

type Iptunnelparam struct {
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Enablestrictrx       string `json:"enablestrictrx,omitempty"`
	Enablestricttx       string `json:"enablestricttx,omitempty"`
	Mac                  string `json:"mac,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Useclientsourceip    string `json:"useclientsourceip,omitempty"`
}
