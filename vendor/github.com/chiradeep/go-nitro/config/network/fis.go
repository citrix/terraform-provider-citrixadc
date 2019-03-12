package network

type Fis struct {
	Ifaces    string `json:"ifaces,omitempty"`
	Name      string `json:"name,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
}
