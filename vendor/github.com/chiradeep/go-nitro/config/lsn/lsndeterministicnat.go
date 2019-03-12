package lsn

type Lsndeterministicnat struct {
	Clientname string `json:"clientname,omitempty"`
	Firstport  int    `json:"firstport,omitempty"`
	Lastport   int    `json:"lastport,omitempty"`
	Natip      string `json:"natip,omitempty"`
	Natip2     string `json:"natip2,omitempty"`
	Natprefix  string `json:"natprefix,omitempty"`
	Nattype    string `json:"nattype,omitempty"`
	Network6   string `json:"network6,omitempty"`
	Srctd      int    `json:"srctd,omitempty"`
	Subscrip   string `json:"subscrip,omitempty"`
	Subscrip2  string `json:"subscrip2,omitempty"`
	Td         int    `json:"td,omitempty"`
}
