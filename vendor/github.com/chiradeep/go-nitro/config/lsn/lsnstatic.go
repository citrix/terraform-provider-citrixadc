package lsn

type Lsnstatic struct {
	Destip            string `json:"destip,omitempty"`
	Dsttd             int    `json:"dsttd,omitempty"`
	Name              string `json:"name,omitempty"`
	Natip             string `json:"natip,omitempty"`
	Natport           int    `json:"natport,omitempty"`
	Nattype           string `json:"nattype,omitempty"`
	Network6          string `json:"network6,omitempty"`
	Status            string `json:"status,omitempty"`
	Subscrip          string `json:"subscrip,omitempty"`
	Subscrport        int    `json:"subscrport,omitempty"`
	Td                int    `json:"td,omitempty"`
	Transportprotocol string `json:"transportprotocol,omitempty"`
}
