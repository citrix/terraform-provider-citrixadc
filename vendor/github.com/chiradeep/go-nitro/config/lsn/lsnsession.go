package lsn

type Lsnsession struct {
	Clientname        string `json:"clientname,omitempty"`
	Destip            string `json:"destip,omitempty"`
	Destport          int    `json:"destport,omitempty"`
	Dsttd             int    `json:"dsttd,omitempty"`
	Ipv6address       string `json:"ipv6address,omitempty"`
	Natip             string `json:"natip,omitempty"`
	Natport           int    `json:"natport,omitempty"`
	Natport2          int    `json:"natport2,omitempty"`
	Natprefix         string `json:"natprefix,omitempty"`
	Nattype           string `json:"nattype,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Network           string `json:"network,omitempty"`
	Network6          string `json:"network6,omitempty"`
	Nodeid            int    `json:"nodeid,omitempty"`
	Sessionestdir     string `json:"sessionestdir,omitempty"`
	Srctd             int    `json:"srctd,omitempty"`
	Subscrip          string `json:"subscrip,omitempty"`
	Subscrport        int    `json:"subscrport,omitempty"`
	Td                int    `json:"td,omitempty"`
	Transportprotocol string `json:"transportprotocol,omitempty"`
}
