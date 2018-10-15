package lsn

type Lsnsipalgcalldatachannelbinding struct {
	Callid          string `json:"callid,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
}
