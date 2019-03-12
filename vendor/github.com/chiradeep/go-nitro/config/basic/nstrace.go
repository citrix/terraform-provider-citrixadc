package basic

type Nstrace struct {
	Capdroppkt       string      `json:"capdroppkt,omitempty"`
	Capsslkeys       string      `json:"capsslkeys,omitempty"`
	Doruntimecleanup string      `json:"doruntimecleanup,omitempty"`
	Fileid           string      `json:"fileid,omitempty"`
	Filename         string      `json:"filename,omitempty"`
	Filesize         int         `json:"filesize,omitempty"`
	Filter           string      `json:"filter,omitempty"`
	Inmemorytrace    string      `json:"inmemorytrace,omitempty"`
	Link             string      `json:"link,omitempty"`
	Merge            string      `json:"merge,omitempty"`
	Mode             interface{} `json:"mode,omitempty"`
	Nf               int         `json:"nf,omitempty"`
	Nodeid           int         `json:"nodeid,omitempty"`
	Nodes            interface{} `json:"nodes,omitempty"`
	Pernic           string      `json:"pernic,omitempty"`
	Scope            string      `json:"scope,omitempty"`
	Size             int         `json:"size,omitempty"`
	Skiplocalssh     string      `json:"skiplocalssh,omitempty"`
	Skiprpc          string      `json:"skiprpc,omitempty"`
	State            string      `json:"state,omitempty"`
	Time             int         `json:"time,omitempty"`
	Tracebuffers     int         `json:"tracebuffers,omitempty"`
	Traceformat      string      `json:"traceformat,omitempty"`
	Tracelocation    string      `json:"tracelocation,omitempty"`
}
