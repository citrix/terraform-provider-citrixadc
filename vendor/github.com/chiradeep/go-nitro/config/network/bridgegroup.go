package network

type Bridgegroup struct {
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Flags              bool   `json:"flags,omitempty"`
	Id                 int    `json:"id,omitempty"`
	Ifaces             string `json:"ifaces,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Portbitmap         int    `json:"portbitmap,omitempty"`
	Rnat               bool   `json:"rnat,omitempty"`
	Tagbitmap          int    `json:"tagbitmap,omitempty"`
	Tagifaces          string `json:"tagifaces,omitempty"`
}
