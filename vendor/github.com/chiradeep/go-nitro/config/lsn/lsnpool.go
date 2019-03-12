package lsn

type Lsnpool struct {
	Maxportrealloctmq   int    `json:"maxportrealloctmq,omitempty"`
	Nattype             string `json:"nattype,omitempty"`
	Poolname            string `json:"poolname,omitempty"`
	Portblockallocation string `json:"portblockallocation,omitempty"`
	Portrealloctimeout  int    `json:"portrealloctimeout,omitempty"`
}
