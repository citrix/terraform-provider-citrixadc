package lsn

type Lsnip6profile struct {
	Name      string `json:"name,omitempty"`
	Natprefix string `json:"natprefix,omitempty"`
	Network6  string `json:"network6,omitempty"`
	Type      string `json:"type,omitempty"`
}
