package subscriber

type Subscriberprofile struct {
	Avpdisplaybuffer    string      `json:"avpdisplaybuffer,omitempty"`
	Flags               int         `json:"flags,omitempty"`
	Ip                  string      `json:"ip,omitempty"`
	Servicepath         string      `json:"servicepath,omitempty"`
	Subscriberrules     interface{} `json:"subscriberrules,omitempty"`
	Subscriptionidtype  string      `json:"subscriptionidtype,omitempty"`
	Subscriptionidvalue string      `json:"subscriptionidvalue,omitempty"`
	Ttl                 int         `json:"ttl,omitempty"`
	Vlan                int         `json:"vlan,omitempty"`
}
