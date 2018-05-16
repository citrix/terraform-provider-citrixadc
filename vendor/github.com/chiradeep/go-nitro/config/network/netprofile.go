package network

type Netprofile struct {
	Name  string `json:"name,omitempty"`
	Srcip string `json:"srcip,omitempty"`
	Td    int    `json:"td,omitempty"`
}
