package network

type Vpath struct {
	Destip    string `json:"destip,omitempty"`
	Encapmode string `json:"encapmode,omitempty"`
	Name      string `json:"name,omitempty"`
}
