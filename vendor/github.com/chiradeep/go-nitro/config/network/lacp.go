package network

type Lacp struct {
	Clustermac         string `json:"clustermac,omitempty"`
	Clustersyspriority int    `json:"clustersyspriority,omitempty"`
	Devicename         string `json:"devicename,omitempty"`
	Flags              int    `json:"flags,omitempty"`
	Lacpkey            int    `json:"lacpkey,omitempty"`
	Mac                string `json:"mac,omitempty"`
	Ownernode          int    `json:"ownernode,omitempty"`
	Syspriority        int    `json:"syspriority,omitempty"`
}
