package cluster

type Clusternode struct {
	Backplane                  string      `json:"backplane,omitempty"`
	Cfgflags                   int         `json:"cfgflags,omitempty"`
	Clearnodegroupconfig       string      `json:"clearnodegroupconfig,omitempty"`
	Clusterhealth              string      `json:"clusterhealth,omitempty"`
	Delay                      int         `json:"delay"` // 0 is a valid value for this field
	Disabledifaces             string      `json:"disabledifaces,omitempty"`
	Effectivestate             string      `json:"effectivestate,omitempty"`
	Enabledifaces              string      `json:"enabledifaces,omitempty"`
	Hamonifaces                string      `json:"hamonifaces,omitempty"`
	Health                     string      `json:"health,omitempty"`
	Ifaceslist                 interface{} `json:"ifaceslist,omitempty"`
	Ipaddress                  string      `json:"ipaddress,omitempty"`
	Isconfigurationcoordinator bool        `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool        `json:"islocalnode,omitempty"`
	Masterstate                string      `json:"masterstate,omitempty"`
	Name                       string      `json:"name,omitempty"`
	Netmask                    string      `json:"netmask,omitempty"`
	Nodegroup                  string      `json:"nodegroup,omitempty"`
	Nodeid                     int         `json:"nodeid"` // 0 is a valid value for this field
	Nodejumbonotsupported      bool        `json:"nodejumbonotsupported,omitempty"`
	Nodelicensemismatch        bool        `json:"nodelicensemismatch,omitempty"`
	Nodelist                   interface{} `json:"nodelist,omitempty"`
	Nodersskeymismatch         bool        `json:"nodersskeymismatch,omitempty"`
	Operationalsyncstate       string      `json:"operationalsyncstate,omitempty"`
	Partialfailifaces          string      `json:"partialfailifaces,omitempty"`
	Priority                   int         `json:"priority"` // 0 is a valid value for this field
	Routemonitor               string      `json:"routemonitor,omitempty"`
	State                      string      `json:"state,omitempty"`
	Syncstate                  string      `json:"syncstate,omitempty"`
	Tunnelmode                 string      `json:"tunnelmode,omitempty"`
}
