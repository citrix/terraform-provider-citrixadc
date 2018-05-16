package network

type Rnat6 struct {
	Acl6name     string `json:"acl6name,omitempty"`
	Name         string `json:"name,omitempty"`
	Network      string `json:"network,omitempty"`
	Redirectport int    `json:"redirectport,omitempty"`
}
