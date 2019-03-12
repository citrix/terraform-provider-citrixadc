package network

type Vridparam struct {
	Deadinterval  int    `json:"deadinterval,omitempty"`
	Hellointerval int    `json:"hellointerval,omitempty"`
	Sendtomaster  string `json:"sendtomaster,omitempty"`
}
