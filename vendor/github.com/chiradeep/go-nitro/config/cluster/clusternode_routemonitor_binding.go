package cluster

type Clusternoderoutemonitorbinding struct {
	Netmask       string `json:"netmask,omitempty"`
	Nodeid        int    `json:"nodeid,omitempty"`
	Routemonitor  string `json:"routemonitor,omitempty"`
	Routemonstate int    `json:"routemonstate,omitempty"`
}
