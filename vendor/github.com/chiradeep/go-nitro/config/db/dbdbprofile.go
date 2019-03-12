package db

type Dbdbprofile struct {
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Name                   string `json:"name,omitempty"`
	Refcnt                 int    `json:"refcnt,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
}
