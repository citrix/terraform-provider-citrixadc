package lsn

type Lsngroup struct {
	Allocpolicy    string `json:"allocpolicy,omitempty"`
	Clientname     string `json:"clientname,omitempty"`
	Ftp            string `json:"ftp,omitempty"`
	Ftpcm          string `json:"ftpcm,omitempty"`
	Groupid        int    `json:"groupid,omitempty"`
	Groupname      string `json:"groupname,omitempty"`
	Ip6profile     string `json:"ip6profile,omitempty"`
	Logging        string `json:"logging,omitempty"`
	Nattype        string `json:"nattype,omitempty"`
	Portblocksize  int    `json:"portblocksize,omitempty"`
	Pptp           string `json:"pptp,omitempty"`
	Rtspalg        string `json:"rtspalg,omitempty"`
	Sessionlogging string `json:"sessionlogging,omitempty"`
	Sessionsync    string `json:"sessionsync,omitempty"`
	Sipalg         string `json:"sipalg,omitempty"`
	Snmptraplimit  int    `json:"snmptraplimit,omitempty"`
}
