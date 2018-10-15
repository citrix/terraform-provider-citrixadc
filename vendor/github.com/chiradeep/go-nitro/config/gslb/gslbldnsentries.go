package gslb

type Gslbldnsentries struct {
	Ipaddress string      `json:"ipaddress,omitempty"`
	Name      string      `json:"name,omitempty"`
	Nodeid    int         `json:"nodeid,omitempty"`
	Numsites  int         `json:"numsites,omitempty"`
	Rtt       interface{} `json:"rtt,omitempty"`
	Sitename  string      `json:"sitename,omitempty"`
	Ttl       int         `json:"ttl,omitempty"`
}
