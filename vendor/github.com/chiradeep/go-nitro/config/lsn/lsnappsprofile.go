package lsn

type Lsnappsprofile struct {
	Appsprofilename   string `json:"appsprofilename,omitempty"`
	Filtering         string `json:"filtering,omitempty"`
	Ippooling         string `json:"ippooling,omitempty"`
	L2info            string `json:"l2info,omitempty"`
	Mapping           string `json:"mapping,omitempty"`
	Tcpproxy          string `json:"tcpproxy,omitempty"`
	Td                int    `json:"td,omitempty"`
	Transportprotocol string `json:"transportprotocol,omitempty"`
}
