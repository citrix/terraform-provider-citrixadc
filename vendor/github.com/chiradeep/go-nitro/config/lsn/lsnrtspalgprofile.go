package lsn

type Lsnrtspalgprofile struct {
	Rtspalgprofilename    string `json:"rtspalgprofilename,omitempty"`
	Rtspidletimeout       int    `json:"rtspidletimeout,omitempty"`
	Rtspportrange         string `json:"rtspportrange,omitempty"`
	Rtsptransportprotocol string `json:"rtsptransportprotocol,omitempty"`
}
