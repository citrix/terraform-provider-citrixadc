package cluster

type Clusternode struct {
	Backplane                  string      `json:"backplane,omitempty"`
	Clusterhealth              string      `json:"clusterhealth,omitempty"`
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
	Nodeid                     int         `json:"nodeid,omitempty"`
	Nodelicensemismatch        bool        `json:"nodelicensemismatch,omitempty"`
	Nodelist                   interface{} `json:"nodelist,omitempty"`
	Nodersskeymismatch         bool        `json:"nodersskeymismatch,omitempty"`
	Operationalsyncstate       string      `json:"operationalsyncstate,omitempty"`
	Partialfailifaces          string      `json:"partialfailifaces,omitempty"`
	Priority                   int         `json:"priority,omitempty"`
	State                      string      `json:"state,omitempty"`
}
