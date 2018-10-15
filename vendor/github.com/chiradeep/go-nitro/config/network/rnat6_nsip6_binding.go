package network

type Rnat6nsip6binding struct {
	Name       string `json:"name,omitempty"`
	Natip6     string `json:"natip6,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}
