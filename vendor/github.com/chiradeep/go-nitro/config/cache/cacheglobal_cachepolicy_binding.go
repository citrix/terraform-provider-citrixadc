package cache

type Cacheglobalcachepolicybinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Precededefrules        string `json:"precededefrules,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
}
