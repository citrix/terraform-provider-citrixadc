package ns

type Nshostname struct {
	Hostname  string `json:"hostname,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
}
