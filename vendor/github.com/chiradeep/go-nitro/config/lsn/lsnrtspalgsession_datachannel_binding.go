package lsn

type Lsnrtspalgsessiondatachannelbinding struct {
	Channelflags    int    `json:"channelflags,omitempty"`
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
	Sessionid       string `json:"sessionid,omitempty"`
}
