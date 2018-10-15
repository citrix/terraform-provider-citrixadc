package lsn

type Lsnclientnsaclbinding struct {
	Aclname    string `json:"aclname,omitempty"`
	Clientname string `json:"clientname,omitempty"`
	Td         int    `json:"td,omitempty"`
}
