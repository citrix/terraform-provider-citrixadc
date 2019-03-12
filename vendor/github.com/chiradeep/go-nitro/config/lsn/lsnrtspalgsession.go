package lsn

type Lsnrtspalgsession struct {
	Callflags    int    `json:"callflags,omitempty"`
	Callrefcount int    `json:"callrefcount,omitempty"`
	Calltimer    int    `json:"calltimer,omitempty"`
	Sessionid    string `json:"sessionid,omitempty"`
	Xlatip       string `json:"xlatip,omitempty"`
}
