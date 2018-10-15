package network

type Rnat6 struct {
	Acl6name         string `json:"acl6name,omitempty"`
	Name             string `json:"name,omitempty"`
	Network          string `json:"network,omitempty"`
	Ownergroup       string `json:"ownergroup,omitempty"`
	Redirectport     int    `json:"redirectport,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,omitempty"`
}
