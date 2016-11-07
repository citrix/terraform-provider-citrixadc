package basic

type Nstrace struct {
	Fileid        string      `json:"fileid,omitempty"`
	Filename      string      `json:"filename,omitempty"`
	Filter        string      `json:"filter,omitempty"`
	Link          string      `json:"link,omitempty"`
	Mode          interface{} `json:"mode,omitempty"`
	Nf            int         `json:"nf,omitempty"`
	Nodes         interface{} `json:"nodes,omitempty"`
	Pernic        string      `json:"pernic,omitempty"`
	Scope         string      `json:"scope,omitempty"`
	Size          int         `json:"size,omitempty"`
	State         string      `json:"state,omitempty"`
	Tcpdump       string      `json:"tcpdump,omitempty"`
	Time          int         `json:"time,omitempty"`
	Tracelocation string      `json:"tracelocation,omitempty"`
}
