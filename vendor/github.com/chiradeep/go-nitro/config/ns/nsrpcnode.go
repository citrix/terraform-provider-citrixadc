package ns

type Nsrpcnode struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Password  string `json:"password,omitempty"`
	Secure    string `json:"secure,omitempty"`
	Srcip     string `json:"srcip,omitempty"`
}
