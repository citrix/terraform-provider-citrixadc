package ns

type Nsweblogparam struct {
	Buffersizemb  int         `json:"buffersizemb,omitempty"`
	Customreqhdrs interface{} `json:"customreqhdrs,omitempty"`
	Customrsphdrs interface{} `json:"customrsphdrs,omitempty"`
}
