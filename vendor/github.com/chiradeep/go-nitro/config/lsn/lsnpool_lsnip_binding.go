package lsn

type Lsnpoollsnipbinding struct {
	Lsnip     string `json:"lsnip,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	Poolname  string `json:"poolname,omitempty"`
}
