package lsn

type Lsnlogprofile struct {
	Logcompact     string `json:"logcompact,omitempty"`
	Logipfix       string `json:"logipfix,omitempty"`
	Logprofilename string `json:"logprofilename,omitempty"`
	Logsubscrinfo  string `json:"logsubscrinfo,omitempty"`
}
