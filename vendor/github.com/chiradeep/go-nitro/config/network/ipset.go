package network

type Ipset struct {
	Name string `json:"name,omitempty"`
	Td   int    `json:"td,omitempty"`
}
