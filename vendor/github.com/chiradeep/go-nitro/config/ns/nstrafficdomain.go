package ns

type Nstrafficdomain struct {
	Aliasname string `json:"aliasname,omitempty"`
	State     string `json:"state,omitempty"`
	Td        int    `json:"td,omitempty"`
	Vmac      string `json:"vmac,omitempty"`
}
