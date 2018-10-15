package subscriber

type Subscribersessions struct {
	Avpdisplaybuffer    string      `json:"avpdisplaybuffer,omitempty"`
	Flags               int         `json:"flags,omitempty"`
	Idlettl             int         `json:"idlettl,omitempty"`
	Ip                  string      `json:"ip,omitempty"`
	Nodeid              int         `json:"nodeid,omitempty"`
	Servicepath         string      `json:"servicepath,omitempty"`
	Subscriberrules     interface{} `json:"subscriberrules,omitempty"`
	Subscriptionidtype  string      `json:"subscriptionidtype,omitempty"`
	Subscriptionidvalue string      `json:"subscriptionidvalue,omitempty"`
	Ttl                 int         `json:"ttl,omitempty"`
	Vlan                int         `json:"vlan,omitempty"`
}
