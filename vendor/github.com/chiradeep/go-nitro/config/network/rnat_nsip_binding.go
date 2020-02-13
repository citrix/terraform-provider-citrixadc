package network

type Rnatnsipbinding struct {
	Name       string `json:"name,omitempty"`
	Natip      string `json:"natip,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}
