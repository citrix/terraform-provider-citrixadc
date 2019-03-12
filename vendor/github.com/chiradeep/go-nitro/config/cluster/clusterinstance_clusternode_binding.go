package cluster

type Clusterinstanceclusternodebinding struct {
	Clid                       int    `json:"clid,omitempty"`
	Clusterhealth              string `json:"clusterhealth,omitempty"`
	Effectivestate             string `json:"effectivestate,omitempty"`
	Health                     string `json:"health,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Isconfigurationcoordinator bool   `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool   `json:"islocalnode,omitempty"`
	Masterstate                string `json:"masterstate,omitempty"`
	Nodeid                     int    `json:"nodeid,omitempty"`
	Nodejumbonotsupported      bool   `json:"nodejumbonotsupported,omitempty"`
	Nodelicensemismatch        bool   `json:"nodelicensemismatch,omitempty"`
	Nodersskeymismatch         bool   `json:"nodersskeymismatch,omitempty"`
	State                      string `json:"state,omitempty"`
}
