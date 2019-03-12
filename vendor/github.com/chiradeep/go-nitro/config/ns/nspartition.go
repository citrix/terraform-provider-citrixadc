package ns

type Nspartition struct {
	Maxbandwidth  int    `json:"maxbandwidth,omitempty"`
	Maxconn       int    `json:"maxconn,omitempty"`
	Maxmemlimit   int    `json:"maxmemlimit,omitempty"`
	Minbandwidth  int    `json:"minbandwidth,omitempty"`
	Partitionid   int    `json:"partitionid,omitempty"`
	Partitionmac  string `json:"partitionmac,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
	Partitiontype string `json:"partitiontype,omitempty"`
	Pmacinternal  bool   `json:"pmacinternal,omitempty"`
}
